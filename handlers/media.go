package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/praveen001/go-boilerplate/handlers/ctx"
	"github.com/praveen001/go-boilerplate/repository"

	"github.com/praveen001/go-boilerplate/app"
)

// MediaHandler .
type MediaHandler struct {
	media  *repository.MediaRepository
	logger *app.Logger
}

// NewMediaHandler .
func NewMediaHandler(c *app.Context) *MediaHandler {
	return &MediaHandler{
		media:  c.DB.Media,
		logger: c.Logger,
	}
}

// List .
func (h *MediaHandler) List(w http.ResponseWriter, r *http.Request) {
	feed := ctx.GetFeed(r.Context())

	feeds, _ := h.media.FindByFeedID(feed.ID)

	json.NewEncoder(w).Encode(feeds)
}
