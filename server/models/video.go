package models

import "time"

type Video struct {
	ID           uint      `gorm:"primaryKey"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	ThumbnailUrl *string    `json:"thumbnailurl"`
	UploadedBy   string    `json:"uploadedBy"`
	UploadedAt  time.Time `json:"uploadedat"`
	UpdatedAt time.Time `json:"updatedat"`
	Tags         *[]string  `json:"tags,omitempty"`
	Duration     int       `json:"duration"`
	CategoryID   string    `json:"categoryid"`
}
