package handlers

import "net/http"

func (h *Handler) API() http.Handler {
	engine := h.engine

	return engine
}
