package main

import (
	"fmt"
)

func main() {
	arr := []int{}
	exit := 0
	fmt.Print("\n****** Welcome to stack structure ******\n\n")
	for exit < 1 {
		fmt.Println("-> Select a function by typing the respective number:\n1- Add data on the stack.\n2- Remove data from de stack\n3- Show the stack\n4- Exit")
		var option int
		_, err := fmt.Scan(&option)
		if err != nil {
			panic(err)
		}
		switch option {
		case 1:
			push(&arr)
			show(arr)
		case 2:
			arr = pop(arr)
			show(arr)
		case 3:
			show(arr)
		case 4:
			exit = 1
		default:
			fmt.Print("\nThat's not a valid option\n")
		}
	}
}

func push(arr *[]int) {
	fmt.Print("Type a number you want add to the stack: ")
	var addNumber int
	_, err := fmt.Scanln(&addNumber)
	if err != nil {
		panic(err)
	}
	*arr = append(*arr, addNumber)
	fmt.Printf("\n[SUCCESS] Number %d was added to the stack!\n\n", addNumber)
}

func pop(arr []int) []int {
	removedNumber := arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	fmt.Printf("\n[SUCCESS] Number %d was removed from the stack!\n\n", removedNumber)
	return arr
}

func show(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Printf("[  %d  ]\n", arr[i])
	}
	fmt.Print("\n")
}
