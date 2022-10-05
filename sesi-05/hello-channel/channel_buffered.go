package main

import (
	"fmt"
	"time"
)

func main() {
	// c := make(chan int,3)

	// c1 := make(chan int)

	// go func(c chan int) {
	// 	fmt.Println("func goroutine starts sending data into the channel")
	// 	c <- 10
	// 	fmt.Println("func goroutine after sending data into the channel")
	// }(c1)

	// fmt.Println("main goroutine sleeps for 2 second")
	// time.Sleep(time.Second * 2)

	// fmt.Println("Main goroutine starts receiving data")
	// d := <-c1
	// fmt.Println("main goroutine received data:", d)

	// close(c1)
	// time.Sleep(time.Second)

	// c1:= make(chan int,3)

	// go func(c chan int){
	// 	for i :=1;i<=5;i++ {
	// 		fmt.Printf("Func goroutine #%d starts sending data into the channel \n",i)
	// 		c<-i
	// 		fmt.Printf("func goroutine #%d after sending data into the channel \n",i)
	// 	}
	// 	close(c)
	// }(c1)

	// fmt.Println("main goroutine skleeps 2 second")
	// time.Sleep(time.Second*2)

	// for v:= range c1{
	// 	fmt.Println("main goroutine received value from channel:",v)
	// }

	// ------------------------------------------------------
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)

		c1 <- "Hello!"
	}()

	go func() {
		time.Sleep(1 * time.Second)

		c2 <- "Salut!"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}
}
