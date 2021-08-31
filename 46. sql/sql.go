package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	// _ "github.com/lib/pq"
)

type student struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:aditya0724@tcp(localhost:3360)/golang_tutorial?charset=utf8&parseTime=True&loc=Local")
	// db, err := sql.Open("postgres", "user=postgres dbname=golang_tutorial password=aditya0724 host=localhost sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery(age int) {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// var age = 27
	// rows, err := db.Query(`select id, name, grade from tb_student where age = $?`, age)
	rows, err := db.Query("select id, name, grade from tb_student where age = ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []student

	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.name)
	}
}

func main() {
	sqlQuery(27) //age = 27
}
