package main

import (
	"fmt"
	"sync"
	"time"
)

//	func PrintMessage(){
//		time.Sleep(2*time.Second)
//		fmt.Println("Hello from go routine")
//	}

type WebsiteChecker func(string) bool
type result struct {
	url    string
	status bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool) //output to be given
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(url string) {    //goroutines started ,anonymous function 
			status := wc(url)

			resultChannel <- result{
				url:    url,
				status: status,
			}
		}(url)

	}
	for i := 0; i < len(urls); i++ {      //receives goroutines as per exec time
		r := <-resultChannel
		results[r.url] = r.status //r is type of result
	}
	return results

}
func SayHello(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Printf("Hello from goroutine, id=%d , wg=%+v\n", id, wg)
	// fmt.Printf("WaitGroup address: %p\n", wg)
	// fmt.Printf("WaitGroup value: %+v\n", *wg)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go SayHello(i, &wg)
	}
	wg.Wait()

	ch := make(chan string)
	go func() {
		ch <- "Hello world from channel"
	}()
	message := <-ch
	fmt.Println(message)

	websites := []string{
		"google.com",
		"facebook.com",
		"youtube.com",
	}
	checker := func(url string) bool {
		fmt.Println("Checking:", url)
		return true
	}
	results:=CheckWebsites(checker,websites)
	fmt.Println(results)

	fmt.Println("all goroutines are finished")
}
