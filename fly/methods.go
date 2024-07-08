package fly

import "net/http"

// Get 方法用于添加处理 Get 请求的路由。
func (e *Engine) Get(path string, handler func(http.ResponseWriter, *http.Request)) {
	e.router.Get(path, http.HandlerFunc(handler))
}

// Post 方法用于添加处理 Post 请求的路由。
func (e *Engine) Post(path string, handler func(http.ResponseWriter, *http.Request)) {
	e.router.Post(path, http.HandlerFunc(handler))
}

// Put 方法用于添加处理 Put 请求的路由。
func (e *Engine) Put(path string, handler func(http.ResponseWriter, *http.Request)) {
	e.router.Put(path, http.HandlerFunc(handler))
}

// Delete 方法用于添加处理 Delete 请求的路由。
func (e *Engine) Delete(path string, handler func(http.ResponseWriter, *http.Request)) {
	e.router.Delete(path, http.HandlerFunc(handler))
}

// Patch 方法用于添加处理 Patch 请求的路由。
func (e *Engine) Patch(path string, handler func(http.ResponseWriter, *http.Request)) {
	e.router.Delete(path, http.HandlerFunc(handler))
}

// Options 方法用于添加处理 Options 请求的路由。
func (e *Engine) Options(path string, handler func(http.ResponseWriter, *http.Request)) {
	e.router.Delete(path, http.HandlerFunc(handler))
}

// Head 方法用于添加处理 Head 请求的路由。
func (e *Engine) Head(path string, handler func(http.ResponseWriter, *http.Request)) {
	e.router.Delete(path, http.HandlerFunc(handler))
}
