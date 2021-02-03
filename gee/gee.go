package gee

import (
	"net/http"
)

type HandlerFunc func(*Context)

//key由method（post/get）和pattern（地址）构成 比如GET-/index
//value-HandlerFunc是对应的方法
type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method,pattern,handler)
}

//调用GET or POST会注册到router中
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

//封装的http的listenAndServe
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	c :=newContext(w,req)
	engine.router.handle(c)
}
