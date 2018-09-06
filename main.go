package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var user, password string
	fmt.Printf("\nEnter user name (skip for 'root' ) : ")
	fmt.Scanf("%s", &user)
	if user == "" {
		user = "root"
	}
	fmt.Printf("\nEnter password : ")
	fmt.Scanf("%s", &password)
	url := fmt.Sprint(user, ":", password, "@/information_schema")
	doStuff(url)
}

func doStuff(url string) {
	db, _ := sql.Open("mysql", url)
	defer db.Close()
	stmt, _ := db.Prepare("SELECT COLUMN_NAME FROM COLUMNS WHERE TABLE_SCHEMA=? AND TABLE_NAME=?")
	defer stmt.Close()
	fmt.Print("Enter Table Name ( e.g. mydb.mytable) : ")
	var name string
	fmt.Scan(&name)
	var names = strings.Split(name, ".")
	rows, _ := stmt.Query(names[0], names[1])
	defer rows.Close()
	var cols, q []string
	var col string
	var i = 0
	for rows.Next() {
		rows.Scan(&col)
		cols = append(cols, col)
		q = append(q, "?")
		i = i + 1
	}
	fmt.Println(fmt.Sprint("INSERT INTO ", names[0], ".", names[1], "(", strings.Join(cols, ","), ")VALUES(", strings.Join(q, ","), ")"))

}
