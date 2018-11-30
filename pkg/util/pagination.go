package util

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"

	"github.com/binlihpu/gin-blog/pkg/setting"
)

// GetPage GetPage
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.AppConf.PageSize
	}

	return result
}
