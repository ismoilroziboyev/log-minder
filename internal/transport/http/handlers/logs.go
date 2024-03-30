package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ismoilroziboyev/go-pkg/errors"
	"github.com/ismoilroziboyev/log-minder/internal/domain"
)

// @Summary insert logs request
// @Tags logs
// @Param request body domain.InsertLogPayload true "payload"
// @Success 200 {object} domain.CommonResponse
// @Failure 400 {object} domain.CommonResponse
// @Failure 401 {object} domain.CommonResponse
// @Failure 500 {object} domain.CommonResponse
// @Security UsersAuth
// @Router /logs [post]
func (h *Handler) insertLog(c *gin.Context) {
	ctx, cancel := h.makeContext(c)
	defer cancel()

	var payload domain.InsertLogPayload

	if err := c.ShouldBind(&payload); err != nil {
		h.handleError(c, errors.NewBadRequestErrorw("cannot parse body", err))
		return
	}

	if err := h.service.LogsService.Insert(ctx, &payload); err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, domain.CommonResponse{
		Message:    "sucess",
		StatusCode: http.StatusOK,
	})
}

// @Summary retreive logs request
// @Tags logs
// @Param request body domain.RetreiveLogsFilter true "filter"
// @Success 200 {object} domain.RetreiveLogsResponse
// @Failure 400 {object} domain.CommonResponse
// @Failure 401 {object} domain.CommonResponse
// @Failure 500 {object} domain.CommonResponse
// @Security UsersAuth
// @Router /logs/retreive [post]
func (h *Handler) retreiveLogs(c *gin.Context) {
	ctx, cancel := h.makeContext(c)
	defer cancel()

	var payload domain.RetreiveLogsFilter

	if err := c.ShouldBind(&payload); err != nil {
		h.handleError(c, errors.NewBadRequestErrorw("cannot parse query params", err))
		return
	}

	res, err := h.service.LogsService.Retreive(ctx, &payload)

	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
