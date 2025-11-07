package main

import "fmt"

func main() {
	name1 := "holang"
	name2 := "golang"

	var result bool = name1 == name2 // bandingkan apakah name1 sama dengan name2
	fmt.Println(result)

	var result2 bool = name1 != name2 // bandingkan apakah name1 tidak sama dengan name2
	fmt.Println(result2)

	// "<, >" jarang digunakan untuk membandingkan string di Go
	var result3 bool = name1 > name2 // bandingkan apakah name1 kurang dari name2
	fmt.Println(result3)

	value1 := 10
	value2 := 20

	var result4 bool = value1 < value2 // bandingkan apakah value1 kurang dari value2
	fmt.Println(result4)

	var result5 bool = value1 >= value2 // bandingkan apakah value1 lebih besar atau sama dengan value2
	fmt.Println(result5)

	var result6 bool = value1 <= value2 // bandingkan apakah value1 kurang dari atau sama dengan value2
	fmt.Println(result6)

	var result7 bool = value1 == value2 // bandingkan apakah value1 sama dengan value2
	fmt.Println(result7)

	var result8 bool = value1 != value2 // bandingkan apakah value1 tidak sama dengan value2
	fmt.Println(result8)
}