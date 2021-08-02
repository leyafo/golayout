package api

import (
	"github.com/go-chi/jwtauth/v5"
	"golayout/pkg/daemon"
	"golayout/pkg/etcd"
	"google.golang.org/grpc"
	"net/http"
)

//Server server for HTTP API
type Server struct{
	opt *daemon.ApiOption
	grpcMonitorBackend *grpc.ClientConn
}

func (s *Server) MiddleWare(next http.Handler) http.Handler {
	return jwtauth.Authenticator(next)
}

func NewServer(monitorServiceName string, opt *daemon.ApiOption)(*Server, error){
	var err error
	s := &Server{opt: opt}
	s.grpcMonitorBackend, err = etcd.DialGrpc(opt.Etcd.EndpointsKey, monitorServiceName, grpc.WithInsecure())
	if err != nil{
		return nil, err
	}
	return s, nil
}

type Router struct {
	Method string
	Path string
	Handler http.HandlerFunc
}

func (s *Server) RegisterHandlers()(version string, routers []struct{
	Method string
	Path string
	Handler http.HandlerFunc
}){
	routers = []struct{
		Method string
		Path string
		Handler http.HandlerFunc
	}{
		{ Method: "GET", Path:"/version", Handler: s.Version},
	}

	return "v1", routers
}
