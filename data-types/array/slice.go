package main

import "fmt"

func main() {
	// slice from array
	var months = [...]string{
		"Januari", "Februari", "Maret", "April", "Mei", "Juni",
		"Juli", "Agustus", "September", "Oktober", "November", "Desember",
	}

	fmt.Println(months)

	var sliceOne = months[4:7] // one as pointer 
	fmt.Println(sliceOne)
	fmt.Println(len(sliceOne)) //get len of array 
	fmt.Println(cap(sliceOne)) // get capacity of array

	sliceOne[0] = "April Lagi" // if array capacity still enough, change value in slice will change value in array
	fmt.Println(sliceOne)
	fmt.Println(months)

	// make slice
	sliceTwo := make([]string, 2, 5) // len 2, cap 5

	sliceTwo[0] = "kale"
	sliceTwo[1] = "morales"

	fmt.Println(sliceTwo)
	fmt.Println(len(sliceTwo))
	fmt.Println(cap(sliceTwo))

	sliceTwo = append(sliceTwo, "joko")
	fmt.Println(sliceTwo)

	// copy slice
	sliceThree := make([]string, len(sliceTwo), cap(sliceTwo))
	copy(sliceThree, sliceTwo) // (destination, source)
	fmt.Println(sliceThree)

	// becarefull when make a slice because if u forget to set len, it will be array
	typeArray := [...]string{"a", "b", "c"}
	typeSlice := []string{"a", "b", "c"}

	fmt.Println(typeArray)
	fmt.Println(typeSlice)
}