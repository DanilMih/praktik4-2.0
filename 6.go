package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	resultChan := make(chan string, 3)
	go searchdb("База1 (Основная)", resultChan)
	go searchdb("База2 (Дополнительная)", resultChan)
	go searchdb("База3 (Прикладная)", resultChan)

	firstResult := <-resultChan
	fmt.Println("Успех", firstResult)
}

func searchdb(dbName string, resultChan chan<- string) {
	timer := time.Duration(rand.Intn(700)+100) * time.Millisecond
	time.Sleep(timer)
	resultChan <- "Результат найден в " + dbName
}
