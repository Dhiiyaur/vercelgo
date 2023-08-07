package api

import (
	"net/http"

	. "github.com/tbxark/g4vercel"
)

func HandlerHello(w http.ResponseWriter, r *http.Request) {

	server := New()
	server.POST("/hello", func(context *Context) {
		context.JSON(200, H{
			"data": "go",
		})
	})
	server.Handle(w, r)

}
