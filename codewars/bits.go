package main

import (
	"fmt"	
)

func main() {
	fmt.Println(CountBits(0))
	fmt.Println(CountBits(4))
	fmt.Println(CountBits(7))
	fmt.Println(CountBits(9))
	fmt.Println(CountBits(10))



}

func CountBits(number uint ) int {
	bits := 0

	for number > 0 {
		if number % 2 == 1{
			bits++
		}  
		number = number / 2
	}

	return bits

}