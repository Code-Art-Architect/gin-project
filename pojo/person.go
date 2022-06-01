package pojo

import "errors"

type person struct {
	Name string
	age  int
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	if age > 0 && age < 100 {
		p.age = age
	} else {
		_ = errors.New("年龄不合法!")
	}
}

func (p *person) GetAge() int {
	return p.age
}
