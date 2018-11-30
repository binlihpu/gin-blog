package main

import (
	"fmt"
	"log"
	"syscall"

	"github.com/binlihpu/gin-blog/pkg/setting"
	"github.com/binlihpu/gin-blog/routers"
	"github.com/fvbock/endless"
)

func main() {
	endless.DefaultReadTimeOut = setting.ServerConf.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerConf.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerConf.HTTPPort)
	server := endless.NewServer(endPoint, routers.InitRouter())

	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
