package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount = 1000.0
	var expectedReturnRate = 5.5
	var years = 10.0

	fmt.Print("Investment Amount : ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Number of years :")
	fmt.Scan(&years)

	fmt.Print("Expected return Rate :")
	fmt.Scan(&expectedReturnRate)
	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("the future value is : %.2f \n", futureValue)

	fmt.Printf("the future  real value is : %.2f \n", futureRealValue)

}
