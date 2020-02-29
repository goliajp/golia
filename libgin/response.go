package libgin

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func API(ctx *gin.Context, function func() error) {
	err := function()
	if err != nil {
		Error(ctx, err)
	}
	ctx.Next()
}

func Res(ctx *gin.Context, res interface{}, err error, extras ...gin.H) error {
	if err == nil {
		Success(ctx, res, extras...)
	} else {
		Error(ctx, err, extras...)
	}
	return nil
}

func Success(ctx *gin.Context, data interface{}, extras ...gin.H) {
	res := gin.H{
		"data":    data,
		"success": true,
	}
	for _, ex := range extras {
		for k, v := range ex {
			res[k] = v
		}
	}
	ctx.JSON(http.StatusOK, res)
}

func Error(ctx *gin.Context, err error, extras ...gin.H) {
	if err != nil {
		res := gin.H{
			"error":   err.Error(),
			"success": false,
		}
		for _, ex := range extras {
			for k, v := range ex {
				res[k] = v
			}
		}
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, res)
	}
}
