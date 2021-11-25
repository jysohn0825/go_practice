package practice

import "fmt"

func MapPractice(){
	var m = make(map[int]string)

	m[1] = "서울"
	m[2] = "경기"
	m[3] = "인천"

	fmt.Println(m)

	m[3] = "인천 수정"

	fmt.Println(m)

	delete(m, 3)

	fmt.Println(m)
	
	
	val, exist := m[4] 
	fmt.Println(val, exist) // 없을 경우 비어 있는 값 나옴

	val = m[2] 
	fmt.Println(val)

	_, exist = m[1] 
	fmt.Println(exist)
}