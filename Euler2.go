package main

import "fmt"

func main() {
	var fib,old,top int
	old,fib = 1,old+fib
	for old+fib <4000000{
		old,fib = fib,old+fib
		if fib%2 == 0{
			top +=fib
		}
	}
	fmt.Println(top)
}
