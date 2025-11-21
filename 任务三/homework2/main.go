package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Employees struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}
type Books struct {
	ID     int    `db:"id"`
	Title  string `db:"title"`
	Author string `db:"author"`
	Price  int    `db:"price"`
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/go_demo?charset=utf8mb4&parseTime=True"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	//测试连接
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("连接成功")

	// sqlx.createTable(db.MapperName("users"))
	getDtail(db)
	getDtail2(db)
}
func getDtail(db *sqlx.DB) {
	employees := Employees{}
	// var employees Employees
	var employeesList []Employees
	query := "select * from employees where department = ?"
	query1 := "select * from employees order by salary desc limit 1"
	db.Get(&employees, query, "技术部")
	db.Select(&employeesList, query1)
	fmt.Printf("%#v\n", employees)
	fmt.Printf("%+v\n", employeesList)

}
func getDtail2(db *sqlx.DB) {
	var books []Books
	query := "select * from books where price >50 order by price desc"
	db.Select(&books, query)
	fmt.Printf("%+v\n", books)

}
