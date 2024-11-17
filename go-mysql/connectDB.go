package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	ID       int64     `json:"id"`
	EmpNum   string    `json:"EmpNum"`
	EmpName  string    `json:"EmpName"`
	HireDate time.Time `json:"HireDate"`
	Salary   int64     `json:"Salary"`
	Position string    `json:"Position"`
	DepNo    string    `json:"DepNo"`
	HeadNo   string    `json:"HeadNo"`
}

// func query(db *sql.DB) {
// 	var tag Employee
// 	queryStr := "SELECT id,EmpNum FROM tb_employee"
// 	if err := db.QueryRow(queryStr, 1).Scan(&tag.ID, &tag.EmpName); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(tag)
// }

func main() {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:root(127.0.0.1:3306)/localdb")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT id,EmpNum FROM tb_employee")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tag Employee
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.ID, &tag.EmpName)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.Printf(tag.EmpName)
	}

}
