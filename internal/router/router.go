package router

import (
	"github.com/gin-gonic/gin"
	"github.com/usmaarn/blogg_api/internal/handler"
)

func Initialize(addr string) {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(handler.PanicHandler())
	r.Use(handler.ErrorHandler())

	//V1 Routes
	V1Router(r)

	r.Run(addr)
}
