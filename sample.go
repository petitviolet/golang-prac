package main

import (
	"fmt"
	"os"

	p "net.petitviolet/prac/person"
)

func ops(i int, j int) int {
	person := p.New("alice", 2, false)
	person.Say("hello")
	return person.Add(i, j)
}

func show(target []string) (string, error) {
	var res string
	len := len(target)
	if len > 1 {
		fmt.Printf("target: %s\n", target)
		return fmt.Sprintf("target: \n %s \n %v \n %+v \n %#v", target, target, target, target), nil
	} else {
		return res, fmt.Errorf("argument is too short...")
	}
}

func main() {
	res, err := show(os.Args)
	if err != nil {
		fmt.Printf("error!!! %#v\n", err)
	} else {
		fmt.Printf("result => %s\n", res)
	}

	fmt.Println("-----------------------")

	fmt.Printf("ops: %d", ops(1, 2))
}
