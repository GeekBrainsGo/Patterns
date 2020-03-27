package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	ch1 := make(chan int, 5)
	ch2 := make(chan string, 5)
	ch3 := make(chan []byte, 5)

	ch1 <- 1
	ch1 <- 3
	ch1 <- 5

	go F1(ch1, ch2)
	go F2(ch2, ch3)
	F3(ch3)

	<- time.After(1 * time.Second)

}

func F1(ch1 chan int, ch2 chan string) {
	for data := range ch1 {
		ch2 <- strconv.Itoa(data)
	}
	close(ch2)
}

func F2(ch2 chan string, ch3 chan []byte) {
	for data := range ch2 {
		ch3 <- []byte(data)
	}
	close(ch3)
}

func F3(ch3 chan []byte) {
	for data := range ch3 {
		fmt.Println(data)
	}
}