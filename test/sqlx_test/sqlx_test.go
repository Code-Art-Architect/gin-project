package sqlx_test

import (
	"fmt"
	"testing"

	"github.com/code-art/gin-project/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	var err error
	if db, err = util.InitDataBaseWithSqlX(); err != nil {
		fmt.Printf("Connected Failed Err: %v\n", err)
	}
}

type user struct {
	Id uint64
	Name string
	Age uint32
}

func TestQuerySingleRow(t *testing.T) {
	sqlStr := "select * from `tb_student` where id = ?"
	var u user
	if err := db.Get(&u, sqlStr, 1); err != nil {
		fmt.Printf("Query Failed Err: %v\n", err)
	}
	fmt.Println(u)
}

func TestUpdateRow(t *testing.T) {
	sqlStr := "update `tb_student` set age = ? where id = ?"
	res, err := db.Exec(sqlStr, 100, 2)
	if err != nil {
		fmt.Printf("Updated Failed Err: %v\n", err)
	}
	
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected Failed Err: %v\n", err)
	}
	
	fmt.Println("n = ", n)
}

func TestInsertRow(t *testing.T) {
	sqlStr := "insert into `tb_student` values(?, ?, ?)"
	res, err := db.Exec(sqlStr, nil, "莱布尼茨", 200)
	if err != nil {
		fmt.Printf("Updated Failed Err: %v\n", err)
	}
	
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected Failed Err: %v\n", err)
	}
	
	fmt.Println("Inserted Successfully : ", n)
}

func TestDeleteRow(t *testing.T) {
	sqlStr := "delete from `tb_student` where id = ?"
	res, err := db.Exec(sqlStr, 6)
	if err != nil {
		fmt.Printf("Updated Failed Err: %v\n", err)
	}

	n, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("Get RowsAffected Failed Err: %v\n", err)
	}
	
	fmt.Println("n = ", n)
}

func TestQueryName(t *testing.T) {
	sqlStr := "select * from `tb_student` where age = :age"
	rows, err := db.NamedQuery(sqlStr, map[string]interface{} {
		"age": 200,
	})
	
	if err != nil {
		fmt.Printf("Name Query Failed Err: %v\n", err)
	}
	
	defer rows.Close()
	
	for rows.Next() {
		var u user
		if err := rows.StructScan(&u); err != nil {
			fmt.Printf("Strduct Scan Failed Err: %v\n", err)
		}
		fmt.Println(u)
	}
}

func TestBatchInsert(t *testing.T) {
	users := []user{
		{Name: "111", Age: 100},
		{Name: "222", Age: 150},
		{Name: "333", Age: 200},
	}
	
	sqlStr := "insert into `tb_student`(name, age) values(:name, :age)"
	_, err := db.NamedExec(sqlStr, users)
	if err != nil {
		fmt.Printf("Batch Insert Failed Err: %v\n", err)
	}
	fmt.Println("Success")
}