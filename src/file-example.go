package main

import (
	"io/ioutil"
	"log"
)

func main()  {
	s := "Hello File!"

	err := ioutil.WriteFile("hello.txt", []byte(s), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
