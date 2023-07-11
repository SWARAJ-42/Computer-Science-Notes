package main

import (
	"fmt"
)

func getExpenseReport(e expense) (string, float64) {
	// e.(sms) returns a bool too as second var to show whether the argument type is sms or not
	// phone, ok := e.(sms)
	// email, ok_1 := e.(email)

	// if ok {
	// 	return phone.toPhoneNumber, e.cost()
	// }
	// if ok_1 {
	// 	return email.toAddress, e.cost()
	// }
	// return "", 0

	// switch case statement for type assertion
	switch T := e.(type) {
		case email:
			return T.toAddress, T.cost() 
		case sms:
			return T.toPhoneNumber, T.cost()
		default:
			return "", 0
	}
}

// don't touch below this line

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func estimateYearlyCost(e expense, averageMessagesPerYear int) float64 {
	return e.cost() * float64(averageMessagesPerYear)
}

func test_assert(e expense) {
	address, cost := getExpenseReport(e)
	switch e.(type) {
	case email:
		fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	case sms:
		fmt.Printf("Report: The sms going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	default:
		fmt.Println("Report: Invalid expense")
		fmt.Println("====================================")
	}
}

func main() {
	test_assert(email{
		isSubscribed: true,
		body:         "hello there",
		toAddress:    "john@does.com",
	})
	test_assert(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		toPhoneNumber: "+155555504444",
	})
	test_assert(invalid{})
}