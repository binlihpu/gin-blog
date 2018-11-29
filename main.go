package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/binlihpu/gin-blog/pkg/setting"
)

func main() {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerConf.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ServerConf.ReadTimeout,
		WriteTimeout:   setting.ServerConf.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
