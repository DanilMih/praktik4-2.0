package main
import (
"fmt"
"time"
"sync"

)
func main() {
fmt.Println("Начало")
var wg sync.WaitGroup
wg.Add(1)
go goruitina(&wg)
wg.Wait()
fmt.Println("Конец")
}

func goruitina(wg * sync.WaitGroup) {

for i := 1; i<=5; i++ {
time.Sleep(1 * time.Second)
	fmt.Println(i)
}
wg.Done()

}

