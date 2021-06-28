# url-checker

### Goroutines
- Goroutines는 호출하려는 func 앞에 go라고 붙여줌으로써 동시에 진행됨.
- 다만 main은 goroutine을 기다려주지 않기때문에 아래처럼 호출하는 모든 함수에 go를 붙이면 아무런 일도 일어나지 않고 끝나버림. 왜냐하면 goroutine은 main이 실행중일 때에만 동작할 수 있기 때문.
```
func main() {
    go count("Hannah")
	go count("Santi")
}

func count(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
```