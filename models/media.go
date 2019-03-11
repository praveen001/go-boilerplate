package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/praveen001/go-boilerplate/utils"
)

/*
+----------------------+---------------+------+-----+---------+----------------+
| Field                | Type          | Null | Key | Default | Extra          |
+----------------------+---------------+------+-----+---------+----------------+
| id                   | int(11)       | NO   | PRI | NULL    | auto_increment |
| asset_id             | varchar(255)  | YES  | MUL | NULL    |                |
| media_type           | varchar(255)  | YES  |     | NULL    |                |
| title                | varchar(255)  | YES  | MUL | NULL    |                |
| duration             | int(11)       | YES  |     | NULL    |                |
| video_preview_src    | varchar(255)  | YES  |     | NULL    |                |
| image_preview_src    | varchar(255)  | YES  |     | NULL    |                |
| video_bit_rate       | int(11)       | YES  |     | NULL    |                |
| resolution           | varchar(255)  | YES  |     | NULL    |                |
| status               | int(11)       | YES  |     | 0       |                |
| uploaded_by_id       | int(11)       | YES  |     | NULL    |                |
| created_at           | datetime      | YES  |     | NULL    |                |
| updated_at           | datetime      | YES  |     | NULL    |                |
| asset_src            | varchar(255)  | YES  |     | NULL    |                |
| old_priority         | int(11)       | YES  | MUL | 0       |                |
| filename             | varchar(255)  | YES  |     | NULL    |                |
| md5sum               | varchar(255)  | YES  |     | NULL    |                |
| size                 | bigint(20)    | YES  |     | NULL    |                |
| size_uploaded        | bigint(20)    | YES  |     | 0       |                |
| category             | int(11)       | YES  | MUL | 0       |                |
| aasm_state           | varchar(255)  | YES  | MUL | initial |                |
| upload_start_time    | datetime      | YES  |     | NULL    |                |
| upload_end_time      | datetime      | YES  |     | NULL    |                |
| region_id            | int(11)       | YES  | MUL | NULL    |                |
| meta                 | mediumtext    | YES  |     | NULL    |                |
| parent_media_id      | int(11)       | YES  | MUL | NULL    |                |
| params               | varchar(1024) | YES  |     | NULL    |                |
| template             | varchar(255)  | YES  |     | NULL    |                |
| account_id           | int(11)       | YES  |     | NULL    |                |
| type                 | varchar(255)  | YES  |     | NULL    |                |
| generated            | tinyint(1)    | YES  |     | 0       |                |
| broadcaster_media_id | varchar(255)  | YES  |     | NULL    |                |
| tc_in                | varchar(255)  | YES  |     | NULL    |                |
+----------------------+---------------+------+-----+---------+----------------+
*/

//
var (
	MediaCategory = map[int]string{
		1: "media",
		2: "rescue",
		3: "graphics",
		4: "subtitles",
		5: "audio",
		7: "signatures",
		8: "media_segment",
	}
	MediaCategoryReverse = map[string]int{
		"media":         1,
		"rescue":        2,
		"graphics":      3,
		"subtitles":     4,
		"audio":         5,
		"signatures":    7,
		"media_segment": 8,
	}
)

// Media .
type Media struct {
	ID        int       `json:"id" gorm:"PRIMARY_KEY"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	// Belongs to many Feeds
	Feeds []*Feed `json:"-" gorm:"MANY2MANY:feeds_media"`

	// Has many segments
	Segments []*Segment `json:"segments"`

	AssetID          string `json:"assetId"`
	Title            string `json:"title"`
	DurationInFrames int    `json:"-" gorm:"COLUMN:duration"`
	ImagePreviewSrc  string `json:"imagePreviewSrc"`
	Status           string `json:"status" gorm:"COLUMN:aasm_state"`
	CategoryID       int    `json:"-" gorm:"COLUMN:category"`
	TCInTimecode     string `json:"tc" gorm:"COLUMN:tc_in"`
	Type             string `json:"type" gorm:"COLUMN:media_type"`

	TCIn     int    `json:"tcIn" gorm:"-"`
	Category string `json:"category" gorm:"-"`
	Duration int    `json:"duration" gorm:"-"`

	// TODO:
	// *Dynamic should be read from meta (-_-)
	Dynamic bool `json:"dynamic" gorm:"-"`
}

// AfterFind .
func (m *Media) AfterFind(db *gorm.DB) error {
	m.Category = MediaCategory[m.CategoryID]

	if err := db.Model(m).Related(&m.Feeds, "Feeds").Error; err != nil {
		return err
	}

	if len(m.Feeds) != 0 {
		feed := m.Feeds[0]
		feed.PostFetch()
		m.TCIn = utils.TimecodeToFrames(m.TCInTimecode, feed.FPS)
		m.Duration = utils.FramesToMsec(m.DurationInFrames, feed.FPS)
	}

	return nil
}
