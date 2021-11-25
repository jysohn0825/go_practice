package practice

import(
	"fmt"
	"log"
)

func errorChek(n int) (string, error) {
    if n == 1 {       
        s := "Not Error"
        return s, nil
    }

	return "", fmt.Errorf("Error : %d != 1", n)
}

func ErrorPractice(){
    s, err := errorChek(1)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(s)

    s, err = errorChek(2)
    if err != nil {
        log.Print(err)
    }
    fmt.Println(s)
    
    defer func() {      
        s, err = errorChek(4)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(s)
    }()
    
    
    s, err = errorChek(3)
    if err != nil {
        log.Panic(err) // defer 함수로 이동
    }
    fmt.Println(s)
    
    // 에러가 발생하여 프로그램이 종료되었으므로 이 아래부터는 실행되지 않음
    fmt.Println(s)

    fmt.Println("Hello, world!")
}