// The practical use of interface

package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empID    int
	basicpay int
	pf       int
}

type Contract struct {
	empID    int
	basicpay int
}

type Freelancer struct {
	empID       int
	ratePerHour int
	totalHours  int
}

// Salary of permanent employee
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

// Salary of contract employee
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

// Salary for Freelancers
func (f Freelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

/*
total expense is calculated by iterating through the SalaryCalculator slice and summing
the salaries of the individual employees (both contract and permanent)
*/
func totalExpense(s []SalaryCalculator) {
	expense := 0

	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense per month $%d", expense)
}

func main() {

	perm1 := Permanent{
		empID:    1,
		basicpay: 50000,
		pf:       20,
	}

	perm2 := Permanent{
		empID:    2,
		basicpay: 30000,
		pf:       30,
	}

	cemp1 := Contract{
		empID:    3,
		basicpay: 20000,
	}

	femp1 := Freelancer{
		empID:       4,
		ratePerHour: 10,
		totalHours:  40,
	}

	employees := []SalaryCalculator{perm1, perm2, cemp1, femp1}
	totalExpense(employees)

}
