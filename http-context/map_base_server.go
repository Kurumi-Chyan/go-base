package main

import (
	"net/http"
)

type Handle struct {
	handlers map[string]func(ctx *Context)
}

func (h *Handle) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := h.Key(request.Method, request.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		handler(NewContext(writer, request))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("notFound"))
	}

}

func (h *Handle) Key(method string, pattern string) string {
	return method + "#" + pattern
}
