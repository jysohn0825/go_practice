package chapter

import "fmt"

func For_range_practice(){
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