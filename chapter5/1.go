package main

import (
	"fmt"
)

func main() {
	//exerTwo()
	exerThree()
}

func exerTwo() {
	for i := 1; i < 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func exerThree() {
	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz", i)
		} else if i%3 == 0 && i%5 != 0 {
			fmt.Println("Fizz", i)
		} else if i%5 == 0 && i%3 != 0 {
			fmt.Println("Buzz", i)
		}
	}
}
