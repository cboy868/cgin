package cgin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	*gin.Engine
}

// Context 拓展gin.Context后的Context
type Context struct {
	*gin.Context
	ID int
}

// RouterGroup 拓展后的自定义gin.RouterGroup
type RouterGroup struct {
	*gin.RouterGroup
}

// NewServer 新建gin服务器
func NewServer() *Server {
	server := &Server{Engine: gin.New()}
	return server
}

func Default() *Server {
	server := &Server{Engine: gin.Default()}
	return server
}

// handleFunc 实现gin.Context到自定义Context的转换。
func handleFunc(handler func(c *Context)) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		handler(&Context{Context: c})
	}
}

// Group 重写路由组注册
func (server *Server) Group(relativePath string, handlers ...func(c *Context)) *RouterGroup {
	RHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		RHandles = append(RHandles, handleFunc(handle))
	}
	return &RouterGroup{server.Engine.Group(relativePath, RHandles...)}
}

// GET 拓展Get请求（根）
func (server *Server) GET(relativePath string, handlers ...func(c *Context)) gin.IRoutes {
	RHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		RHandles = append(RHandles, handleFunc(handle))
	}
	return server.Engine.GET(relativePath, RHandles...)
}

// POST 拓展POST请求（根）
func (server *Server) POST(relativePath string, handlers ...func(c *Context)) gin.IRoutes {
	RHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		RHandles = append(RHandles, handleFunc(handle))
	}
	return server.Engine.POST(relativePath, RHandles...)
}

// Group 重写路由组注册
func (r *RouterGroup) Group(relativePath string, handlers ...func(c *Context)) *RouterGroup {
	RHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		RHandles = append(RHandles, handleFunc(handle))
	}
	return &RouterGroup{r.RouterGroup.Group(relativePath, RHandles...)}
}

// GET 拓展Get请求（子）
func (r *RouterGroup) GET(relativePath string, handlers ...func(c *Context)) gin.IRoutes {
	rHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		rHandles = append(rHandles, handleFunc(handle))
	}
	return r.RouterGroup.GET(relativePath, rHandles...)
}

// POST 拓展Post请求（子）
func (r *RouterGroup) POST(relativePath string, handlers ...func(c *Context)) gin.IRoutes {
	rHandles := make([]gin.HandlerFunc, 0)
	for _, handle := range handlers {
		rHandles = append(rHandles, handleFunc(handle))
	}
	return r.RouterGroup.POST(relativePath, rHandles...)
}

// Use 拓展中间件注册
func (r *RouterGroup) Use(middlewares ...func(c *Context))*RouterGroup {
	rMiddlewares := make([]gin.HandlerFunc, 0)
	for _, middleware := range middlewares {
		rMiddlewares = append(rMiddlewares, handleFunc(middleware))
	}
	r.RouterGroup.Use(rMiddlewares...)
	return &RouterGroup{r.RouterGroup}
}


func (c *Context) Success(data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Success",
		"data":    data,
	})
}

func (c *Context) Error(code interface{}, err interface{}, data ...interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": err,
		"data":    data,
	})
}
