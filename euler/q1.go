package main

import (
	"fmt"
)

func main() {
	fmt.Println("Project Eular Question 1")
	c := make(chan int)
	b := make(chan bool, 1)

	go func() {
		for i := 1; i < 1000; i++ {
			if i%3 == 0 || i%5 == 0 {
				c <- int(i)
			}
		}
		c <- 0
	}()

	go func() {
		sum := 0
		for {
			number := <-c

			if number != 0 {
				sum = sum + number
			} else {
				fmt.Println("Total =", sum)
				b <- true
				return
			}
		}
	}()

	<-b
}
