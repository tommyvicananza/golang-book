package main

import (
	"fmt"
)

func add(args ...int) int {
	total := 0

	for _, value := range args {
		total += value
	}
	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func sum(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(add(1, 2, 3, 3))

	addinside := func(x, y int) int {
		return x + y
	}
	fmt.Println(addinside(1, 1))
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4

	sl := []int{1, 2, 3, 4, 5, 6}
	totalsl := sum(sl)
	fmt.Println(totalsl)
	n := fib(10)
	fmt.Println(n)
}
