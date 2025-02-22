package main

import "fmt"

func random() interface{} {
	//return "Nasi Goreng" // test case, if the return value is a string
	//return 120 // test case, if the return value is a number
	return true // if the return value are not string or number
}

func main() {

	result := random()

	switch value := result.(type) {
	case string:
		fmt.Println(value, " is string")
	case int:
		fmt.Println(result, " is numeric")
	default:
		fmt.Println(result, "is other type")
	}
}
