package dao

import (
	"short-link-go/global"
	"short-link-go/models"
)

func GetShortByUrl(url string) models.ShortLink {
	shortLink := models.ShortLink{}
	global.DB.Find(&shortLink, models.ShortLink{Url: url})
	return shortLink
}
func Short(short string) models.ShortLink {
	shortLink := models.ShortLink{}
	global.DB.Find(&shortLink, models.ShortLink{ShortUrl: short})
	return shortLink
}

func SaveShortUrl(model *models.ShortLink) {
	global.DB.Select("id", "url", "short_url").Create(model)
}
