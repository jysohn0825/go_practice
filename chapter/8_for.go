package chapter

import "fmt"

func For_practice(){
	sum := 0
	
	for i := 1; i <= 10; i++{
		sum += i
	}
	fmt.Println("1부터 10까지의 합 : ",sum)
	
	multi := 1
	
	for multi < 100 {
		multi *= 2
	}
	fmt.Println("100보다 작은 2의 배수 중 가장 큰 수 : ", multi)
}