package main

import "fmt"

func main() {
	var operation int8 = 0

	for {
		fmt.Printf("--- Thermal unit converter ---\n\n")

		fmt.Println("Insert the operation below:")
		fmt.Println("1 - Table of thermal units")
		fmt.Println("2 - Direct convertion")
		fmt.Print("> ")
		fmt.Scanf("%d", &operation)

		if operation == 1 || operation == 2 {
			break
		}
		ClearScreen()
	}

	if operation == 1 {
		TableInput()
	}

	if operation == 2 {
		DirectInput()
	}
}
