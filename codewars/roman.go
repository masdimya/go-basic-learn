
/*
codewars challenge : codewars.com/r/5eykug

*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	
	// fmt.Println(195,ConvertRoman(195))
	// fmt.Println(216,Solution(216))
	// fmt.Println(373,Solution(373))
	// fmt.Println(480,Solution(480))
	// fmt.Println(599,Solution(599))
	// fmt.Println(610,Solution(610))
	fmt.Println(2476,Solution(2476))





}

func Solution(number int) string {
	var roman string

	digit:=len(strconv.Itoa(number))	

	switch digit{
		case 1:
			roman = RomanOneDigit(number)
		case 2:
			roman = RomanTwoDigit(number)
		case 3:
			roman = RomanThreeDigit(number)
		case 4:
			roman = RomanFourDigit(number)
	}	

  return roman
}

func RomanOneDigit(number int) string {
	if number == 5{
		return "V"
	} else if number == 9{
		return "IX"
	} else if number == 4{
		return "IV"
	} else if number > 5{
		return "V" + strings.Repeat("I", number - 5)
	} else {
		return strings.Repeat("I", number)
	}
}

func RomanTwoDigit(number int) string {
	if number % 10 == 0 && number < 40 {
		return strings.Repeat("X", number / 10) 
	} else if number == 40{
		return "XL"
	} else if number == 50{
		return "L"
	}	else if number == 90{
		return "XC"
	} else if number % 10 == 0 && number > 50{
		return "L" +  strings.Repeat("X", (number - 50 ) / 10 ) 
	} else {
		unit:= number % 10
		doz := number - unit
		return RomanTwoDigit(doz) + RomanOneDigit(unit)
	}	
}

func RomanThreeDigit(number int) string {
	if number % 100 == 0 && number < 400 {
		return strings.Repeat("C", number / 100) 
	} else if number == 400{
		return "CD"
	} else if number == 500{
		return "D"
	}	else if number == 900{
		return "CM"
	} else if number % 100 == 0 && number > 500{
		return "D" +  strings.Repeat("C", (number - 500 ) / 100 ) 
	} else {
		doz := number % 100
		hundreds:= number - doz 
		return RomanThreeDigit(hundreds) + RomanTwoDigit(doz) 
	} 
}

func RomanFourDigit(number int) string {
	
	if number % 1000 == 0 && number < 4000 {
		return strings.Repeat("M", number / 1000) 
	} else {
		hundreds := number % 1000
		thousands := number - hundreds
		return RomanFourDigit(thousands) + RomanThreeDigit(hundreds) 
	} 
}