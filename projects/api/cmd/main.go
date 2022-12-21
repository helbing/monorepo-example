package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/helbing/monorepo-example/packages/tool"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"key": "value",
		})
	})

	r.GET("/gotools", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"foo": tool.Foo(),
		})
	})

	r.Run()
}
