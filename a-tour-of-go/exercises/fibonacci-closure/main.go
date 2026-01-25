package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n1 := 0
	n2 := 1
	var n3 int

	return func() int {
		n3 = n1 + n2
		n1 = n2
		n2 = n3
		
		return n1
	}
}

func main() {
	f := fibonacci()
	for range 10 {
		fmt.Println(f())
	}
}
