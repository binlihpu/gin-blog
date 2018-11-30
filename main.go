package main

import (
	"fmt"
	"net/http"

	"github.com/binlihpu/gin-blog/pkg/setting"
	"github.com/binlihpu/gin-blog/routers"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerConf.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ServerConf.ReadTimeout,
		WriteTimeout:   setting.ServerConf.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
