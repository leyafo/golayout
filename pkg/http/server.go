package http

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golayout/pkg/logger"
	"net/http"
	"path"
)

type Logical interface {
	MiddleWare(next http.Handler)http.Handler
	RegisterHandlers()(version string, routers []struct{
		Method string
		Path string
		Handler http.HandlerFunc
	})
}

type Server struct{
	chiRouter *chi.Mux
	addr string
}

func NewServer(host string, port int)*Server{
	s := &Server{}
	s.chiRouter = chi.NewRouter()
	s.chiRouter.Use(middleware.Logger)
	s.addr = fmt.Sprintf("%s:%d", host, port)
	return s
}

func (s *Server)Run()error{
	logger.Infof("listening %s", s.addr)
	return http.ListenAndServe(s.addr, s.chiRouter)
}

func (s *Server)RouterRegister(l Logical)error{
	s.chiRouter.Group(func(r chi.Router) {
		r.Use(l.MiddleWare)
		version, routers := l.RegisterHandlers()
		for _, router := range routers {
			path := path.Clean(fmt.Sprintf("/%s/%s", version, router.Path))
			switch router.Method {
			case "GET":
				r.Get(path, router.Handler)
			case "HEAD":
				r.Head(path, router.Handler)
			case "PUT":
				r.Put(path, router.Handler)
			case "POST":
				r.Post(path, router.Handler)
			case "PATCH":
				r.Patch(path, router.Handler)
			case "DELETE":
				r.Delete(path, router.Handler)
			case "CONNECT":
				r.Connect(path, router.Handler)
			case "OPTIONS":
				r.Options(path, router.Handler)
			case "TRACE":
				r.Trace(path, router.Handler)
			default:
				logger.Errorf("unsupported method: %s",
					router.Method)
			}
		}
	})
	return nil
}