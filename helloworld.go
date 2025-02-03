package main

import (
	"fmt"
)

// Structure

type Person struct {
	name string
	age  int
}

// Multiple Return Values
func divAndMul(x int, y int) (mul int, div int) {
	mul = x * y
	div = x / y
	return
}

// Named Return Values
func sub(x int, y int) (result int) {
	result = x - y
	return
}

func add(x int, y int) int {
	return x + y
}

func helloMessage(name string) {
	fmt.Println("Hello from , ", name, " !!!")
}
func main() {

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

	// if condition

	numX := 20
	numY := 30

	if numX > numY {
		fmt.Println("NumX is greater")
	} else if numY > numX {
		fmt.Println("NumY is greater")
	} else {
		fmt.Println("Both are same")
	}

	mark := 97

	switch mark {
	case 99, 100:
		fmt.Println("Grade A")
	case 98:
		fmt.Println("Grade B")
	default:
		fmt.Println("Grade C")
	}

	// for loop

	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}

	// range in loop

	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}

	for idx, _ := range fruits {
		fmt.Printf("%v\n", idx)
	}

	// Functions
	helloMessage("Sathish")

	sum := add(10, 20)
	fmt.Println(sum)

	fmt.Println(sub(10, 20))

	_, div := divAndMul(20, 10)
	fmt.Println(div)

	// Struct

	var student Person
	student.name = "SK"
	student.age = 21

	fmt.Println("Name : ", student.name, " Age : ", student.age)

	// Create Map using var and :=

	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	fmt.Println(a["brand"])

	// Create Maps Using make() Function

	var map_a = make(map[string]string)
	map_a["brand"] = "Ford"
	map_a["model"] = "Mustang"
	map_a["year"] = "1964"

	map_b := make(map[string]int)
	map_b["Oslo"] = 1
	map_b["Bergen"] = 2
	map_b["Trondheim"] = 3
	map_b["Stavanger"] = 4

	fmt.Printf("a\t%v\n", map_a)
	fmt.Printf("b\t%v\n", map_b)

	// Delete a key in map

	delete(map_b, "Oslo")
	fmt.Printf("b\t%v\n", map_b)

	// Check For Specific Elements in a Map

	// val, ok :=map_name[key]

	var check_a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": "Satur"}

	val1, ok1 := check_a["brand"] // Checking for existing key and its value
	val2, ok2 := check_a["color"] // Checking for non-existing key and its value
	val3, ok3 := check_a["day"]   // Checking for existing key and its value
	_, ok4 := check_a["model"]    // Only checking for existing key and not its value

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)

	// Reference Map

	ref_map := check_a
	fmt.Println(ref_map)

	// Iterate Over Maps

	for k, v := range ref_map {
		fmt.Println(k, v)
	}

	// Variadic function
	fmt.Println("Sum of 1, 2, 3:", sumVariadic(1, 2, 3))

	// Anonymous function in Go Language

	value := func(data string) {
		fmt.Println(data)
	}
	value("Anonymous functions")

}

func sumVariadic(nums ...int) int {
	total := 0
	for n := range nums {
		total += n
	}
	return total
}
