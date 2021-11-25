package practice

import "fmt"

func ForPractice(){
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

func ForRangePractice(){
	// forEach와 유사한 개념
	var arr [5]int = [5]int{1,2,3,4,5}
	for index, num := range arr{
		fmt.Println(index,"번째 값 ", num)
	}

	fmt.Println()
	
	for _, num := range arr{ //index 주기 싫을 때
		fmt.Println(num)
	}

	fmt.Println()
	
	var fruits map[string]string = map[string]string{
		"apple":  "red",
		"banana": "yellow",
		"grape":  "purple",
	}

	for fruit, color := range fruits {
		fmt.Printf("%s의 색깔은 %s입니다.\n", fruit, color)
	}
}