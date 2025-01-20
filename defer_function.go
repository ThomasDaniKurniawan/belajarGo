package main
import "fmt"


/*
content : recurtion defer function, panic, recover

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
	defer function ini akan diexecute setelah function yg memanggilnya 
	selesai diexecute, sekalipun error akan tetap dijalankan.
	**/
	defer logging()
	fmt.Println("menjalankan aplikasi")
}


func main (){
	
	factorial := factorial(3)
	fmt.Println(factorial)

	runApplication()
}