package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var days = []string{
		"mon","tue","wed","thu", "fri", "sat", "sun",
	}

	workdays := days[0:5]
	meetingdays := workdays[2:3]
	meetingdays[0] = "Rakor"

	var fiveFristMonth = []string{
		"jan","feb", "mar", "apr", "may",
	}

	fiveFristMonth[0] = "Happy new year"
	valentineMonth:= fiveFristMonth[1]

	fmt.Println("days",days)
	fmt.Println("workdays",workdays)
	fmt.Println("meetingdays",meetingdays)
	fmt.Println("fiveFristMonth",fiveFristMonth)
	fmt.Println("valentineMonth",reflect.TypeOf(valentineMonth))




}