package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ismoilroziboyev/go-pkg/errors"
)

func (h *Handler) authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username, password, ok := ctx.Request.BasicAuth()

		if !ok {
			h.handleError(ctx, errors.NewUnauthorizedErrorf("basic auth is required"))
			return
		}

		if username != h.cfg.AdminUser {
			h.handleError(ctx, errors.NewUnauthorizedErrorf("user not found"))
			return
		}

		if password != h.cfg.AdminPassword {
			h.handleError(ctx, errors.NewUnauthorizedErrorf("wrong password"))
			return
		}

		ctx.Next()
	}
}
