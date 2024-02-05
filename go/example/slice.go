package main

import "fmt"


func main(){
	arr := [5]int{1,2,3,4,5}
	fmt.Println(arr)

	slc := arr[1:3]
	fmt.Println(len(slc))
	fmt.Println(cap(slc))
}
