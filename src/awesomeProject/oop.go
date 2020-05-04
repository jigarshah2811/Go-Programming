package main

/* How do you define something in Go?
// Keyword (some)identifier type
// keyword: https://golang.org/ref/spec#Keywords
// identifier: https://golang.org/ref/spec#Identifiers
// type: https://golang.org/ref/spec#Types
*/

type human struct {
    name string
}

type UserDB struct {
	users []User
}

type UserDBOperations interface {
	addUser()
	removeUser()
}

func () addUser()  {
	
}

func main()  {
	EncapsulationExample()
}

func EncapsulationExample()  {
    fmt.Println("EncapsulationExample")
}