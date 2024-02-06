/*
This program looks in a specified directory and finds all files with a certain extension. It finds and replaces the specified words.
It is a rewrite of a program of the same name I wrote in python
Developed by: Sage Hourihan
Date: 2/6/24
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var path, textToSearch, textToReplace string
	fmt.Print("File path: ")
	fmt.Scan(&path)
	fmt.Print("Text to search for: ")
	fmt.Scan(&textToSearch)
	fmt.Print("Text to replace it with: ")
	fmt.Scan(&textToReplace)

	wd, err := os.Getwd()
	if err != nil {
		println(err)
	}

	if path != wd {
		os.Chdir(path)
	}

	startTime := time.Now()

	data, err := os.ReadFile(path)

	lines := strings.Split(string(data), "\n")

	for i, line := range lines {
		if strings.Contains(line, textToSearch) {
			// println(lines[i])
			lines[i] = strings.ReplaceAll(line, textToSearch, textToReplace)
		}
	}

	output := strings.Join(lines, "\n")
	err = os.WriteFile(path, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
	endTime := time.Now()

	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("It took %s to run", elapsedTime)
}
