package chapter_two

import "fmt"

func Calculate_additional_practice() {
	ch := make(chan int) //정수형 채널 생성

	go func() {
		ch <- 10
	}() //채널에 10을 보냄

	result := <-ch //채널로부터 10을 전달받음
	fmt.Println(result)
}

func Pointer_practice(){
	var num int = 5
	var pnum = &num

	fmt.Println("num : ", num)   //num 값
	fmt.Println("pnum :", pnum)  //num의 메모리 주소
	fmt.Println("pnum :", *pnum) //num의 주소로 메모리에 할당돼있는 값 접근

	*pnum++
	fmt.Println("num : ", num)
	fmt.Println("pnum :", *pnum)
}