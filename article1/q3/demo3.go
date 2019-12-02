package main

import "fmt"

func main() {
	var fahr, celsius float64
	var lower, upper, step float64

	lower = -60
	upper = 60
	step = 20

	fmt.Printf("%s\t%s\n", "Celsius", "Fahr")
	for celsius = lower; celsius <= upper; celsius = celsius + step {
		fahr = (celsius * (9.0 / 5.0)) + 32
		fmt.Printf("%6.1f\t%3.0f\n", celsius, fahr)
	}
}
