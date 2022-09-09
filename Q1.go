package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func sumSlice(z []int, c chan int) int {
	sum := 0
	//iterate over slice and sum numbers
	for _, val := range z {
		sum += val
	}

	if c != nil {
		c <- sum //sends the sum to a channel if this is a summation by a goroutine
	}

	return sum
}

func main() {

	x, _ := strconv.Atoi(os.Args[1]) //Converts the command line arg to an int

	p := fmt.Println //taken from gobyexample, makes a func variable to print

	sli := make([]int, x) //Create slice and store random numbers
	for i := 0; i < x; i++ {
		rand.Seed(time.Now().UnixNano())
		sli[i] = rand.Intn(100)
	}

	now := time.Now()               //records current time
	normalSum := sumSlice(sli, nil) //sum the slice in main thread
	then := time.Now()              //records end time
	p(normalSum)                    //prints sum
	p(now.Sub(then))                //prints time summation took

	//Sum with goroutines
	//Code partially from https://gist.github.com/hackintoshrao/5a88f0fab29bd2d2a4062989a9520d91\

	//create channels to receive each routine's respective sum
	c1 := make(chan int)
	c2 := make(chan int)

	now = time.Now()
	go sumSlice(sli[:len(sli)/2], c1) //goroutine sum first half of slice
	go sumSlice(sli[len(sli)/2:], c2) //goroutine sum second half of slice
	s1, s2 := <-c1, <-c2              // get both sums from the channels
	then = time.Now()
	p(s1 + s2) //print parallel sum
	p(now.Sub(then))

}
