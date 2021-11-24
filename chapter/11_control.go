package chapter

import "fmt"

func Control_practice(){
	
	i := 0
	
	BREAK_TEST:
		for i < 10{
			i++
			if i == 1{
				fmt.Println(i)
				break BREAK_TEST
			}
		}
}