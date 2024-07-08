package fly

import (
	"fmt"
	"github.com/martinadamsdev/go-fly/route"
	"log"
	"net/http"
)

// 定义支持的运行模式
const (
	DebugMode   = "debug"
	TestMode    = "test"
	ReleaseMode = "release"
)

// Engine 结构体现在包含运行模式。
type Engine struct {
	router *route.Router
	mode   string
}

// Default 创建并返回一个带有默认配置的 Engine 实例，默认为 Debug 模式。
func Default() *Engine {
	router := route.NewRouter()
	engine := &Engine{
		router: router,
		mode:   DebugMode, // 默认为 Debug 模式
	}
	// 添加日志和错误恢复中间件
	engine.Use(route.LoggingMiddleware)
	engine.Use(route.RecoverMiddleware)
	return engine
}

// SetMode 设置 Engine 的运行模式。
func (e *Engine) SetMode(mode string) {
	e.mode = mode
}

// Use 添加中间件到路由器。
func (e *Engine) Use(middleware route.Middleware) {
	// 根据模式调整日志中间件的行为
	if e.mode == DebugMode {
		log.Println("Adding middleware in debug mode")
	}
	e.router.Use(middleware)
}

// Run 启动HTTP服务器。
func (e *Engine) Run(addr string) error {
	if e.mode == DebugMode {
		log.Printf("Debug mode is enabled")
	}
	fmt.Printf("Listening and serving HTTP on %s\n", addr)
	return http.ListenAndServe(addr, e.router)
}
