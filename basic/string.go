package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(len("dimas"))
    fmt.Println(strings.Compare("a", "a"))
    fmt.Println(strings.Compare("a", "z"))
    fmt.Println(strings.Compare("s", "c"))
}