package routes

import (
	"klik/handler"

	"github.com/gin-gonic/gin"
)

func Router(setup *gin.Engine) {
	setup.GET("/klikdaily/stocks", handler.GetAllstock)
	setup.POST("/klikdaily/adjustment", handler.Adjustment)
	setup.GET("/klikdaily/logs", handler.GetLogStock)
}
