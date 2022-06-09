package test

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnect(t *testing.T) {
	const driverName = "mysql"
	const dataSourceName = "root:root1234@tcp(localhost:3306)/gin"

	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// 尝试建立连接
	if err := db.Ping(); err != nil {
		fmt.Println("Connected Failed!")
		panic(err)
	}

	// 设置最大连接和闲置连接数目
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("Connected Successfully!")
}
