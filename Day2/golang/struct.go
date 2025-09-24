package main

import "fmt"

type Rectangle struct {
  length int
  width  int
}

//Method - these are functions that can only be invoked via struct variables
//Area is member function of Rectangle structure
//Area function takes zero arguments and it return an integer
func ( rect Rectangle ) Area() int {
   area := rect.length * rect.width
   return area
}

func ( rect Rectangle ) GetLength() int {
   return rect.length
}

func ( rect Rectangle ) GetWidth() int {
   return rect.width
}

func main() {
     rectangle := Rectangle {
        length: 100,
	width : 200,
     }

     fmt.Printf("Length of rectangle : %d\n", rectangle.GetLength() )
     fmt.Printf("Width  of rectangle : %d\n", rectangle.GetWidth () )
     fmt.Printf("Area of rectangle   : %d\n", rectangle.Area() )
}
