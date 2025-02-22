package main

import "fmt"

/**
variadic func is a variabel arguments (varargs) and varargs is an array, 
so i can input the parameter value more than one in a one variabel
*/

// ... is a varargs of numbers, and the position always at the last (right side of param variable)
func sumAll(numbers ... int) int{
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func main(){
	// cuz number is an arguments, so i can put more values in there
	result := sumAll(10,5,10,2,3)

	fmt.Println(result)
}