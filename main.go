package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	Tom := Person{"G", 34}
	bob := Tom
	bob.age = 23
	fmt.Println(Tom, bob)
	inc(&Tom)
	fmt.Println(Tom, bob)
	fmt.Println("now its time")
}

func inc(person *Person) {
	person.age++
}
