package util

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func InitDataBase() (db *sql.DB, err error) {
	const driverName = "mysql"
	const dataSourceName = "root:root1234@tcp(localhost:3306)/gin"

	db, err = sql.Open(driverName, dataSourceName)
	
	if err != nil {
		return nil, err
	}
	
	if err := db.Ping(); err != nil {
		return nil, err
	}
	
	return db, nil
}

func InitDataBaseWithSqlX() (db *sqlx.DB, err error) {
	const dataSourceName = "root:root1234@tcp(localhost:3306)/gin"
	db, err = sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		fmt.Printf("Connected Failed: %v\n", err)
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	fmt.Println("Connected Successfully!")
	
	return db, nil 
}
