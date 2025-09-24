package main

import "fmt"

func main() {

        // index              0   1   2   3   4   5
	intArray := [6]int { 10, 20, 30, 40, 50, 60 } 

	fmt.Println ("Array elements are ...")
	fmt.Println(intArray)

	//30,40,50 are pointed by this slice
	var mySlice []int = intArray[2:5] // 2 is the lower bound index, while 5 is the upper bound index, 5 is not inclussive

	fmt.Println("Slice elements are ...")
	fmt.Println(mySlice)

	// Let's modify the slice at certain indices
	// when the slice is modified, it will reflect in the original array that is referred by slice
	mySlice[0] = 100 //mySlice[0] is nothing but intArray[2]
	mySlice[1] = 200 //mySlice[1] is nothing but intArray[3]
	mySlice[2] = 300 //mySlice[2] is nothing but intArray[4]

	fmt.Println("Slice elements after modifying are ...")
	fmt.Println(mySlice)

	fmt.Println("Array elements after modifying slice are ...")
	fmt.Println(intArray)

	mySlice = append(mySlice, 400)
	fmt.Println("Array elements after appending 400 in slice are ...")
	//intArray will be {10,20,100,200,300,400} at this point
	fmt.Println(intArray)

	mySlice = append(mySlice, 500)
	fmt.Println("Array elements after appending 500 in slice are ...")
	//intArray will be {10,20,100,200,300,400} at this point
	//mySlice will be {100,200,300,400,500} at this point the mySlice is no more associated with intArray
	//mySlice will be pointing to a fresh array of int with 5 integer elements as slice can grow(resize) at run-time on demand
	fmt.Println(intArray)

	fmt.Println("Slice elements after appending 500 into slice are ...")
	fmt.Println(mySlice)

}
