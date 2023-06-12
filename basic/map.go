package main

import (
	"fmt"
)


func main()  {
	response:=map[string]string{
		"status": "500",
		"message": "server error",
	}

	delete(response,"message")

	fmt.Println(response)
	fmt.Println(len(response))

}