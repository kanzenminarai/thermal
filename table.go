package main

import "fmt"

func TableInput() {
	var value float64 = 0

	ClearScreen()
	fmt.Print("Insert the value: ")
	fmt.Scanf("%f", &value)
	tableConvertion(value)
}

func tableConvertion(value float64) {
	fmt.Printf("%f °C = %f °F\n", value, (value*1.8)+32)
	fmt.Printf("%f °C = %f °K\n", value, value+273.15)
	fmt.Printf("%f °F = %f °C\n", value, (value-32)/1.8)
	fmt.Printf("%f °F = %f °K\n", value, ((value-32)/32)+273.15)
	fmt.Printf("%f °K = %f °C\n", value, value-273.15)
	fmt.Printf("%f °K = %f °F\n", value, ((value-273.15)*1.8)+32)
}
