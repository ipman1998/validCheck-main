package main

import (
	"log"
	"os"
	check "validCheck/checkStatus"
	fHandle "validCheck/fileHandle"
	"validCheck/util"
	"fmt"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	util.Init()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("URI")
	url := "https://gist.githubusercontent.com/am1ru1/275af8ab41b12d72b94bd18ba779c51b/raw/e43c672905132bea1d138b581081dc59d69035c1/log4j%2520exploit%2520payload%2520samples.txt"
	arrayString := make([]string, 0)
	if err := fHandle.DownloadFile("file", url); err != nil {
		log.Panicln("Cannot download file, problem is: ", err.Error())
	}
	fHandle.ReadFile("file/tempFile.txt", &arrayString)
	dt := time.Now().UTC().Format("2006-01-02 15:04:05")
	totalCase := 0
	errorCase := 0

	fmt.Printf(
	`********Start*********Time: %s
	`, dt)
	
	util.Status200Logger.Printf(
	`********Start*********Time: %s
	`, dt)
	
	for _, jndi := range arrayString {
		total, err :=check.CheckStatus(uri, jndi)
		totalCase = totalCase + total
		errorCase = errorCase + err
	}
	fmt.Printf(
	`********Final Result*********
	Error: %d
	Successful: %d
	Total: %d
	
	`, errorCase, (totalCase - errorCase), totalCase)
	
	util.Status200Logger.Printf(
	`********Final Result*********
	Error: %d
	Successful: %d
	Total: %d
	
	`, errorCase, (totalCase - errorCase), totalCase)
}
