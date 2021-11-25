package practice

import "fmt"

const ( 
	c1 = 1
	c2
	c3
	c4 = "four"
	c5
	c6 = iota
	c7
	c8
	c9 = 9
	c10
	c11 = "End"
)

func ConstPractice() {
	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11)
}