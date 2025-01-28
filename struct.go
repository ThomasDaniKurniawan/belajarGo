package main

import (
	"fmt"
	"strconv"
)

type Customer struct {
	name, address string
	age           int
}

type loanDataCustomer struct {
	loanAmount float64
	interest   float64
}

func (loanData loanDataCustomer) finalAmount() float64 {
	finalAmount := loanData.loanAmount * (loanData.interest / 100)
	return finalAmount
}

func main() {
	var cus Customer
	cus.name = "Thomas"
	cus.address = "Yogyakarta"
	cus.age = 24
	fmt.Println(cus)

	dataAmount := loanDataCustomer{
		loanAmount: 567000900,
		interest:   5}

	result := strconv.FormatFloat(dataAmount.finalAmount(), 'f', -1, 64)
	fmt.Println(result)
}
