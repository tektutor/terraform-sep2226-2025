package main

import "fmt"

func main() {

	//Declares a count variable of type int and assigns a value 5
	count := 5

	//similar to while loop
	for count > 0 {
		fmt.Println("Before decrementing ", count)

		count--  //equivalent to count = count - 1
		//golang doesn't support pre-decrement or pre-increment unlike C or C++

		fmt.Println("After decrementing ", count )
	}
	fmt.Println("Value of count is ", count, " after for loop")

	//Resetting already declared variable count with value 0
	count = 0

	fmt.Println()
	//Regular for loop
	for count=1; count<10; count++ {
           fmt.Printf( "%d\t", count )
	}
	fmt.Println()

	count = 0

	//similar to do while - infinite loop
	for {

		fmt.Printf("Inside for loop %d\n", count )
		count++

		if count > 3 {
		   break
		}
	}
	fmt.Println("Control reached outside infinite for loop")
}
