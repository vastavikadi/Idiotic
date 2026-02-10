package models

import "time"

type Video struct {
	ID           uint      `gorm:"primaryKey"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	ThumbnailUrl *string    `json:"thumbnailurl"`
	UploadedBy   string    `json:"uploadedBy"`
	PublishedAt  time.Time `json:"publishedat"`
	Tags         *[]string  `json:"tags,omitempty"`
	Duration     int       `json:"duration"`
	Status       string    `json:"status"`
	CategoryID   string    `json:"categoryid"`
	ViewCount    int       `json:"viewcount"`
	LikeCount    int       `json:"likecount"`
	DislikeCount int       `json:"dislikecount"`
	CommentCount int       `json:"commentcount"`
}
