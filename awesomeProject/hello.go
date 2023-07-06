package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Println("1 + 1 =", 1.2+1.0)
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	c := 2
	//These two lines are same
	fmt.Printf("%T %[1]v\n", c)
	fmt.Printf("%T %v\n", c, c)
	//*********

	var x string
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)

	var y string
	y = "first"
	fmt.Println(y)
	y = y + "second"
	fmt.Println(y)
	/*
		These three are same
		var str="Hello"
		str:="Hello"
		var str string ="Hello"
	*/
	//Sample program for taking input
	/*
		var input float32
		fmt.Scanf("%f", &input)
		output := 2 * input
		fmt.Println(output)
	*/
	//Control structures
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
	inp := 10
	switch inp {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println("1")
	default:
		fmt.Println("Unknown number")
	}

	//Arrays declaration and printing
	var arr [5]int
	arr[4] = 100
	fmt.Println(arr)
	for index, i := range arr {
		fmt.Println(index, "=", i)
	}
	for _, i := range arr {
		fmt.Println(i)
	}
	slice1 := []int{1, 2, 3}
	slice1 = append(slice1, 4, 5)
	fmt.Println(slice1)
	slice2 := make([]int, 0, 5)
	fmt.Println(slice2)
}
