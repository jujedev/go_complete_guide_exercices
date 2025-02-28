package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate float64 = 2.5

	var investmentAmount float64
	var expectedReturnRate float64
	var years float64
	var nombre string

	fmt.Scan(&nombre)
	fmt.Println("Hello", nombre)

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future value is:", futureValue)
	fmt.Println("Future real value is:", futureRealValue)
}
