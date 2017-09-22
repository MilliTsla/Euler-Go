package main

import "fmt"

func main() {
var top int
	for i :=0;i<1000;i++ {
		if i%3 ==0{
			top +=i
		}else if i%5==0{
			top +=i
		}
	}
	fmt.Println(top)
	}

