package chapter

import "fmt"

func Var_practice(){
	var a string = "string init"
	fmt.Println(a)
	
	b := "string check"
	fmt.Println(b)
	
	c,d,e := "multi", "string", "check"
	fmt.Println(c,d,e)
}