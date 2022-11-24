package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	content := "Hemlo bro ! Good morning dffdf\n\tsdsds"

	file, err := os.Create("./sample.txt")
	if err != nil {
		panic(err)
	}

	length, err :=  io.WriteString(file, content)
	if err != nil {
		panic(err)
	}else{
		println("File created with length: ", length)
	}
	file.Close()
	readFile("./sample.txt")
}


func readFile(filename string){
	databytes, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(databytes)
	fmt.Println(string(databytes))
}