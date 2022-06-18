package test_gorm

import (
	"fmt"
	"testing"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
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

func TestBatchInsert(t *testing.T) {
	var products = []Product{
		{Code: "100", Price: 200},
		{Code: "200", Price: 300},
		{Code: "300", Price: 400},
	}
	
	db.Create(products)
}




