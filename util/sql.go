package util

import "database/sql"

var db *sql.DB

func InitDataBase() error {
	const driverName = "mysql"
	const dataSourceName = "root:root1234@tcp(localhost:3306)/gin"

	db, err := sql.Open(driverName, dataSourceName)
	
	if err != nil {
		return err
	}
	
	if err := db.Ping(); err != nil {
		return err
	}
	
	return nil
}
