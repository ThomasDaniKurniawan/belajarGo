package main

import "fmt"


// panic
func endApp(){
	fmt.Println("end app")

	/*
	recover ini dijalankan sebelum code program panic diexecute. fungsinya adalah utk merecover program 
	agar tetap dijalankan setelah terjadi panic.
	**/
	message := recover()
	fmt.Println("terjadi panic ",message)
	
}

func runApp(error bool){
	defer endApp()  // krn di defer, endApp akan tetap dijalankan meskipun panic dijalankan. 
	if error{
		/*
		panic function akan menghentikan program jika kondisi terpenuhi (in case: true).
		tapi defer function akan tetap dijalankan.
		**/
		panic("THERE is an ERROR")
	}
	
}

func main (){
	
	runApp(true)
	fmt.Println("logg ...")
}