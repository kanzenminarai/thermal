package main

import "fmt"

func DirectInput() {
	var option int8 = 0
	var value float64 = 0

	for {
		ClearScreen()
		fmt.Printf(
			"Insert the unit to convert:\n" +
				"1 - Celsius to Fahrenheit\n" +
				"2 - Celsius to Kelvin\n" +
				"3 - Fahrenheit to Celsius\n" +
				"4 - Fahrenheit to Kelvin\n" +
				"5 - Kelvin to Celsius\n" +
				"6 - Kelvin to Fahrenheit\n" +
				"> ")
		fmt.Scanf("%d", &option)
		fmt.Scan()
		if option < 7 && option > 0 {
			fmt.Print("Insert the value: ")
			fmt.Scanf("%f", &value)
			directConvertion(value, option)
			break
		}
	}
}

func directConvertion(value float64, option int8) {
	switch option {
	case 1:
		fmt.Printf("%f °C = %f °F\n", value, (value*1.8)+32)
	case 2:
		fmt.Printf("%f °C = %f °K\n", value, value+273.15)
	case 3:
		fmt.Printf("%f °F = %f °C\n", value, (value-32)/1.8)
	case 4:
		fmt.Printf("%f °F = %f °K\n", value, ((value-32)/32)+273.15)
	case 5:
		fmt.Printf("%f °K = %f °C\n", value, value-273.15)
	case 6:
		fmt.Printf("%f °K = %f °F\n", value, ((value-273.15)*1.8)+32)
	default:
		fmt.Println("Unknown option.")
	}
}
