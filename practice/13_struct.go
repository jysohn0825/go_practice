package practice

import "fmt"

type person struct {
	name    string
	age     int
	mobile string
}

func newPerson() *person{
	p := person{}
	p.name = "초기화"
	p.age = 999
	p.mobile = "01012345678"
	return &p
}


func addAgeRef(a *person) {
	a.age += 4 //자동 역참조 * 생략
}

func addAgeVal(a person) { //Pass by value
	a.age += 4
}


func StrcutPractice(){
	// var p1 = new(person) //포인터 구조체 객체 생성
	p1 := newPerson()    // 빈 구조체 객체 생성
	var p2 = person{}
	
	fmt.Println(p1, p2)
	
	p1.age = 25
	p2.age = 25

	addAgeRef(p1) //&을 쓰지 않아도 됨
	addAgeVal(p2)

	fmt.Println(p1)
	fmt.Println(p2)

	addAgeRef(&p2) //&을 붙이면 pass by reference 가능
	fmt.Println(p2)

}