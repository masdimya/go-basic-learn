package main

import (
	"fmt"
)

func main()  {
	var status int16 = 500

	if status == 200{
		fmt.Println("Success")
	} else {
		fmt.Println("Server Error")
	}

	switch{
	case status == 200:
		fmt.Println("Success")
	case status == 500:
		fmt.Println("Server Error")
	}




}