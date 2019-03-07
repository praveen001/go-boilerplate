package models

import "time"

// Account .
type Account struct {
	ID        uint      `json:"id" gorm:"PRIMARY_KEY"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	Name    string `json:"name"`
	Domain  string `json:"-"`
	IP      string `json:"-"`
	Storage string `json:"-"`
	Meta    string `json:"-"`
}
