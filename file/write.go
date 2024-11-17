package main

import "os"

func main() {
	// update file
	data1 := []byte("Golang write file testing55555")
	err := os.WriteFile("D:/Learning/golang_basic/file/write.txt", data1, 0644)
	if err != nil {
		panic(err)
	}

	// create file
	f, f_err := os.Create("D:/Learning/golang_basic/file/name.txt")
	if f_err != nil {
		panic(f_err)
	}
	defer f.Close()

	data2 := []byte("Golang create file testing")
	os.WriteFile("D:/Learning/golang_basic/file/name.txt", data2, 0644)
}
