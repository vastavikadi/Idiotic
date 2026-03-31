package models

type SharedWith struct{
	VideoID uint `gorm:"primaryKey"`
	SharedWithEmails *[]string `json:"sharedWithEmails"`
}