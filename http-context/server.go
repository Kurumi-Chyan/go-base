package main

import (
	"net/http"
)

type Server interface {
	Route(method string, pattern string, handlerFunc func(ctx *Context))
	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler *Handle
}

// Route 路由注册
func (s *sdkHttpServer) Route(method string,
	pattern string,
	handlerFunc func(ctx *Context)) {

	key := s.handler.Key(method, pattern)
	s.handler.handlers[key] = handlerFunc

}

func (s *sdkHttpServer) Start(address string) error {
	http.Handle("/", s.handler)
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}
}
