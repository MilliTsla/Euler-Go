package main

import (
	"math"
	"fmt"
)

func main() {
fmt.Println(two()-one())
}
func one()int{
	a :=0
	b :=0
	for i:=0;i<=100;i++ {
		a = int(math.Pow(float64(i),2))
		b +=a
		}
		return b
}
func two()int{
	a:=0
	b:=0
	for i:=0;i<=100;i++{
		b +=i
	}
	a = int(math.Pow(float64(b),2))
	return a
}