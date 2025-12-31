package main

import "fmt"

type payment struct{}

func (p payment) makePayment(amount float32) {

	razorpayPaymentGw:=razorpay{}
	razorpayPaymentGw.pay(amount)

}

type razorpay struct{}

func (r razorpay) pay(amount float32) {

	//? logic to make payment
	fmt.Println("making payment using result pay" ,amount)

}


func main() {

	newPayment:=payment{}
	newPayment.makePayment(300)

}