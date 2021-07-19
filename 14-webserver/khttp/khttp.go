package khttp

import "net/http"

type ServeMux struct {
	routes map[string]http.Handler
}

func (h *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	route := r.URL.Path
	handler := h.routes[route]

	handler.ServeHTTP(w,r)

}

func (h *ServeMux) Handle(route string, handler http.Handler) {
	h.routes[route] = handler
}

func NewServeMux() *ServeMux {
	mux := ServeMux{}
	mux.routes = make(map[string]http.Handler)
	return &mux
}


type DefaultServeMux struct {
	routes map[string]func(w http.ResponseWriter, r *http.Request)
}

func (h *DefaultServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	route := r.URL.Path
	handler := h.routes[route]

	handler(w,r)

}
var defaultMux DefaultServeMux
defaultMux.route = make(map[string]func(w http.ResponseWriter, r *http.Request))

func HandleFunc(route string, handleFunc func(w http.ResponseWriter, r *http.Request)) {
	DefaultServeMux.route[route] = handleFunc
}




