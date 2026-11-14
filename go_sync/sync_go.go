package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

type Counter struct {
	value int
	mu    sync.Mutex
}
type Counter2 struct {
	value atomic.Int64
}

func (c *Counter2) Inc() {
	c.value.Add(1)
}
func (c *Counter2) Val() int64 {
	return c.value.Load()

}
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func (b *BankAccount) Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done() //defer is last in first out
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance += amount
}

func (b *BankAccount) ShowBalance() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.balance
}

func main() {

	account := BankAccount{balance: 100}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go account.Deposit(100, &wg)
	}
	wg.Wait()
	fmt.Println("balance is ", account.ShowBalance())

	counter := Counter{value: 0}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()
	fmt.Println("counter is(with Mutex) :", counter.Value())

	counter2 := Counter2{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter2.Inc()
		}()

	}
	wg.Wait()

	fmt.Println("counter is (without mutex): ", counter2.Val())
}
