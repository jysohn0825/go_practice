package chapter

import "fmt"

func Change_type_pratice(){
	var num int = 10
	var changeFloat float32 = float32(num) //int형을 float32형으로 변환
	changeInt := int8(num)               //int형을 int8형으로 변환

	var str string = "konai"
	changeStr := []byte(str) //바이트 배열
	str2 := string(changeStr) //바이트 배열을 다시 문자열로 변환

	fmt.Println(num)
	fmt.Println(changeFloat, changeInt)

	fmt.Println(str)
	fmt.Println(changeStr)
	fmt.Println(str2)
	
	// float = int / int 불가능! 모든 형을 동일하게 가져가야 함
}