package main

import (
	"fmt"
	"log"
	"math"
)

func main1() {
	var a, b, c float64
	fmt.Println("a ni kiriting")
	fmt.Scan(&a)
	fmt.Println("b ni kiriting")
	fmt.Scan(&b)
	fmt.Println("c ni kiriting")
	fmt.Scan(&c)
	x1, x2 := Solutions(a, b, c)
	fmt.Println("x1 =", x1)
	fmt.Println("x2 =", x2)
}

func Solutions(a, b, c float64) (float64, float64) {

	var x1, x2 float64

	discr := b*b - 4*a*c

	if discr > 0 {
		x1 = (-b + math.Sqrt(discr)) / (2 * a)
		x2 = (-b - math.Sqrt(discr)) / (2 * a)
	} else if discr == 0 {
		x1 = -b / (2 * a)
		x2 = x1
	} else {
		log.Fatal(" Diskriminant manfiy ! ∅ ")
		return 0, 0
	}

	return x1, x2
}