package model

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/victor-leee/portal-be/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func MustInit(cfg *config.Config) {
	ormDB, err := gorm.Open(mysql.Open("%s:%s@tcp(%s:%d)/portal?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		logrus.Fatal(err)
	}
	db = ormDB
}

func GetMysql(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}
