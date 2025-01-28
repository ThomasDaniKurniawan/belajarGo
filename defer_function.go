package main
import "fmt"


/*
content : recurtion, defer function

**/

// recurtion
func factorial(number int) int{
	if number == 1{
		return 1
	}else{
		return number * factorial(number-1)
	}

}


// defer function
func logging(){
	fmt.Println("selesai memanggil function")
}

func runApplication(){
	/*
	defer function akan diexecute setelah function yg memanggilnya 
	selesai diexecute, sekalipun error akan tetap dijalankan.
	**/
	defer logging() // logging akan diexecute setelah semua code program di runApplication selesai diexecute.
	fmt.Println("menjalankan aplikasi")
}


func main (){
	
	factorial := factorial(3)
	fmt.Println(factorial)

	runApplication()
}