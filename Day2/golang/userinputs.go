package main

import "fmt"

func main() {
	//The below line declares variables named x and y and initializes them with value 0
	x := 0

	//This is yet another way we can declare variables in go lang
	var y int

	//This is an ansignment
	y = 0

	fmt.Print("Enter your first integer input :" )
	fmt.Scanf("%d", &x)

	fmt.Print("Enter your second integer input :" )
	fmt.Scanf("%d", &y)

	fmt.Println("Value of x :", x )
	fmt.Println("Value of y :", y )

	var temp string
	fmt.Println("Press any key to exit ...")
	fmt.Scanln(&temp)
}
