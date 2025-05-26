package router

import (
	"github.com/gin-gonic/gin"
	"github.com/usmaarn/blogg_api/internal/handler"
)

func V1Router(router *gin.Engine) {
	//Authentication
	r := router.Group("/v1")
	r.POST("/auth/register", handler.RegisterHandler)
	r.POST("/auth/login", handler.LoginHandler)
	//
}
