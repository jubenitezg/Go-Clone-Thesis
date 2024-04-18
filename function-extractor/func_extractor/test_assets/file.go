package test_assets

import "strconv"

type Person struct {
	Name string
	Age  int
}

func (p *Person) SayHello() string {
	return "Hello " + p.Name
}

func (p *Person) SayAge() string {
	return "I am " + strconv.Itoa(p.Age)
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}
