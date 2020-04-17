package main

import (
	"log"
	"time"
)

func main() {
	log.Println("Hello, World!")
	ch := make(chan int)
	go sync(ch, 1, 1)
	go sync(ch, 1, 2)
	v1, v2 := <-ch, <-ch
	log.Println(v1, v2)
}

func wait(sec int) <-chan time.Time {
	timer := time.NewTimer(time.Duration(sec) * time.Second)
	return timer.C
}

func sync(channel chan int, sec int, val int) {
	<-wait(sec)
	channel <- val
}
