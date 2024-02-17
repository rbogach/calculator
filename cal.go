package main

import (
	"fmt"
	"math"
)

func main() {

	var investmentAmount float64
	var years int64
	const inflationRate = 2.5
	var expectedReturnRate = 5.5
	fmt.Print("Please insert the investment amount ___")
	fmt.Scan(&investmentAmount)
	fmt.Print("Please insert the number of years you want to invest for __")
	fmt.Scan(&years)

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, float64(years))

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
