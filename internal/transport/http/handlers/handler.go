package handlers

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	pkg "github.com/ismoilroziboyev/go-pkg/errors"
	"github.com/ismoilroziboyev/log-minder/internal/config"
	"github.com/ismoilroziboyev/log-minder/internal/domain"
	"github.com/ismoilroziboyev/log-minder/internal/services"
)

type Handler struct {
	cfg     *config.Config
	service *services.Service

	engine *gin.Engine
}

func New(cfg *config.Config, service *services.Service) *Handler {
	var engine *gin.Engine

	defaultConfig := cors.DefaultConfig()
	defaultConfig.AllowAllOrigins = true
	defaultConfig.AllowCredentials = true
	defaultConfig.AllowHeaders = append(defaultConfig.AllowHeaders,
		"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token,"+
			"Authorization, accept, origin, Cache-Control, X-Requested-With",
		"upload-offset, upload-metadata, upload-length, tus-resumable")

	defaultConfig.AllowHeaders = append(defaultConfig.AllowHeaders, "*")
	defaultConfig.AllowMethods = append(defaultConfig.AllowMethods, "OPTIONS")

	engine = gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(cors.New(defaultConfig))

	return &Handler{
		cfg:     cfg,
		service: service,

		engine: engine,
	}
}

func (h *Handler) handleError(c *gin.Context, err error) {
	myErr, ok := err.(pkg.Error)

	if !ok || myErr.Code() == 0 {
		h.handleError(c, pkg.NewInternalServerError(err))
		return
	}

	lang := h.getLang(c)
	errCause := myErr.Error()
	errLangMsg := myErr.ErrorI18nMsg(lang)
	errCode := myErr.Code()

	c.AbortWithStatusJSON(myErr.Code(), domain.CommonResponse{
		Error:      &errCause,
		Message:    errLangMsg,
		StatusCode: errCode,
	})
}

func (h *Handler) makeContext(c *gin.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(c, time.Second*5)
}

func (h *Handler) getUserId(c *gin.Context) string {
	return c.GetString("user_id")
}

func (h *Handler) getUserRole(c *gin.Context) string {
	return c.GetString("role")
}

func (h *Handler) getLang(c *gin.Context) string {
	return c.GetString("lang")
}

func (h *Handler) extractPathID(c *gin.Context, key string) (int64, error) {
	return strconv.ParseInt(c.Param(key), 10, 64)
}
