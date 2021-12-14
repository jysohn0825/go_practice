package practice

import (
	"fmt"
	"time"
)

func syncPractice(){
	done := make(chan bool)

	go func() {
		for i := 0; i < 4; i++ {
			done <- true

			fmt.Println("고루틴 : ", i)
		}
	}()

	for i := 0; i < 4; i++ {
		<-done
		
		fmt.Println("메인 함수 : ", i)
		
		time.Sleep(time.Second) //눈에 보이게 하기 위해 넣음
	}	
}


func asyncPractice(){
	done := make(chan bool, 2)

	go func() {
		for i := 0; i < 4; i++ {
			done <- true

			fmt.Println("고루틴 : ", i)
		}
	}()

	for i := 0; i < 4; i++ {
		<-done                    
		
		fmt.Println("메인 함수 : ", i)
	}	
}

func ChannelPractice(){
	fmt.Println("동기 시작")
	syncPractice()
	fmt.Println("\n\n비동기 시작")
	asyncPractice()
}