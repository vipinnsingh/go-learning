package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	Id string
}

func main() {

	p := Person{Name: "abc", Age: 32}

	emp := Employee{Person: p, Id: "12455342"}

	fmt.Println(emp)
}
