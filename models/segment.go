package models

import (
	"strconv"
	"strings"
	"time"
)

/*
{
	"segment_id": 1,
	"default_segment": true,
	"segment_offset": 0,
	"duration": 125,
	"file_offset": 5640,
	"start_timecode": 2040,
	"segment_offset_tc": "00:00:00:00",
	"duration_tc": "00:00:05:00",
	"som": "00:00:00:00",
	"eom": "00:00:05:00"
}
*/

// Segment .
type Segment struct {
	ID        int       `json:"id" gorm:"PRIMARY_KEY"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	// Belongs to Media
	Media   *Media `json:"-"`
	MediaID int    `json:"-"`

	SegmentID int    `json:"segmentId"`
	Data      string `json:"-"`

	DefaultSegment bool `json:"defaultSegment" gorm:"-"`
	Duration       int  `json:"duration" gorm:"-"`
	Offset         int  `json:"offset"` // Used to compute SOM and EO gorm:"-"M
	FileOffset     int  `json:"-" gorm:"-"`
	StartTimecode  int  `json:"-" gorm:"-"`
}

// AfterFind .
func (s *Segment) AfterFind() error {
	if len(s.Data) == 0 {
		return nil
	}

	data := strings.TrimPrefix(s.Data, "---")
	for _, keyvalue := range strings.Split(data, "\n") {
		if keyvalue == "" {
			continue
		}

		keyvalarr := strings.Split(keyvalue, ": ")
		key := keyvalarr[0]
		val := keyvalarr[1]

		switch key {
		case "default_segment":
			s.DefaultSegment, _ = strconv.ParseBool(val)

		case "duration":
			s.Duration, _ = strconv.Atoi(val)

		case "file_offset":
			s.FileOffset, _ = strconv.Atoi(val)

		case "start_timecode":
			s.StartTimecode, _ = strconv.Atoi(val)

		case "offset":
			s.Offset, _ = strconv.Atoi(val)

		}
	}

	return nil
}
