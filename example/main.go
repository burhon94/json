package main

import (
	"github.com/burhon94/json/cmd/reader"
	"github.com/burhon94/json/cmd/writer"
	"log"
)

type exStruct struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

var exData = exStruct{
	Id:   1,
	Name: "Vasya",
}

func main() {
	write, err := writer.JsonWrite(exData)
	if err != nil {
		panic(err)
	}
	log.Print(string(write))

	fileUpload, err := writer.JsonFileUpload("./example")
	if err != nil {
		panic(err)
	}
	log.Print(fileUpload)

	err = reader.ReadJSON(write, &exStruct{})
	if err != nil {
		panic(err)
	}

}
