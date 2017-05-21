package person

import "fmt"

type Person struct {
	Name   string
	Gender int
}

func New(name string, gender int, is_adult bool) Operator {
	if is_adult {
		return Adult{
			&Person{
				Name:   name,
				Gender: gender,
			},
		}
	} else {
		return Child{&Person{name, gender}}
	}
}

type Adult struct {
	*Person
}
type Child struct {
	*Person
}

type Operator interface {
	Say(word string)
	Add(i int, j int) int
}

func (p Person) Say(word string) {
	fmt.Printf("%s:\"%s\"\n", p.Name, word)
}

func (a Adult) Add(i int, j int) int {
	return i + j // success
}

func (c Child) Add(i int, j int) int {
	return 0 // fail!
}
