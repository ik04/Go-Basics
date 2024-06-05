package main

import "fmt"

func main(){
	menu := map[string]float64{
		"soup":90.99,
		"pie":88.00,
	}
	// fmt.Println(menu["soup"])

	for k,v := range menu{
		fmt.Println(k,"-",v)
	}
}