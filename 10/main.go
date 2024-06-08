package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string,r *bufio.Reader)(string,error){
	fmt.Println(prompt)
	input,err := r.ReadString('\n')
	return strings.TrimSpace(input),err
}

func createBill() bill {
	reader:= bufio.NewReader(os.Stdin)
	name,_ := getInput("Create new Bill name: ",reader)
	name = strings.TrimSpace(name)
	b := newBill(name)
	fmt.Println("Created bill for", name)
	return b
}

func main(){
	mybill := createBill()
	mybill.updateTip(10.22)
	mybill.addItem("ishaan",10.99)
	summary := mybill.format()
	fmt.Println(summary)
	mybill.save()

}