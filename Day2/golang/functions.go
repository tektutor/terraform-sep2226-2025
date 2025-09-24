package main

import "fmt"

func yetAnotherFunction() {
     fmt.Println("Yet Another Function invoked")
}

func main() {
	fmt.Println( sayHello("Golang") )
	fmt.Println( sayHello("World") )
	yetAnotherFunction()
}

//This function accepts a string input argument and returns a string output
func sayHello( msg string ) string {
  return "Hello, " + msg + " !"
}

/* function overloading is not supported in golang
func sayHello() string {
   return "Hello World !"
}
*/
