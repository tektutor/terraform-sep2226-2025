package main

import "fmt"

func main() {
	x := 100

	if x % 2 == 0 {
	    fmt.Println(x, "is an even number")
	} else {
            fmt.Println(x, "is an odd number")
	}
}
