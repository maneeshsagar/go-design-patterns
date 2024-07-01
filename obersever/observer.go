package main

import "fmt"

// observer is somthing that oberves in certain thing if it finds anything then noties to the subscribers

type Observer interface {
	GetId() int
	Update()
}
type Person struct {
	Id   int
	Name string
}

func (p *Person) GetId() int {
	return p.Id
}
func (p *Person) Update() {
	fmt.Println("update from ", p.Id)
}

func NewPerson(id int, name string) Observer {
	return &Person{
		Id:   id,
		Name: name,
	}
}
