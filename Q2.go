package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {

	x, _ := strconv.Atoi(os.Args[1]) //Converts the command line arg to an int

	p := fmt.Println //taken from gobyexample, makes a func variable to print

	sli := make([]int, x) //Create slice and store random numbers
	for i := 0; i < x; i++ {
		rand.Seed(time.Now().UnixNano())
		sli[i] = rand.Intn(100)
	}

	//Slice Sort
	now := time.Now()                                               //records current time
	sort.Slice(sli, func(i, j int) bool { return sli[i] < sli[j] }) //sorts the slice in ascending order
	then := time.Now()                                              //records end time
	p(now.Sub(then))                                                //prints time summation took
	p(sli)

	//Stable sort
	//Equal elements will not change order
	now = time.Now()                                                      //records current time
	sort.SliceStable(sli, func(i, j int) bool { return sli[i] > sli[j] }) //sorts the slice in descending order (so it changes)
	then = time.Now()                                                     //records end time
	p(now.Sub(then))                                                      //prints time summation took
	
}
