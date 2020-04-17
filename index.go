package main

import (
	"context"
	"log"
	"math/rand"
	"time"
)

func main() {
	go bg()
	log.Println("Hello, World!")
	ch := make(chan int)
	go sync(ch, 1, 1)
	go sync(ch, 1, 2)
	v1, v2 := <-ch, <-ch
	log.Println(v1, v2)

	ctx := context.Background()
	testContext(ctx)
	log.Println("done")
}

func wait(sec int) <-chan time.Time {
	timer := time.NewTimer(time.Duration(sec) * time.Second)
	return timer.C
}

func sync(channel chan int, sec int, val int) {
	<-wait(sec)
	channel <- val
}

func bg() {
	<-wait(3)
	log.Println("after 3 sec!")
}

func testContext(ctx context.Context) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	rand.Seed(time.Now().UnixNano())
	for {
		n := rand.Intn(30)
		log.Println(n)

		select {
		case <-c.Done():
			log.Println("cancel!")
			return
		default:
			if n < 5 {
				log.Println("return!")
				return
			}
		}
		<-wait(1)
	}
}
