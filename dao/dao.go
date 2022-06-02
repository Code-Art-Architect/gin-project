package dao

import (
	"gin-project/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Manager interface {
	AddUser(user *model.User)
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to init db:", err)
	}

	Mgr = &manager{db: db}
	_ = db.AutoMigrate(&model.User{})
}

func (m manager) AddUser(user *model.User) {
	m.db.Create(user)
}
