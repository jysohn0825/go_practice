package practice

import "fmt"

func SlicePractice(){
	// slice는 참조 타입이기 때문에 슬라이스를 복사한다 = 같은 주소를 참조한다
	// make(슬라이스 타입, 슬라이스 길이, 슬라이스의 용량) : e.g. 30인승 버스 = 용량, 탑승 인원 17명 = 길이
	
	c := make([]int, 0, 3) //용량이 3이고 길이가0인 정수형 슬라이스 선언
	c = append(c, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(c), cap(c))

	l := c[1:3] //인덱스 1요소부터 2요소까지 복사
	fmt.Println(l)

	l = c[2:] //인덱스 2요소부터 끝까지 복사
	fmt.Println(l)

	l[0] = 6

	fmt.Println(c) //슬라이스 l의 값을 바꿨는데 c의 값도 바뀜
	//값을 복사해온 것이 아니라 기존 슬라이스 주솟값을 참조
}
