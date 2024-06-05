package main

import "fmt"

func sayHi(n string){
fmt.Println("hello" , n)
}

func cycleNames(names []string, greet func(string)){
	for _,v := range names{
		greet(v)
	}
}

func main(){
	// sayHi("ishaan")
	names := []string{"ishaan","jai","dhruv","stanley","saira"}
	cycleNames(names,sayHi)

}