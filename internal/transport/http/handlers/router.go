package handlers

import (
	"net/http"

	_ "github.com/ismoilroziboyev/log-minder/api/openapi"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (h *Handler) API() http.Handler {
	engine := h.engine

	router := engine.Group("v1")
	{
		router.POST("/logs", h.authenticate(), h.insertLog)
		router.POST("/logs/retreive", h.authenticate(), h.retreiveLogs)

		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return engine
}
