package main

import (
	"fmt"
)

func main() {
	arrayExample()
	fmt.Println("")
	programArray()
	fmt.Println("")
	programArrayTwo()
	fmt.Println("")
	sliceExample()
	fmt.Println("")
	sliceFunc()
	fmt.Println("")
	sliceCopy()
	fmt.Println("")
	mapExample()
	fmt.Println("")
	mapBigger()
	fmt.Println()
	problemOne()
	fmt.Println()
	problemTwo()
	fmt.Println()
	problemThree()
	fmt.Println()
	problemFour()
}

func arrayExample() {
	fmt.Println("arrayExample")
	var x [5]int
	x[4] = 100
	fmt.Println(x)
}

func programArray() {
	fmt.Println("programArray")
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
}

func programArrayTwo() {
	fmt.Println("programArrayTwo")
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for _, value := range x {
		fmt.Println(value)
		total += value
	}
	fmt.Println(total / float64(len(x)))
}

func sliceExample() {
	//var x []float64
	//x := make([]float64, 5, 10)
	arr := [5]float64{1, 2, 3, 4, 5}
	x := arr[0:5]
	fmt.Println(x)
}

func sliceFunc() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}

func sliceCopy() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

func mapExample() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}
}

func mapBigger() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}

func problemOne() {
	x := [5]int{1, 2, 3, 4, 5}
	fmt.Println(x[3])
}

func problemTwo() {
	x := make([]int, 3, 9)
	fmt.Println(len(x))
}

func problemThree() {
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x[2:5])
}

func problemFour() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	init := 0
	for _, value := range x {
		if init < value {
			init = value
		}
	}
	fmt.Println(init)
}
