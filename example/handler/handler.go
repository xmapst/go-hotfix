package handler

import (
	"fmt"
	"net/http"
	"time"
)

var svc = new(HttpSvc)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", svc.Index)
	mux.HandleFunc("/hello", svc.Hello)
	return mux
}

type HttpSvc struct {
}

func (h *HttpSvc) Index(w http.ResponseWriter, r *http.Request) {
	h.Hello(w, r)
}

func (h *HttpSvc) Hello(w http.ResponseWriter, r *http.Request) {
	h._test(w, r)
}

func (h *HttpSvc) _test(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "hello %s %s", r.RemoteAddr, time.Now().String())
}