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

	// Slices

	slice_1 := []int{}
	fmt.Println(len(slice_1))
	fmt.Println(cap(slice_1))

	// Create a Slice From an Array

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice_from_arr := arr[3:6]
	fmt.Println((slice_from_arr))

	// Create Slice from make

	slice_make := make([]int, 5, 10)

	fmt.Println("Slice len : ", len(slice_make))
	fmt.Println("Slice cap : ", cap(slice_make))

	// Append elements to slide

	slice_append := []int{1, 2, 3}
	slice_append = append(slice_append, 4, 5)
	fmt.Println(slice_append)

	myslice1 := []int{1, 2, 3}
	myslice2 := []int{4, 5, 6}
	myslice3 := append(myslice1, myslice2...)

	fmt.Printf("myslice3=%v\n", myslice3)
	fmt.Printf("length=%d\n", len(myslice3))
	fmt.Printf("capacity=%d\n", cap(myslice3))
}
