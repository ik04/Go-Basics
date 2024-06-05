package main

import "fmt"

// loops

func main(){
	arr := []int{1,2,3,4,5}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	for index, value := range arr{
		fmt.Println(index,value)
	}
}