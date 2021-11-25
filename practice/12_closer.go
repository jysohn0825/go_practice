package practice

import "fmt"

func calCoin() func(coin int, num int) int { 
	return func(coin int, num int) int  { //클로저
		return coin * num
	}
}
/* 
입력값
1
4
5
6
출력값
3710
*/
func CloserPractice(){
	
	var coin10, coin50, coin100, coin500 int
	fmt.Scan(&coin10, &coin50, &coin100, &coin500)
	
	calculate := calCoin()
	
	add10 := calculate(10, coin10)
	add50 := calculate(50, coin50)
	add100 := calculate(100, coin100)
	add500 := calculate(500, coin500)
	
	totalmoney := add10 + add50 + add100 + add500
	
	fmt.Println(totalmoney)
}