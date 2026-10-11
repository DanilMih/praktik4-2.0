package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Начало обработки")
	sem := make(chan struct{}, 5)
	tick := time.Tick(200 * time.Millisecond)

	for i := 1; i <= 15; i++ {
		go process(i, sem)
		<-tick
	}

}
func process(id int, sem chan struct{}) {
	sem <- struct{}{}
	defer func() { <-sem }()
	fmt.Printf("Запрос номер %d\n", id)

}
