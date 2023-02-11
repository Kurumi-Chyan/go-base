package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (Context *Context) ReadJson(req interface{}) error {
	r := Context.R
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, req)
	if err != nil {
		return err
	}
	return nil
}

func (Context *Context) WriteJson(status int, data interface{}) error {
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = Context.W.Write(bs)
	if err != nil {
		return err
	}
	Context.W.WriteHeader(status)
	return nil
}

func (Context *Context) OkJson(data interface{}) error {
	return Context.WriteJson(http.StatusOK, data)
}

func (Context *Context) SystemErrJson(data interface{}) error {
	return Context.WriteJson(http.StatusInternalServerError, data)
}

func (Context *Context) BadRequestJson(data interface{}) error {
	return Context.WriteJson(http.StatusBadRequest, data)
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		R: r,
		W: w,
	}
}
