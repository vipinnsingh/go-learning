package main

import "fmt"

type Database struct {
}

type Storer interface {
	Store(string)
}

func (d Database) Store(data string) {
	fmt.Println(data)
}

type Businesslogic struct {
	Storer Storer
}

func (b Businesslogic) SaveData(data string) {
	b.Storer.Store(data)
}

func main() {

	d := Database{}
	b := Businesslogic{Storer: d}

	b.SaveData("vjhbk")

}
