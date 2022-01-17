package util

import (
	"log"
	"os"
)

var (
	InfoLogger      *log.Logger
	Status200Logger *log.Logger
)

func Init() {
	fileStatus200, err := os.OpenFile("result.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	infoFile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(infoFile, "", 0)
	Status200Logger = log.New(fileStatus200, "", 0)
}
