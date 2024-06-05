// * return multiple values
package main

import (
	"fmt"
	"strings"
)

func getInitials(n string)(string,string){
	s :=strings.ToUpper(n)
	names := strings.Split(s, " ")
	var initials []string 
	for _,v := range names{
		initials = append(initials, v[:1] )
	}
	if len(initials) > 1{
		return initials[0], initials[1]
	}
	return initials[0],"_"

}

func main(){
	name := "ishaan khurana"
	fn1, sn1 := getInitials(name)
	fmt.Println(fn1,sn1)
}