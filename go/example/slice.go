package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr) // [1 2 3 4 5]

	slc := arr[1:3]
	fmt.Println(slc)      // [2 3]
	fmt.Println(len(slc)) // 2
	fmt.Println(cap(slc)) // 4

	fmt.Println(slc[0]) // 2
	fmt.Println(slc[1]) // 3
	// fmt.Println(slc[2]) // Error out of len

	new_slc := slc[:4]   // access to array and create new slice
	fmt.Println(new_slc) // [2 3 4 5]

}
