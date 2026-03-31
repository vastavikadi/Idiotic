package models


type AI_Description struct {
	VideoID uint `gorm:"primaryKey"`
	Description string `json:"description"`
}