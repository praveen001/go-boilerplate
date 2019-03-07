package repository

import (
	"net/url"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/praveen001/go-boilerplate/models"
)

// MediaRepository .
type MediaRepository struct {
	db *gorm.DB
}

// NewMediaRepository .
func NewMediaRepository(db *gorm.DB) *MediaRepository {
	return &MediaRepository{db}
}

// FilterMedia .
func (r *MediaRepository) FilterMedia(feed *models.Feed, params url.Values) (map[string]interface{}, error) {
	var medias []*models.Media

	tx := r.db.Model(feed)
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
		return nil, err
	}

	resp := map[string]interface{}{
		"media": medias,
		"total": total,
	}

	return resp, nil
}
