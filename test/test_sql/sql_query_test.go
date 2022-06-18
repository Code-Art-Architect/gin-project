package sql

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/code-art/gin-project/util"
)

var db *sql.DB

func init() {
	var err error
	if db, err = util.InitDataBase(); err != nil {
		panic(err)
	}

	fmt.Println("Connected Successfully!")
}

type user struct {
	Id uint64
	Name string
	Age uint32
}

func TestQuerySingleRow(t *testing.T) {
	const sqlStr = "select * from `tb_student` where id = ?"
	var u user

	if err := db.QueryRow(sqlStr, 1).Scan(&u.Id, &u.Name, &u.Age); err != nil {
		log.Printf("Scan Failed Error : %v\n", err)
	}
	log.Println(u)
}

func TestQueryMultiRow(t *testing.T) {
	const sqlStr = "select * from `tb_student`"
	rows, err := db.Query(sqlStr)
	
	if err != nil {
		log.Println(err)
	}
	
	defer rows.Close()
	users := make([]user, 0)
	
	for rows.Next() {
		var u user
		err := rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			log.Println(err)
		}
		users = append(users, u)
	}
	
	fmt.Println(users)
}

func TestUpdateRow(t *testing.T) {
	const sqlStr = "update `tb_student` set name = ? where id = ?"
	res, err := db.Exec(sqlStr, "张三", 1)
	if err != nil {
		fmt.Printf("Updated Failed! %v\n", err)
	}
	
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("get Rows Affected Failed %v\n", err)
	}

	fmt.Println(n)
	fmt.Println("Updated Success!")
}

func TestDeleteRow(t *testing.T) {
	const sqlStr = "delete from `tb_student` where id = ?"
	res, err := db.Exec(sqlStr, 5)
	if err != nil {
		fmt.Printf("Deleted Failed err: %v\n", err)
	}
	
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Get Rows Affected Failed %v\n", err)
	}
	
	fmt.Println(n)
	fmt.Println("Deleted Success!")
}

func TestInsertRow(t *testing.T) {
	const sqlStr = "insert into `tb_student` (name, age) values(?, ?)"
	res, err := db.Exec(sqlStr, "洛必达", 10)
	if err != nil {
		fmt.Printf("Inserted Failed Err: %v\n", err)
	}
	
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("Get LastInserted Faild Err: %v\n", err)
	}
	
	fmt.Println(id)
}

// 测试mysql的预处理
func TestPreparedStatement(t *testing.T) {
	const sqlStr = "select * from `tb_student` where name = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	
	rows, _ := stmt.Query("张三")
	defer rows.Close()
	
	users := make([]user, 0)
	
	for rows.Next() {
		var u user
		err := rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			log.Println(err)
		}
		users = append(users, u)
	}
	
	fmt.Println(users)
}