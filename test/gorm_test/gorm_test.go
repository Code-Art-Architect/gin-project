package gorm_test

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func init() {
	// 配置全局logger
	newLog := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Info,
		},
	)
	
	dsn := "root:root1234@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLog,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected Successfully!")
}

func TestMigrate(t *testing.T) {
	// 迁移 schema
	db.AutoMigrate(&Product{})
}

func TestCreate(t *testing.T) {
	// Create
	db.Create(&Product{Code: "D42", Price: 100})
}

func TestFindFirst(t *testing.T) {
	// Read
	var product Product
	db.First(&product, 1) // 根据整型主键查找
	fmt.Println(product)
	
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	fmt.Println(product)
}

func TestUpdate(t *testing.T) {
	var product Product
	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
}

func TestDelete(t *testing.T) {
	var product Product
	// Delete - 删除 product
	db.Delete(&product, 1)
}


