package routes

////routes/router.go

import (
	"myGin/context"
	"myGin/controller"
	"myGin/response"

	"github.com/gin-gonic/gin"
)

// 这里可以按照层级结构配置路由还有路由组
func Load(r *gin.Engine) {

	router := newRouter(r)

	router.Group("/api", func(api group) {

		api.Registered(GET, "/info", controller.Index)

		api.Group("/user", func(user group) {

			user.Registered(GET, "/info", controller.Index)
			user.Registered(GET, "/order", controller.Index)
			user.Registered(GET, "/money", controller.Index)

		})

	})

}

type router struct {
	engine *gin.Engine
}

type group struct {
	engine *gin.Engine
	path   string
}

type method int

const (
	GET    method = 0x000000
	POST   method = 0x000001
	PUT    method = 0x000002
	DELETE method = 0x000003
	ANY    method = 0x000004
)

func newRouter(engine *gin.Engine) *router {

	return &router{
		engine: engine,
	}
}

func (r *router) Group(path string, callback func(group)) {

	callback(group{
		engine: r.engine,
		path:   path,
	})

}

func (g group) Group(path string, callback func(group)) {

	g.path += path

	callback(g)
}

func (g group) Registered(method method, url string, action func(*context.Context) *response.Response) {

	handlerFunc := convert(action)

	finalUrl := g.path + url

	switch method {

	case GET:

		g.engine.GET(finalUrl, handlerFunc)

	case POST:

		g.engine.GET(finalUrl, handlerFunc)

	case PUT:

		g.engine.PUT(finalUrl, handlerFunc)

	case DELETE:

		g.engine.DELETE(finalUrl, handlerFunc)

	case ANY:

		g.engine.Any(finalUrl, handlerFunc)

	}

}

func convert(f func(*context.Context) *response.Response) gin.HandlerFunc {

	return func(c *gin.Context) {

		resp := f(&context.Context{Context: c})

		data := resp.GetData()

		switch item := data.(type) {

		case string:

			c.String(200, item)

		case gin.H:

			c.JSON(200, item)

		}

	}

}
