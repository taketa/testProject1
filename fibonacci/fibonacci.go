package main

import (
	"fmt"
	"encoding/json"
	"time"
)

var (
	x, y                     = 0, 1
	mistakes, correct, index,input int
	c                        = make(chan int)
	fin                      = make(chan bool)
	f                        = fibonacci()
)

type CorrectAnswer struct {
	Answer int
	Index  int
}

func main() {
	go func(){
		for{
			fmt.Scanln(&input)
			c<-input
		}
	}()

	go func() {
		for {
			select {
			case m := <-c:
				if m != x {
					answer()
					mistakes++
					if isLoose() {
						fin <- true
					}
					f()
				} else {
					correct++
					if correct == 10 {
						fin <- true
					}
					mistakes = 0
					f()
				}
			case <-time.After(10 * time.Second):
				answer()
				mistakes++
				if isLoose() {
					fin <- true
				}
				f()
			}
		}
	}()
	<-fin

}

func answer() {
	answ, err := json.Marshal(&CorrectAnswer{x, index})
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(answ))

}
func isLoose() bool {
	if mistakes == 3 {
		return true
	} else {
		return false
	}
}

func fibonacci() func() int {

	return func() int {
		index++
		x, y = y, x+y
		return x
	}
}

