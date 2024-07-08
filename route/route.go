package route

import (
	"net/http"
	"strings"
)

type Route struct {
	Pattern string
	Method  string
	Handler http.HandlerFunc
}

type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) AddRoute(pattern, method string, handler http.HandlerFunc) {
	r.routes = append(r.routes, Route{pattern, method, handler})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if strings.EqualFold(route.Pattern, req.URL.Path) && strings.EqualFold(route.Method, req.Method) {
			route.Handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}
