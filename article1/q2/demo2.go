package main

import "fmt"

func main() {
	var fahr, celsius float64
	var lower, upper, step float64

	lower = 0
	upper = 300
	step = 20

	fmt.Printf("%s\t%s\n", "Fahr", "Celsius")
	for fahr = lower; fahr <= upper; fahr = fahr + step {
		celsius = (5.0 / 9.0) * (fahr - 32.0)
		fmt.Printf("%3.0f\t%6.1f\n", fahr, celsius)
	}
}
