package main

import (
	"fmt"
	"os"
	"encoding/json"
)
type quotes struct{
	quotes []quotes `json:"quotes"`
}

func Readfile(){
	fmt.Println("Reading file ....")
	filename := "quotes.json"
	jsonFile, err0r := os.Open(filename)
	if err0r != nil{
		println("Damn error")
	}



	fmt.Println("Filename: "+filename)
	fmt.Println(data)

}

func main(){
	Readfile()
}