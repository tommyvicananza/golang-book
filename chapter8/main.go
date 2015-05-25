package main

import (
	"fmt"
)

func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func problemOne(xPtr *int) {
	fmt.Println(&xPtr)
}

func square(x *float64) {
	*x = *x * *x
}

func swap(ptrx, ptry *int) {
	tmpx := *ptrx
	*ptrx = *ptry
	*ptry = tmpx
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)

	problemOne(&x)

	y := 1.5
	square(&y)
	fmt.Println(y)

	ptrx := 1
	ptry := 2
	swap(&ptrx, &ptry)
	fmt.Println(ptrx, ptry)
}
