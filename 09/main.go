package main

import "fmt"

// pointers
func updateName(n *string){
	*n = "strap"

}

func main(){
	name := "tifa"
	updateName(&name)
	fmt.Println(&name)
	fmt.Println(name)
}