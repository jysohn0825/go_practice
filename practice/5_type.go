package practice

import "fmt"

func TypePractice(){
	var str string
	str = `"check raw string and size"\n`
	fmt.Println(str)
}

func TypeChangePractice(){
	var num int = 10
	var changeFloat float32 = float32(num) 
	changeInt := int8(num)

	var str string = "konai"
	changeStr := []byte(str)
	str2 := string(changeStr)

	fmt.Println(num)
	fmt.Println(changeFloat, changeInt)

	fmt.Println(str)
	fmt.Println(changeStr)
	fmt.Println(str2)
	
	// float = int / int 불가능! 모든 형을 동일하게 가져가야 함
}