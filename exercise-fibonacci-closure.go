package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
  prev := 0
  next := 0
  return func() int {
	if temp:= prev+next; temp==0 {
	  	next = 1
		return 1
	} else {
		prev = next
	  	next = temp
	  	return temp
	}
  }
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
