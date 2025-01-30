package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	// Variables
	var num1 int = 1 // var <variable_name> datatype
	num2 := 2
	var num3 int = 3

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)

	// Const
	const PI = 3.14
	const rupees int = 1
	fmt.Println(PI)
	fmt.Println(rupees)

	// Array Declaration
	var num = [5]int{1, 2, 3, 4, 5}
	alphabets := [...]string{"a", "b", "c"}
	fmt.Println(num)
	fmt.Println(alphabets)

	// Index Access
	a_alpha := alphabets[0]
	fmt.Println(a_alpha)

	num[3] = 33
	fmt.Println(num)

	// Array Initialization
	init_arr1 := [5]int{}
	init_arr2 := [5]int{1, 2}
	init_arr3 := [5]int{1, 2, 3, 4, 5}

	fmt.Println(init_arr1, init_arr2, init_arr3)

	// Array Initialization - Specific Elements
	init_spec := [5]int{1, 4: 5}
	fmt.Println(init_spec)

	arr_mod := [...]int{1, 3: 4}
	arr_len := len(arr_mod)
	fmt.Println(arr_mod, arr_len)
}
