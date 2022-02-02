package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup //this makes the main func wait so that other go routine can work.
	//if the main function exits before the other go routine then the other go routine won't be executed
	wait.Add(1) //1 here is saying that there is one go routine and if we add more value than the present go routine
	//then it'll go into deadlock as wait.Wait() waits for the response of wait.Done(). wait.Wait() makes the main function wait
	//until the counter reaches 0. This counter is the value in wait.Add(1)here it is 1.

	//go.hello(wait)//this is pass by reference and it wont work

	defer wait.Wait()
	go hello(&wait) //here we are receiving value by reference and pass by value does not work as the value doesnt get
	//updated properly
}
func hello(wait *sync.WaitGroup) {
	fmt.Println(1)
	wait.Done()
	//if we have multiple go routines then the go routine does not get executed serially , it gets executed randomly
}

//there is also a chance of
