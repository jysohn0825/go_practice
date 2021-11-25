package practice

import "fmt"

func panicAndRecoverTest() {
	defer func() {
		r := recover() //복구 및 에러 메시지 초기화
		fmt.Println(r)
	}()
	
    var a = [4]int{1,2,3,4}
    
    for i := 0; i < 10; i++ {
        fmt.Println(a[i])
    }       
}

func DeferPractice(){
	panicAndRecoverTest()
		fmt.Println("Hello World!")
}