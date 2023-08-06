package api

import (
	"net/http"
	"verceltest/service"

	. "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	server := New()
	server.GET("/", func(context *Context) {

		// ctx := context.Background()
		data, _ := service.MangatownIndex(context.Req.Context())
		context.JSON(200, H{
			"data": data,
		})
	})

	server.Handle(w, r)
}
