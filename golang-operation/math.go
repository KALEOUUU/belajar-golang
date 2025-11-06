package main

import "fmt"

func main() {
	a := 10
	b := 3

	// penjumlahan
	sum := a + b
	fmt.Println("Penjumlahan:", sum)

	// pengurangan
	diff := a - b
	fmt.Println("Pengurangan:", diff)

	// perkalian
	product := a * b
	fmt.Println("Perkalian:", product)

	// pembagian
	quotient := a / b
	fmt.Println("Pembagian:", quotient)

	// sisa bagi
	remainder := a % b
	fmt.Println("Sisa Bagi:", remainder)

	sum += 5
	fmt.Println("Setelah Penjumlahan dengan 5:", sum)

	sum -= 2
	fmt.Println("Setelah Pengurangan dengan 2:", sum)

	sum *= 3
	fmt.Println("Setelah Perkalian dengan 3:", sum)

	sum /= 4
	fmt.Println("Setelah Pembagian dengan 4:", sum)

	sum %= 3
	fmt.Println("Setelah Sisa Bagi dengan 3:", sum)

	sum ++
	fmt.Println("Setelah Increment:", sum)
}