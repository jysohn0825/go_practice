package practice

import (
    "fmt"
    "math/rand"
    "time"
    "sync"
)

func hello(n int, w *sync.WaitGroup) {
    defer w.Done() //끝났음을 전달
    r := rand.Intn(3)
    time.Sleep(time.Duration(r) * time.Second)
    fmt.Println(n)  
}

func GoroutinePractice(){
	var wait sync.WaitGroup
    wait.Add(102)
 
	str := "kona!"
	
    go func() {
        defer wait.Done()
        fmt.Println("Hello")
    }()
	
	go func() {
        defer wait.Done()
        fmt.Println(str)
    }()
 
	for i := 0; i<100; i++ {
		go func(n int) {
			defer wait.Done()
			
			fmt.Println(n)
		}(i)
	}
 
    wait.Wait()
}