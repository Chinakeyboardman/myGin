package context

//context/context.go

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

func (c *Context) Domain() string {

	return c.Request.Host[:strings.Index(c.Request.Host, ":")]
}
