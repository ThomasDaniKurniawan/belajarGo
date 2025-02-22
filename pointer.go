package main

import "fmt"

/*

digolang jika menginisiasi suatu variabel (person2) dgn variabel lainnya (person1), dia akan menduplikat value.
jadi jika person2 diupdate person1 tidak akan berubah. ini disebut pass by value.
eg:
 person1 := Person{"Budi", "Adi", "Victor"}
 person2 := person1

nah, bagaimana jika person1 ikut keubah jika person2 diubah (pass by reference) ?
gunakan pointer
**/

type Person struct {
	name, hobby, ethnic string
	age, height         int
}

func main() {

	//person1 := Person{"Aldi", "fishing", "manado", 26, 173} // var person1 Person = Person{var1, var2, ..., varn}
	//person2 := &person1 // var person2 *Person = &person1

	// tau bisa juga menggunakan operatir new utk menginisiasi objek pointer 
	person1 := new(Person)
	person2 := person1

	person2.name = "Devan"
	fmt.Println(person1)
	fmt.Println(person2)

	/*
		operator asterik (*), akan mengganti value dari objek yg dia reference.
		maka semua variabel yg mengacu pada objek Person akan keganti juga valuenya (e.g. person1).

		**/
	*person2 = Person{"Dodi", "cooking", "Dayak", 29, 175}
	fmt.Println(person1)
	fmt.Println(person2)
}
