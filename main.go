package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liu-shilong/go-gin-demo/pkg/setting"
	"github.com/liu-shilong/go-gin-demo/routers"
	"net/http"
)

func main() {
	gin.SetMode(setting.RunMode)
	routersInit := routers.InitRouter()

	s := http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        routersInit,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		return
	}
}
