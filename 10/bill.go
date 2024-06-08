package main

import (
	"fmt"
	"os"
)

type bill struct{
	name string
	items map[string]float64
	tip float64
}

func newBill(name string)bill{
	b:= bill{
		name: name,
		items: map[string]float64{
			"pie":2.99,
			"cake":3.00,
		},
		tip: 0,
	}
	return b
}

// * reciever function, only be called from bill objects
func (b bill) format() string{
	fs:="Bill Breakdown \n" // * formatted string
	var total float64 = 0
	for k,v := range b.items{
		fs += fmt.Sprintf("%-25v ...$%v \n",k+":",v)
		total += v
	}
	if b.tip != 0{
		fs += fmt.Sprintf("%-25v ...$%0.2f \n","Tip:",b.tip)
		total += b.tip
	}
	fs += fmt.Sprintf("%-25v ...$%0.2f","Total:",total)
	return fs
}

func (b *bill) updateTip(tip float64){
	(*b).tip = tip
}

func (b *bill) addItem(name string,price float64){
	(*b).items[name] = price
}
func (b bill) save(){
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt",data,0644)
	if err != nil{
		panic(err)
	}
	fmt.Println("Bill saved!")
}
