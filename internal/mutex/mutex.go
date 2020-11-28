package m

import (
	"fmt"
	"sync"
)

var (
	balance int
)

func deposit(value int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	fmt.Printf("Depositing %d to account with balance %d\n", value, balance)
	balance += value
	mutex.Unlock()
	wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	fmt.Printf("Withdrawing %d from account with balance %d ", value, balance)
	balance -= value
	mutex.Unlock()
	wg.Done()
}

func main() {
	balance = 1000
	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)
	go withdraw(700, &wg, &mutex)
	go deposit(500, &wg, &mutex)
	wg.Wait()
	fmt.Printf("\nNew Balance %d\n ", balance)
}
