package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	ID       int64     `json:"id"`
	EmpNum   string    `json:"EmpNum"`
	EmpName  string    `json:"EmpName"`
	HireDate time.Time `json:"HireDate"`
	Salary   float64   `json:"Salary"`
	Position string    `json:"Position"`
	DepNo    string    `json:"DepNo"`
	HeadNo   string    `json:"HeadNo"`
}

func main() {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/localdb?parseTime=true")

	// if there is an error opening the connection, handle it

	if err != nil {
		fmt.Println("Database connect fail")
	} else {
		fmt.Println("Database connect success")
		defer db.Close()
		// var dataList []Employee
		// dataList = getUser(db)
		// if len(dataList) > 0 {
		// 	for i := 0; i < len(dataList); i++ {
		// 		data, err := json.Marshal(&dataList[i])
		// 		if err != nil {
		// 			panic(err)
		// 		}
		// 		fmt.Println(string(data))
		// 	}
		// }

		fmt.Println(getUser(db))
	}
}

func getUser(db *sql.DB) []Employee {
	result, err := db.Query(`SELECT * FROM tb_employee`)
	if err != nil {
		fmt.Println(err)
	}
	var userDataList []Employee
	for result.Next() {
		var user Employee
		err := result.Scan(
			&user.ID,
			&user.EmpNum,
			&user.EmpName,
			&user.HireDate,
			&user.Salary,
			&user.Position,
			&user.DepNo,
			&user.HeadNo,
		)

		if err != nil {
			panic(err.Error())
		}
		userDataList = append(userDataList, user)
	}
	return userDataList
}

func addUser(db *sql.DB) bool {
	currentTime := time.Now()
	result, err := db.Exec(`INSERT INTO tb_employee (EmpNum,EmpName,HireDate,Salary,Position,DepNo,HeadNo) VALUE(?,?,?,?,?,?,?)`, "tr", "eerr", currentTime, 50000, "", "", "")
	if err != nil {
		fmt.Println(err)
	} else {
		id, err := result.LastInsertId()
		if err != nil {
			fmt.Println(err.Error())
			return false
		}
		fmt.Println("item id:", id)
	}
	return true
}

func deleteUser(db *sql.DB) bool {
	_, err := db.Exec(`DELETE FROM tb_employee WHERE id = ?`, 22)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
