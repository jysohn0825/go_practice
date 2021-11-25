package practice

import "fmt"

func passByValue(a int){
	a *= a
	fmt.Println(a)
}

func passByReference(a *int){
	*a *= *a
	fmt.Println(*a)
}

func checkMutable(num ...int){ // 가변 매개 변수 인자
	var sum int
	for i := 0; i < len(num); i++{
		sum += num[i]
	}
	fmt.Println(sum)
}

func checkNamedRetun(a int) (num int){ //반환 변수 명을 미리 지정할 수 있으며, 반환 값이 하나더라도 괄호를 적용해야 한다
	num = a*a
	return
}

//함수 원형 정의
type calculatorNum func(int, int) int 

func calNum(f calculatorNum, a int, b int) (result int) {
	result = f(a, b)
	return
}


func FuncPractice(){
	a := 4
	
	passByValue(a)
	
	passByReference(&a)
	
	checkMutable(a, a+1, a+2, a+3, a+4)
	
	fmt.Println(checkNamedRetun(a))
	
	multi := func(i int, j int) int {
		return i * j
	}
	fmt.Println(calNum(multi, 3, 4))
	
}