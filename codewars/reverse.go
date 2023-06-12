package main

import (
	"fmt"
)

func main()  {
	var city string = ".god yzal eht revo spmuj xof nworb kciuq ehT"
	
	fmt.Println(ReverseWords(city))

}

func ReverseWords(str string) string {
  reverse := ""
  
  for i := len(str) - 1; i >= 0; i-- {
	  reverse += string(str[i])  
  }
  
  return reverse // reverse those words
}