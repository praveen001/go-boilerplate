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
			catValues[i] = models.MediaCategoryReverse[c]
		}
		tx = tx.Where("category IN (?)", catValues)
	} else {
		tx = tx.Where("category = ?", models.MediaCategoryReverse["media"])
	}

	if search := params.Get("search"); search != "" {
		sq := strings.ToLower(search) + "%"
		tx = tx.Where("LOWER(title) LIKE ? or LOWER(asset_id) LIKE ?", sq, sq)
	}

	total := 0
	tx = tx.Preload("Segments").Related(&medias, "Medias").Find(&medias).Count(&total)

	limit := params.Get("limit")
	if limit != "" {
		tx = tx.Limit(limit)
	} else {
		tx = tx.Limit(25)
	}

	offset := params.Get("offset")
	if offset != "" {
		tx = tx.Offset(offset)
	} else {
		tx = tx.Offset(0)
	}

	if err := tx.Find(&medias).Error; err != nil {
		return nil, err
	}

	resp := map[string]interface{}{
		"medias": medias,
		"total":  total,
		"offset": offset,
		"limit":  limit,
	}

	return resp, nil
}
