package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/victor-leee/portal-be/internal/handler"
)

func main() {
	logrus.Info("starting service")
	r := gin.Default()
	r.POST("/create-service", handler.CreateService)
	logrus.Fatal(r.Run())
}