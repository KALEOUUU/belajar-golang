// type declaration untuk mendefinisikan tipe data kustom

package main

import "fmt"

func main() {
	type NoKTP string // mendefinisikan tipe data kustom "NoKTP" dengan tipe dasar string
	type Married bool // mendefinisikan tipe data kustom "Married" dengan tipe dasar bool

	var marriedStatus Married = false
	fmt.Println("Married Status =", marriedStatus)

	var ktpKale NoKTP = "317xxxxxxxxxxxx"
	fmt.Println("No KTP Kale =", ktpKale)


}