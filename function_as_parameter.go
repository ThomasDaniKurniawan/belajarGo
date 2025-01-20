package main

import "fmt"

type Filter func(string,int) string 

func sayHeloWithFilter(name string, age int, filter Filter){
	filterName := filter(name, age)
	fmt.Println ("welcome, ",filterName)
}

func spamFilter (name string, age int) string{
	if name == "student"{
		if age >= 18 {
			return "your order in process"
		}else{
			return "your age enough yet"
		}
	}else{
		if age < 24 {
			return "young teacher"
		}
		return name
	}
}


/*

 why called anonymos function?,
 cuz the declaration of function name is general. 

usually we declare function with name, 
tapi kadang ada kalanya lebih mudah membuat function secara langsung di variable atau parameter 
	tanpa harus membuat function terlebih dahulu.
	Hal tersebut dinamakan anonymous function, atau function tanpa nama 
**/

// declare function as a data type is named Blacklist
type Blacklist func(string) bool 

func registerUser(name string, blacklist Blacklist){
	if blacklist(name){
		fmt.Println("you're blocked", name)
	}else{
		fmt.Println("welcome ",name)
	}
}


func main(){
	
	// func as parameter
	sayHeloWithFilter("student", 18, spamFilter)
	filter := spamFilter
	sayHeloWithFilter("teacher", 20, filter)

	// anonymous function
	blacklist := func(name string) bool {
		return name == "Bobi"
	}

	registerUser("Susi", blacklist)




}