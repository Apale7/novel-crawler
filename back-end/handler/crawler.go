package handler

import (
	"net/http"
	"novel-crawler-back-end/crawler"
	"time"

	"github.com/gin-gonic/gin"
)

func GetList(c *gin.Context) {
	startTime := time.Now().UnixNano()

	url := c.DefaultQuery("url", "")
	name, list, err := crawler.GetList(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "获取小说列表失败"})
		return
	}
	endTime := time.Now().UnixNano()

	c.JSON(http.StatusOK, gin.H{
		"msg":  "获取小说列表成功",
		"list": list, "cost": float64(endTime-startTime) / 1e9,
		"name": name,
	})
}

func GetNovel(c *gin.Context) {
	startTime := time.Now().UnixNano()
	url := c.DefaultQuery("url", "")
	err := crawler.GetNovel(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "爬取小说失败，请重试"})
		return
	}
	endTime := time.Now().UnixNano()
	c.JSON(http.StatusOK, gin.H{"msg": "爬取小说成功", "cost": float64(endTime-startTime) / 1e9})
}
