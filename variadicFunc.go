package main

import "fmt"

/**
variadic func is a variabel arguments, 
so i can input the parameter value more than one in a one variabel
*/

// ... is a varargs of numbers, and the position always at the last (right side)
func sumAll(numbers ... int) int{
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func main(){
	// cuz number is an argument, so i can input more values in there
	result := sumAll(10,5,10,2,3)

	fmt.Println(result)
}