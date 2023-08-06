package handler

import (
	"net/http"

	. "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	server := New()

	server.GET("/", func(context *Context) {
		context.JSON(200, H{
			"message": "text vercel",
		})
	})

	server.Handle(w, r)
}
