package controller

//controller/Index.go

import (
	"myGin/response"

	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) *response.Response {

	if 1+1 == 2 {

		return response.Resp().Json(gin.H{"msg": "hello world"})

	}

	return response.Resp().Json(gin.H{"msg": "hello world"})
}
