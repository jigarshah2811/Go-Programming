package main

import (
	"io/ioutil"
	"os"
	"strings"
	"fmt"
)

func main()  {
	if len(os.Args) != 3 {
		panic("Usage: fileIO.go <filename1> <filename2>")
	}
	filename1 := os.Args[1]
	filename2 := os.Args[2]
	fileIOExamples(filename1, filename2)
}

func fileIOExamples(filename1 string, filename2 string)  {

	fileBytes1, err := ioutil.ReadFile(filename1)
	if err != nil {
		panic("File1 read failed!")
	}

	fileBytes2, err := ioutil.ReadFile(filename2)
	if err != nil {
		panic("File2 read failed!")
	}

	originalData := string(fileBytes1)
	words := strings.Split(originalData, " ")

	fmt.Printf("Fun text location:")
	for index, word := range words {
		if strings.Compare(word, "my") == 0 {
			fmt.Printf("\t%d", index)
		}
	}
	fmt.Println()

	replacedData := strings.Replace(originalData, "my", "your", 5)

	combinedText := string(replacedData) + string(fileBytes2)
	err = ioutil.WriteFile("combined.txt", []byte(combinedText), 0644)
	if err != nil {
		panic("CombinedFile Write failed!")
	}
}