package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"short-link-go/models"
	"short-link-go/response"
	"short-link-go/service"
	"short-link-go/utils"
)

func GetShort(c *gin.Context) {
	shortLink := models.ShortLink{}
	err := c.ShouldBind(&shortLink)
	if err != nil {
		utils.HandleValidatorError(c, err)
		return
	}
	response.Success(c, 0, "", service.GetShortUrl(shortLink.Url))
}

func Short(c *gin.Context) {
	path := c.Param("path")
	url := service.Short(path)
	//302 重定向
	c.Redirect(http.StatusFound, url)
}
func DefaultUrl(c *gin.Context) {
	response.Success(c, 10, "无法解析当前url", nil)
}
