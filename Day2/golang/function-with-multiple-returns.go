package main

import "fmt"

func myFunction() (int,int) {
  return 10, 20
}

func main() {
   x, y := myFunction() // := is a short form of declaring a new variable and initialized with some value

   fmt.Println( "Value of x is ", x )
   fmt.Println( "Value of y is ", y )
}
