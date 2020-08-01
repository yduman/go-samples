package main

import (
	"fmt"
	myMath "go-samples/utils"
	"reflect"
)

// ------------------------------------------ Types ------------------------------------------ //
var age int = 26
var height float64 = 1.9
var name string = "Yadullah"
var isDrinking bool = true
var inferred = "I am inferred as a string"

// ------------------------------------------ Funcs ------------------------------------------ //
func printAge(age int) int {
	fmt.Println("I am the function printAge and returning: ", age)
	return age
}

func multipleReturn() (arg1 string, arg2 string) {
	arg1 = "foo"
	arg2 = "bar"
	return
}

func multipleArgs(ages ...int) {
	for _, value := range ages {
		fmt.Println(value)
	}
}

func main() {
	anotherVariable := "I need to be declared inside a function"

	fmt.Printf("Hello I am %s and I am %d years old. I am %.2fm tall and I drink often, which is %t!\n", name, age, height, isDrinking)
	fmt.Println(reflect.TypeOf(height))
	fmt.Println(reflect.TypeOf(inferred))
	fmt.Println(anotherVariable)

	// ------------------------------------------ Control Structs ------------------------------------------ //
	if age > 43 {
		fmt.Println(age)
	} else if age == 42 {
		fmt.Println(age)
	} else {
		fmt.Println(age)
	}

	// if this function returns a value, it will be an error of type Error
	// err := someFunction()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	//
	// if err := someFunction(); err != nil {
	// 	fmt.Println(err.Error())
	// }

	switch name {
	case "abc":
		fmt.Println("abc")
	case "def", "hij":
		fmt.Println("def or hij")
	default:
		fmt.Println("fuck")
	}

	// ------------------------------------------ Loops ------------------------------------------ //

	// classic foo
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// like while
	j := 0
	for j <= 10 {
		fmt.Printf("%d ", j)
		j++
	}
	fmt.Println("")

	// range loop
	for index, letter := range name {
		fmt.Println("Index: ", index, "Letter: ", string(letter))
	}

	printAge(age)
	foo, bar := multipleReturn()
	fmt.Println(foo, bar)

	// ------------------------------------------ Arrays ------------------------------------------ //
	var arr [3]float64
	arrr := [3]float64{6.9, 4.2, 420}
	arrrr := [...]float64{6.9, 4.2, 420}
	fmt.Println(arr)
	fmt.Println(arrr)
	fmt.Println(arrrr)

	for _, val := range arrrr {
		fmt.Println("==> ", val)
	}

	var slice []int = make([]int, 5, 10)
	fmt.Println("length of slice: ", len(slice))
	fmt.Println("capacity of slice: ", cap(slice))

	yeah := [5]string{"best", "shit", "ever", "I", "swear"}
	splicedYeah := yeah[1:3]
	fmt.Println(splicedYeah)
	fmt.Println(len(splicedYeah))
	fmt.Println(cap(splicedYeah))
	moreYeahs := append(splicedYeah, "yeah", "yeah")
	fmt.Println(moreYeahs)

	// ------------------------------------------ Maps ------------------------------------------ //
	fooMap := map[int]string{
		1: "foo",
		2: "bar",
	}

	fooMap[1] = "baz"
	fmt.Println(fooMap)

	youThere, ok := fooMap[69]
	fmt.Println(youThere, ok)

	if youThere, ok := fooMap[69]; ok {
		fmt.Println("found: ", youThere)
	} else {
		fmt.Println("nothing")
	}

	delete(fooMap, 2)

	// ------------------------------------------ Toolkit ------------------------------------------ //

	// call functions from different packages
	myTotal := myMath.Add(1, 2, 3, 4, 5)
	fmt.Println("myTotal=", myTotal)
}
