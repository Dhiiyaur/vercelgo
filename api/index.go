package api

import (
	"net/http"
	"verceltest/service"

	. "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	server := New()
	server.GET("/", func(context *Context) {
		data, _ := service.MangatownIndex(context.Req.Context())
		context.JSON(200, H{
			"data": data,
		})
	})

	server.POST("/", func(context *Context) {
		context.JSON(200, H{
			"data": "go",
		})
	})
	server.Handle(w, r)
}
