package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errRequestFailed = errors.New("request failed")

func main() {
	// var results map[string]string 초기화 되지 않은 map에는 값을 넣을 수 있다. 왜냐면 results가 nil이기 때문!
	// var results = make(map[string]string)
	// urls := []string{
	// 	"https://www.google.com",
	// 	"https://www.facebook.com",
	// 	"https://www.airbnb.com",
	// 	"https://www.amazon.com",
	// 	"https://www.reddit.com",
	// 	"https://www.instagram.com",
	// }

	// for _, url := range urls {
	// 	result := "OK"
	// 	err := hitURL(url)
	// 	if err != nil {
	// 		result = "FAILED"
	// 	}
	// 	results[url] = result
	// }
	// for url, result := range results {
	// 	fmt.Println(url, result)
	// }

	// Goroutines 는 호출하려는 func 앞에 go라고 붙여줌으로써 동시에 진행됨.
	// 다만 main은 goroutine을 기다려주지 않기때문에 호출하는 모든 함수에 go를 붙이면 아무런 일도 일어나지 않고 끝나버림.
	// goroutine은 main이 실행중일 때에만 동작할 수 있기 때문.
	// go count("Hannah")
	// go count("Santi")
	// time.Sleep(time.Second * 5)

	// Channel & make()
	channel := make(chan bool)
	people := [2]string{"Hannah", "Santi"}
	for _, person := range people {
		go isOkay(person, channel)
	}
	// result := <-channel // channel로 부터 뭔가를 받으려고 기다리는 동안에는 main이 살아있음.
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

func hitURL(url string) error {
	fmt.Println("checking: ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}

func count(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

func isOkay(person string, channel chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	channel <- true // channel에 true를 보낸다는 의미.
}
