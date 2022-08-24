package main

import "fmt"

func main() {
	var basicSalary, HRA, DA, total float64
	var employeeName string

	fmt.Println("Welcome to Goland Employee Calculation")
	fmt.Println("Enter Employee Salary")
	fmt.Scanln(&basicSalary)
	fmt.Println("Enter Employee Name")
	fmt.Scanln(&employeeName)

	if basicSalary <= 10000 {
		HRA = (basicSalary * 5) / 100
		DA = (basicSalary * 10) / 100
	} else if basicSalary <= 20000 {
		HRA = (basicSalary * 15) / 100
		DA = (basicSalary * 20) / 100
	} else {
		HRA = (basicSalary * 25) / 100
		DA = (basicSalary * 30) / 100
	}
	total = basicSalary + HRA + DA
	fmt.Println("The total basic gross salary of", employeeName, "is ", total)
}
