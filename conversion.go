package main

import "fmt"

func main() {
	var value32 int32 = 100000
	var value64 int64 = int64(value32) // konversi hanya dengan menyebutkan tipe data tujuan "int64"
	var value8 int8 = int8(value32) // value terlalu besar untuk int8 (128), akan terjadi overflow

	fmt.Println("value32 =", value32)
	fmt.Println("value64 =", value64)
	fmt.Println("value8 =", value8)

	var name = "kale morales"
	var k = name[0]
	var kstring = string(k) // konversi byte ke string

	fmt.Println("name =", name)
	fmt.Println("k =", k)
	fmt.Println("kstring =", kstring)
}