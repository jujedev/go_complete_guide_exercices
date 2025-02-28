package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	stringFutureValue := fmt.Sprintf("Future value is: %.2f\n", futureValue)
	stringFutureRealValue := fmt.Sprintf("Future value (adjusted for Inflation): %.2f\n", futureRealValue)

	fmt.Printf("%v%v", stringFutureValue, stringFutureRealValue)

}

/*
func outputText(text string) {
	fmt.Print(text)
}
*/

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

// Otra forma de utilizar el return
/*func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return
}*/
