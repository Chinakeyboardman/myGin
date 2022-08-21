package routes

//routes/routes.go

import (
	"myGin/controller"
	"myGin/kernel"

	"github.com/gin-gonic/gin"
)

func config(router group) {

	router.Group("/api", func(api group) {

		api.Group("/user", func(user group) {

			user.Registered(GET, "/info", controller.Index)
			user.Registered(GET, "/order", controller.Index)
			user.Registered(GET, "/money", controller.Index)

		})

	})

}

func Load(r *gin.Engine) {

	router := newRouter(r)

	router.Group("", func(g group) {

		config(g)
	}, kernel.Middleware...) //加载全局中间件

}
