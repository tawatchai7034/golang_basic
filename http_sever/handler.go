package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	ID    int64  `json:"Id"`
	Name  string `json:"Name"`
	Tel   string `json:"Tel"`
	Email string `json:"Email"`
}

var empList []Employee

func initData() {
	empJson := `[
	{"Id":12,"Name":"fluke","Tel":"0611804157","Email":"example@gmail.com"},
	{"Id":20,"Name":"test","Tel":"0954568411","Email":"example@gmail.com"},
	{"Id":14,"Name":"book","Tel":"061485244","Email":"example@gmail.com"}
	]`

	err := json.Unmarshal([]byte(empJson), &empList)
	if err != nil {
		log.Fatal(err)
	}
}

func getNextId() int {
	id := -1
	for _, item := range empList {
		if id < int(item.ID) {
			id = int(item.ID)
		}
	}
	return id + 1
}

func empHandler(w http.ResponseWriter, r *http.Request) {
	empJson, err := json.Marshal(empList)

	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(empJson)
	case http.MethodPost:
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
		if newItem.ID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newItem.ID = int64(getNextId())
		empList = append(empList, newItem)
		w.WriteHeader(http.StatusCreated)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
func fineId(id int) (*Employee, int) {
	for i, item := range empList {
		if id == int(item.ID) {
			return &item, i
		}
	}
	return nil, 0
}

func empSingle(w http.ResponseWriter, r *http.Request) {
	pathList := strings.Split(r.URL.Path[1:], "/")
	id, err := strconv.Atoi(pathList[len(pathList)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	e, index := fineId(id)
	if e == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{code:200,message:"ID not found"}`))
		return
	} else {
		switch r.Method {
		case http.MethodGet:
			res, err := json.Marshal(e)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
		case http.MethodPut:
			var updateItem Employee
			bodybtye, err := ioutil.ReadAll(r.Body)
			fmt.Println(index)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			err = json.Unmarshal(bodybtye, &updateItem)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if updateItem.ID != e.ID {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			e = &updateItem
			empList[index] = *e
			// w.WriteHeader(http.StatusOK)
			res, err := json.Marshal(empList[index])
			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
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
	initData()
	empNoparam := http.HandlerFunc(empHandler)
	empGetparam := http.HandlerFunc(empSingle)
	// no param
	http.Handle("/employee", middleware(empNoparam))
	// get param
	http.Handle("/employee/", middleware(empGetparam))
	http.ListenAndServe(":5000", nil)
}
