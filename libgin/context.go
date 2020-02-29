package libgin

import (
	"github.com/gin-gonic/gin"
)

type HandlerFunc func(c *Context)

func Handle(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{
			c,
		}
		h(ctx)
	}
}

type Context struct {
	*gin.Context
}

func (c *Context) Success() {

}
