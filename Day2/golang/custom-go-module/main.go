package main

import (
    "fmt"
    "addition"
    "subtraction"
)

func main() {
     //We are casting/converting float64 into float32
     x := float32(100.123)

     //We are casting/converting float64 into float32
     y := float32(200.456)

     fmt.Println ( "The sum of ", x, " and ", y, " is ", addition.Add( x, y ) )
     fmt.Println ( "The difference of ", x, " and ", y, " is ", subtraction.Subtract( x, y ) )
}
