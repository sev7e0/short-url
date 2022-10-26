package models

import "time"

type ShortLink struct {
	ID         uint       `json:"id" gorm:"primaryKey""`
	Url        string     `json:"url" binding:"min=1"`
	ShortUrl   string     `json:"short_url"`
	CreateTime *time.Time `json:"create_time" gorm:"type:date"`
	UpdateTime *time.Time `json:"update_time" gorm:"type:date"`
}

func (ShortLink) TableName() string {
	return "short_link"
}
