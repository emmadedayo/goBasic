package main

import (
	"fmt"
	"math"
)

func main() {

	//Write a Go Program to calculate Compound Interest. This golang program allows the user to enter the principal amount, totals years, and interest rates and then find the Compound Interest.

	// Using the formula, A = P(1+r/n)^nt
	//
	//A	=	final amount
	//P	=	initial principal balance
	//r	=	interest rate
	//n	=	number of times interest applied per time period
	//t	=	number of time periods elapsed

	var principalBalance, interestRate, interestTime, timePeriod, semiTotal, compoundTotal float64

	fmt.Println("Welcome to Compound interest Calculation")

	fmt.Println("Please Enter your Principal Amount")

	fmt.Scanln(&principalBalance)

	fmt.Println("Please Enter Interest Rate")

	fmt.Scanln(&interestRate)

	fmt.Println("Please Enter Interest Time")

	fmt.Scanln(&interestTime)

	fmt.Println("Please Enter Interest Period")

	fmt.Scanln(&timePeriod)

	semiTotal = principalBalance * (math.Pow(1+interestRate/100, timePeriod))

	fmt.Println("-----------------", semiTotal)

	compoundTotal = semiTotal - principalBalance

	fmt.Println("The compound total is ", compoundTotal)
	fmt.Println("\n")
	fmt.Println("The future total is ", semiTotal)

}
