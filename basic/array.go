package main

import "fmt"

func main() {
  var user [3]string 
	user[0] = "Adi"
	user[1] = "Ani"
	user[2] = "Budi"


	var kota = [3]string{
		"sby", "bkl", "sda",
	}
	fmt.Println(kota)

	var(
		typeStatus = [3]string{
			"todo", "progress", "done",
		}
	
	)

	fmt.Println(typeStatus)


	var animal = [...]string{
		"dog", "cat", "cow", "bug", "fish", "bird",
	}

	slice := animal[:3]

	slice2 := append(slice,"spider1")
	slice3 := append(slice,"spider2")
	slice4 := append(slice,"spider3")



	fmt.Println(animal)
	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice)




}