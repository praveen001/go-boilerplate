package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/praveen001/go-boilerplate/handlers/ctx"
	"github.com/praveen001/go-boilerplate/models"
)

// ListMedias .
func (h *Handler) ListMedias(w http.ResponseWriter, r *http.Request) {
	feed := ctx.GetFeed(r.Context())

	var medias []*models.Media
	r.ParseForm()
	params := r.Form

	tx := h.DB.Model(feed)

	if categories, ok := params["category"]; ok {
		catValues := make([]int, len(categories))
		for i, c := range categories {
			catValues[i] = models.MediaCategory[c]
		}
		tx = tx.Where("category IN (?)", catValues)
	} else {
		tx = tx.Where("category = ?", models.MediaCategory["media"])
	}

	if search := params.Get("search"); search != "" {
		sq := strings.ToLower(search) + "%"
		tx = tx.Where("LOWER(title) LIKE ? or LOWER(asset_id) LIKE ?", sq, sq)
	}

	total := 0
	tx = tx.Preload("Segments").Related(&medias, "Medias").Find(&medias).Count(&total)

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
		h.Logger.Error("Unable to fetch medias", err.Error())
	}

	resp := map[string]interface{}{
		"media": medias,
		"total": total,
	}

	json.NewEncoder(w).Encode(resp)
}
