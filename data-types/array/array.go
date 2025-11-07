// array[index]

package main

import "fmt"

func main() {

	var names [3]string // deklarasi array dengan tipe data string dan panjang 3

	names[0] = "kale"
	names[1] = "morales"
	names[2] = "zheng"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	fmt.Println(len(names)) // untuk mengambil panjang array

	var number = [2]int{
		10,
		20,
	}

	fmt.Println(number)
	fmt.Println(len(number)) 
	fmt.Println(number[0])
	fmt.Println(number[1])

	class := [3]string{ // bentuk lain deklarasi array
		"math",
		"english",
		"computer science",
	}

	fmt.Println(class)
	fmt.Println(len(class))
	fmt.Println(class[0])
	fmt.Println(class[1])
	fmt.Println(class[2])

}