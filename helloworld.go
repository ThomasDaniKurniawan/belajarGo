package main

import "fmt"

func main() {
	fmt.Println("hello world!")

	// initiate value and operator
	a := 10
	b := 10
	var c = a + b

	fmt.Println("hasil =", c)

	// boolean
	nilaiAkhir := 90
	nilaiAbsensi := 85

	lulusNilaiAkhir := nilaiAkhir > 80
	lulusNilaiAbsensi := nilaiAbsensi > 80

	lulus := lulusNilaiAbsensi && lulusNilaiAkhir

	if lulus == true {
		fmt.Println("Complete")
	} else {
		fmt.Println("Fail")
	}

	// array
	var values = [3]int{
		10,
		20,
		30,
	}

	var dinamicLen = [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println("array = ", values)
	fmt.Println("array len = ", dinamicLen)

	dinamicLen[3] = 7
	fmt.Println("array len = ", dinamicLen)

	// tipe data slices
	foods := [...]string{"bakso", "Tahu", "tempe", "lotek", "sate", "bakwan", "coto makasar", "nastar"}
	slice1 := foods[4:6]
	// slice2 := foods[6:]

	slice1[1] = "gudek"
	foods2 := append(slice1, "papeda")

	fmt.Println("slice1 ", slice1)
	fmt.Println("foods ", foods)
	fmt.Println("foods2 ", foods2)



	province := make([]string, 2, 5)  // len =2 , max= 5
	province[0] = "Papua"
	province[1] = "Maluku"

	fmt.Println(province)
	fmt.Println(len(province))
	fmt.Println(cap(province))

	province2 := append(province, "Makassar")

	fmt.Println(province2)
	fmt.Println(len(province2))
	fmt.Println(cap(province2))

	province2[0] = "Yogyakarta"

	fmt.Println(province)
	fmt.Println(province2)


	fromSlices := province2[:]
	toSlices := make([]string, 1, 5)

	copy(toSlices, fromSlices)

	fmt.Println("Copy")
	fmt.Println(toSlices)
	fmt.Println(fromSlices)
}
