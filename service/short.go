package service

import (
	"context"
	"fmt"
	"short-link-go/dao"
	"short-link-go/global"
	"short-link-go/models"
	"short-link-go/utils"
)

// GetShortUrl 生成并记录短链数据
func GetShortUrl(oriUrl string) (shortUrl string) {
	url := dao.GetShortByUrl(oriUrl)
	if url.ShortUrl != "" {
		return global.CONF.App.ShortHost + url.ShortUrl
	}
	genShortUrl := utils.GenShortUrl(oriUrl)
	link := models.ShortLink{ShortUrl: genShortUrl, Url: oriUrl}
	dao.SaveShortUrl(&link)
	redisInsertBF(genShortUrl)
	return global.CONF.App.ShortHost + genShortUrl
}

// 插入布隆过滤器
func redisInsertBF(genShortUrl string) {
	ctx := context.Background()
	inserted, err := global.Redis.Do(ctx, "BF.ADD", "bf_key", genShortUrl).Bool()
	if err != nil {
		panic(err)
	}
	if inserted {
		fmt.Println("当前已经存在于布隆过滤器")
	}
}

// 判定是否命中布隆过滤器
func redisIsExitBF(genShortUrl string) bool {
	ctx := context.Background()
	exists, err := global.Redis.Do(ctx, "BF.EXISTS", "bf_key", genShortUrl).Bool()
	if err != nil {
		panic(err)
	}
	return exists
}

// Short 获取原始链接
func Short(shortUrl string) (oriUrl string) {
	exists := redisIsExitBF(shortUrl)
	if !exists {
		global.LOGGER.Info("过滤后未找到链接")
		return global.CONF.App.DefaultUrl
	}
	url := dao.Short(shortUrl)
	if url.Url != "" {
		return url.Url
	}
	return global.CONF.App.DefaultUrl
}
