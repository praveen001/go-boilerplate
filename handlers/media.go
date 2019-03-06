package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/praveen001/go-boilerplate/models"

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

	var medias []*models.Media
	r.ParseForm()
	params := r.Form

	tx := h.media.DB().Model(feed).Related(&medias, "Medias")

	if categories, ok := params["category"]; ok {
		tx = tx.Where("category IN (?)", categories)
	}

	if search := params.Get("search"); search != "" {
		sq := strings.ToLower(search) + "%"
		tx = tx.Where("LOWER(title) LIKE ? or LOWER(asset_id) LIKE ?", sq, sq)
	}

	if limit := params.Get("limit"); limit != "" {
		tx = tx.Limit(limit)
	} else {
		tx = tx.Limit(25)
	}

	if offset := params.Get("offset"); offset != "" {
		tx = tx.Offset(offset)
	} else {
		tx = tx.Offset(0)
	}

	if err := tx.Find(&medias).Error; err != nil {
		h.logger.Error("Unable to fetch medias", err.Error())
	}

	json.NewEncoder(w).Encode(medias)
}
