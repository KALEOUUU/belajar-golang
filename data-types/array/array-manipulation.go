// array manipulation == "array[index] = value"

package main

import "fmt"

func main() {
	var num = [4]int{
		10,
		20,
		30,
		40,
	}

	fmt.Println("Sebelum diubah:", num)

	num[2] = 100 // manipulation array
	num[3] = 70
	fmt.Println("Setelah diubah:", num)
}