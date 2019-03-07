package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/praveen001/go-boilerplate/repository"

	"github.com/praveen001/go-boilerplate/app"
	"github.com/praveen001/go-boilerplate/handlers/ctx"
)

// MediaHandler .
type MediaHandler struct {
	logger *app.Logger

	media *repository.MediaRepository
}

// NewMediaHandler .
func NewMediaHandler(c *app.Context) *MediaHandler {
	return &MediaHandler{
		logger: c.Logger,

		media: c.DB.Media,
	}
}

// List .
func (h *MediaHandler) List(w http.ResponseWriter, r *http.Request) {
	feed := ctx.GetFeed(r.Context())
	r.ParseForm()

	resp, err := h.media.FilterMedia(feed, r.Form)
	if err != nil {
		h.logger.Error("Unable to fetch medias", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
