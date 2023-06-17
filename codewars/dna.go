package main

import (
	"fmt"	
)

func main() {
	fmt.Println(DNAStrand("TTTT"))
	fmt.Println(DNAStrand("TAACG"))
	fmt.Println(DNAStrand("CATA"))


}

func DNAStrand(dna string) string {
	convert := ""
	for _, char := range dna {
		switch char {
			case 'A':
				convert += "T"
			case 'T':
				convert += "A"
			case 'C':
				convert += "G"
			case 'G':
				convert += "C"
		}
	}

	return convert

}
