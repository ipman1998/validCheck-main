package fileHandle

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(filePath string, arrayString *[]string) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		*arrayString = append(*arrayString, scanner.Text())
	}

	return scanner.Err()
}
