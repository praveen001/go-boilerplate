package models

import "time"

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

// Media .
type Media struct {
	ID        uint      `json:"id" gorm:"PRIMARY_KEY"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	// Belongs to many Feeds
	Feeds []*Feed `json:"feeds" gorm:"MANY2MANY:feeds_media"`

	AssetID          string `json:"assetId"`
	Title            string `json:"title"`
	DurationInFrames uint   `json:"-" gorm:"COLUMN:duration"`
	ImagePreviewSrc  string `json:"imagePreviewSrc"`
	Status           string `json:"status" gorm:"COLUMN:aasm_state"`

	Duration uint `json:"duration" gorm:"-"`
}
