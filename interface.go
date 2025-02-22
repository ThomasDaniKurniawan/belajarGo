package main

import "fmt"

/*
interface is a class of abstarcts.
in this code, i want to :
	1. make an interface,
	2. use a struct at the interface.

note :
di golang, setiap tipe data yg sesuai dgn kontrak interface,
secara otomatis akan dianggap sebagai interface itu sendiri.
sedangakan,
dalam bahasa pemrogaman lain (eg: java) harus menggunakan \
implemets utk mendeklareasikan interface itu.

**/

// declare the interface
type hasName interface {
	getName() string
	getAge() int
}

// declare the struct of person
type Person struct {
	name string
	age  int
}

// declare kontrak untuk struct person
func (person Person) getName() string {
	return person.name
}
func (person Person) getAge() int {
	return person.age
}

// declare struct of animal
type Animal struct {
	name string
	age  int
}

// declare kontrak untuk struct animal
func (animal Animal) getName() string {
	return animal.name
}
func (animal Animal) getAge() int {
	return animal.age
}

func sayHello(value hasName) {
	fmt.Println("hello, ", value.getName())
	fmt.Println("your age is ", value.getAge())
}

func guessAnimal(value hasName) {
	fmt.Println("this is a ", value.getName())
	fmt.Println("the age is ", value.getAge())
}

func main() {
	person := Person{"Budi", 23}
	sayHello(person)

	animal := Animal{"Babi Hutan", 2}
	guessAnimal(animal)

}
