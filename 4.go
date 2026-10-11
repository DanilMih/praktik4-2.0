package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	sem := make(chan struct{}, 3)
	urls := []string{
		"https://www.google.com/",
		"https://workspace.ktk-45.ru/",
		"https://www.microsoft.com/",
		"https://github.com/",
		"https://vk.com"}
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)

		go func() {
			sem <- struct{}{}
			defer func() { <-sem; wg.Done() }()
			r, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("Сайт", url, "Текущий статус", r.Status)
		}()

	}
	wg.Wait()
}
