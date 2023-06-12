package main

import (
    "fmt"
		"strconv"
)

func main() {
  // var message string
  // var name string

	// message = "Selamat Pagi"
	// name = "Dimas"

	// fmt.Println(message)
	// fmt.Println(name)

	// var age int8 = 26

	// fmt.Println(age)

	// address := "Bangkalan"
	
	// fmt.Println(address)

	// var height int8 = 100

	// fmt.Println(height)


	var (
		firstName = "Dimas"
		LastName = "Aulia"
		ageNow = 26
	)

	var height, weight int8 = 100, 75


	fmt.Println(firstName,LastName,ageNow)
	fmt.Println(height,weight)


	var (
		total int
		discount int = 10
		message string = "hello"
	)

	fmt.Println(total,discount, message)


	const (
		phi float32 = 3.14
		STATUS_ID string = "1"
	)

	const ORDER_ID, CANCEL_ID string = "OD-111", "CAN-111"

	fmt.Println(ORDER_ID,CANCEL_ID)

	var nilai = "90"
	nilaiInt, err := strconv.Atoi(nilai)
	fmt.Println(nilaiInt, err)

	fmt.Println(firstName[0], string(firstName[0]))

 
}