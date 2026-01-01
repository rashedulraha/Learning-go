package main

import "fmt"

type payment struct{}

func (p payment) makePayment(amount float32) {

	// razorpayPaymentGw:=razorpay{}
	stripePaymentGW:=stripe{}
	stripePaymentGW.pay(amount)
	// razorpayPaymentGw.pay(amount)

}

type razorpay struct{}

func (r razorpay) pay(amount float32) {

	//? logic to make payment
	fmt.Println("making payment using result pay" ,amount)

}


//!  payment stripe 

 type stripe struct {}
 func ( s stripe) pay(amount float32){
	fmt.Println("Making payment using stripe pay", amount)
 }

func main() {

	newPayment:=payment{
		
	}
	newPayment.makePayment(300)

}