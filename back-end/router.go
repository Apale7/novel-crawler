package main

import (
	"novel-crawler-back-end/handler"

	"github.com/gin-gonic/gin"
)

func collectRouter(r *gin.Engine) {
	r.GET("/list", handler.GetList)
	r.GET("/novel", handler.GetNovel)
}
