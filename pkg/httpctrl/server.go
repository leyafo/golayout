package httpctrl

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golayout/pkg/logger"
	"io"
	"net/http"
	"path"
	"reflect"
	"runtime"
	"strings"
)

type HandlerDoc func()(doc, input, output string)
type Router struct{
	Method string
	Path string
	Handler http.HandlerFunc
	Doc HandlerDoc
}

type Server struct{
	chiRouter *chi.Mux
	addr string
	middlewares []func(http.Handler) http.Handler
	routers map[string]Router
}

func NewServer(host string, port int)*Server{
	s := &Server{}
	s.chiRouter = chi.NewRouter()
	s.chiRouter.Use(middleware.Logger)
	s.addr = fmt.Sprintf("%s:%d", host, port)
	s.routers = make(map[string]Router)
	return s
}

func (s *Server)AddMiddlewares(middlewares ...func(http.Handler) http.Handler){
	s.middlewares = append(s.middlewares, middlewares...)
}

func (s *Server)Run()error{
	logger.Infof("listening %s", s.addr)

	for _, router:= range s.routers{
		path := router.Path
		switch router.Method {
		case "GET":
			s.chiRouter.Get(path, router.Handler)
		case "HEAD":
			s.chiRouter.Head(path, router.Handler)
		case "PUT":
			s.chiRouter.Put(path, router.Handler)
		case "POST":
			s.chiRouter.Post(path, router.Handler)
		case "PATCH":
			s.chiRouter.Patch(path, router.Handler)
		case "DELETE":
			s.chiRouter.Delete(path, router.Handler)
		case "CONNECT":
			s.chiRouter.Connect(path, router.Handler)
		case "OPTIONS":
			s.chiRouter.Options(path, router.Handler)
		case "TRACE":
			s.chiRouter.Trace(path, router.Handler)
		default:
			return fmt.Errorf("unsupported method: %s",
				router.Method)
		}
	}

	return http.ListenAndServe(s.addr, s.chiRouter)
}

func (s *Server)GenerateDocument(writer io.Writer){
	for _, router := range s.routers {
		apiName := runtime.FuncForPC(reflect.ValueOf(router.Handler).Pointer()).Name()
		apiName = apiName[strings.IndexByte(apiName, '.')+1:]
		if router.Doc != nil{
			doc, input, output := router.Doc()
			logger.Infof("apiName=%s  method=%s path=%s doc=%s input=%s, output=%s",
				apiName, router.Method, router.Path, doc, input, output)
		}
	}
}

func (s *Server) RoutersRegister(version string, routers []Router)error{
	for _, router := range routers {
		path := path.Clean(fmt.Sprintf("/%s/%s", version, router.Path))
		routerKey := fmt.Sprintf("%s_%s", router.Method, path)
		if _, ok := s.routers[routerKey]; ok{
			return fmt.Errorf("version:%s method: %s path: %s is existed",
				version, router.Method, router.Path)
		}
		router.Path = path
		s.routers[routerKey]=router
	}
	return nil
}