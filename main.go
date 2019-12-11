package main

import (
	"net/http"
	"os"
	"time"

	"jiazhen/common"
	"jiazhen/config"
	"jiazhen/models"
	"jiazhen/routers"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})

	//写入文件
	src, err := os.OpenFile("log/run.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	log.SetOutput(src)
	log.SetLevel(log.DebugLevel)
}

func main() {
	db := common.Init()
	db.AutoMigrate(&models.Brand{}, &models.Category{}, &models.Shifu{}, &models.Typ{}, &models.FeedBack{})
	s := &http.Server{
		Addr:           ":" + config.ServerConf.Port,
		Handler:        routers.Routers(),
		ReadTimeout:    120 * time.Second,
		WriteTimeout:   120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
