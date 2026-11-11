package main

import (
	"fmt"
	"net/http"
	"time"
)

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
var configureTimeOut=10*time.Second
func Racer(a,b string) (string,error){
	return  ConfigurableRacer(a,b,configureTimeOut)
}
func ConfigurableRacer(a, b string,timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		fmt.Println("got url A",a)
		return a, nil
	case <-ping(b):
		fmt.Println("got url B",b)
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out")
	}

}
func main() {
	// ch1 := make(chan string)
	// ch2 := make(chan string)

	// go func(){
	// 	time.Sleep(2*time.Second)
	// 	ch1<-"task 1 completed"
	// }()

	// go func(){
	// 	time.Sleep(1*time.Second)
	// 	ch2<-"task 2 completed"
	// }()

	// select{
	// case msg:=<-ch1:
	// 	fmt.Println(msg)

	// case msg:=<-ch2:
	// 	fmt.Println(msg)
	// }
	urlA:="http://www.facebook.com"
	urlB:="http://www.google.com" 
	Racer(urlA,urlB)

}
