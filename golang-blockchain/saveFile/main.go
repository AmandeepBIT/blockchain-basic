package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Printf("AmandeepS")
	content := "Amandeep Singh is just saving the file here !!!!!! ep Singh is just saving the file here !!!!!! ep Singh is just saving the file here !!!!!! ep Singh is just saving the file here !!!!!!"
	filename := "./demo.txt"
	file, err := os.Create(filename)
	checkErr(err)
	length, err := io.WriteString(file, content)
	checkErr(err)
	fmt.Println("file length is ", length)
	defer file.Close()
	readFile(filename)
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkErr(err)

	fmt.Println("Raw data of saved file is ", databyte)
	fmt.Println("Actual data of saved file is ", string(databyte))

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
