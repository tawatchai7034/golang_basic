package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func UnmarshalEmployee(data []byte) (Employee, error) {
	var r Employee
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Employee) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

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

type ApiResponse struct {
	Code    int64  `json:"Code"`
	Message string `json:"Message"`
	Result  int64  `json:"Result"`
}

func getEmployee(w http.ResponseWriter, db *sql.DB) {
	result, err := db.Query(`SELECT * FROM tb_employee`)
	if err != nil {
		fmt.Println(err)
	}
	var tableDataList []Employee
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
		tableDataList = append(tableDataList, user)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	empJson, err := json.Marshal(tableDataList)
	w.Write(empJson)
}

func getSingleEmployee(w http.ResponseWriter, db *sql.DB, id int) {
	result, err := db.Query(`SELECT * FROM tb_employee WHERE id = ?`, id)
	if err != nil {
		fmt.Println(err)
	}
	var tableDataList []Employee
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
		tableDataList = append(tableDataList, user)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	empJson, err := json.Marshal(tableDataList)
	w.Write(empJson)
}

func addEmployee(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var newItem Employee
	bodybtye, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bodybtye, &newItem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := db.Exec(`INSERT INTO tb_employee (EmpNum,EmpName,HireDate,Salary,Position,DepNo,HeadNo) VALUE(?,?,?,?,?,?,?)`, newItem.EmpNum, newItem.EmpName, newItem.HireDate, newItem.Salary, newItem.Position, newItem.DepNo, newItem.HeadNo)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	} else {
		id, err := result.LastInsertId()
		if err != nil {
			fmt.Println(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		resJson := ApiResponse{Code: 200, Message: "success", Result: id}
		empJson, err := json.Marshal(resJson)
		w.Write(empJson)
	}
}

func updateEmployee(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var newItem Employee
	bodybtye, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bodybtye, &newItem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := db.Exec(`UPDATE tb_employee 
							SET EmpNum = ?,EmpName=?,HireDate=?,Salary=?,Position=?,DepNo=?,HeadNo=?
							WHERE id = ?`, newItem.EmpNum, newItem.EmpName, newItem.HireDate, newItem.Salary, newItem.Position, newItem.DepNo, newItem.HeadNo, newItem.ID)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	} else {
		_, err := result.LastInsertId()
		if err != nil {
			fmt.Println(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		resJson := ApiResponse{Code: 200, Message: "success", Result: newItem.ID}
		empJson, err := json.Marshal(resJson)
		w.Write(empJson)
	}
}

func deleteEmployee(w http.ResponseWriter, db *sql.DB, id int) {
	_, err := db.Query(`DELETE FROM tb_employee WHERE id = ?`, id)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	} else {

		if err != nil {
			fmt.Println(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		resJson := ApiResponse{Code: 200, Message: "success", Result: 0}
		empJson, _ := json.Marshal(resJson)
		w.Write(empJson)
	}
}

func empHandler(w http.ResponseWriter, r *http.Request) {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/localdb?parseTime=true")
	if err != nil {
		fmt.Println("Database connect fail")
	} else {
		fmt.Println("Database connect success")
		defer db.Close()
		switch r.Method {
		case http.MethodGet:
			getEmployee(w, db)
		case http.MethodPost:
			addEmployee(w, r, db)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func empSingle(w http.ResponseWriter, r *http.Request) {
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/localdb?parseTime=true")
	if err != nil {
		fmt.Println("Database connect fail")
	} else {
		fmt.Println("Database connect success")
		defer db.Close()
		pathList := strings.Split(r.URL.Path[1:], "/")
		id, err := strconv.Atoi(pathList[len(pathList)-1])
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		switch r.Method {
		case http.MethodGet:
			getSingleEmployee(w, db, id)
		case http.MethodPut:
			updateEmployee(w, r, db)
		case http.MethodDelete:
			deleteEmployee(w, db, id)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("brfore handler middle ater")
		handler.ServeHTTP(w, r)
		fmt.Println("middleware finised")
	})
}

func main() {
	empNoparam := http.HandlerFunc(empHandler)
	empGetparam := http.HandlerFunc(empSingle)
	// no param
	http.Handle("/employee", middleware(empNoparam))
	// get param
	http.Handle("/employee/", middleware(empGetparam))
	http.ListenAndServe(":5000", nil)
}
