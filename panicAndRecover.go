package main
import "fmt"


// panic
func endApp(){
	fmt.Println("end app")

	/*
	recover ini dijalankan sebelum panic
	**/
	message := recover()
	fmt.Println("terjadi panic ",message)
	
}

func runApp(error bool){
	//defer endApp()  // krn di defer, endApp akan tetap dijalankan meskipun panic dijalankan. 
	if error{
		/*
		panic function akan menghentikan program jika kondisi terpenuhi (in case: true)
		**/
		panic("THERE is an ERROR")
	}
	defer endApp() 
}

func main (){
	
	runApp(true)
	fmt.Println("logg ...")
}