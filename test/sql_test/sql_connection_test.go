package sql_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/code-art/gin-project/util"
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

// 测试封装 SQL 连接
func TestSqlUtil(t *testing.T) {
	if _, err := util.InitDataBase(); err != nil {
		panic(err)
	}

	fmt.Println("Connected Successfully!")
}
