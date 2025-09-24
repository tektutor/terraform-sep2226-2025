package main

import "fmt"

func sayHello ( msgPtr *string ) {

	//Dereferencing - the values stored at address pointed by msgPtr will be printed here
	fmt.Println( "Inside sayHello funciton", *msgPtr )

	//Here the address pointed by msgPtr pointer will be printed
	fmt.Println("Address pointed by msgPtr is", msgPtr )

	//Print the address of msgPtr
	fmt.Println("Address of msgPtr is", &msgPtr )

	//The values stored at the address pointed by msgPtr is assigned to tmp string
	tmp := *msgPtr

	//We are changing the value stored at address pointed by msgPtr pointer
	*msgPtr = tmp + " Golang" + " !"

	fmt.Println("Inside sayHello before returning ", *msgPtr)

}

func main() {
   //declares a string variable name str with value "Hello"
   msg := "Hello"

   fmt.Println("Message before calling sayHello function is ", msg )
   fmt.Println("Address of msg string is ", &msg )

   sayHello( &msg )

   fmt.Println("Message after calling sayHello function is ", msg )
}
