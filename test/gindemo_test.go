package main

import "fmt"

func (t Test) name() {
	fmt.Println(t.name)
}

func (t Test) age() {
	fmt.Println(t.age)
}

func (t Test) sex() {
	fmt.Println(t.sex)
}
