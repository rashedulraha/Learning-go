package main

import "fmt"

type payment struct{}

func (p payment) makePayment(amount float32) {

	// razorpayPaymentGW:=razorpay{}
	// razorpayPaymentGW.pay(amount)

	// ! payment stripe 

	 stripePaymentGW:=stripe{}
	 stripePaymentGW.pay(amount)


}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment using razorpay" , amount)
}

// payment stript 

 type stripe struct{}
 func (s stripe) pay(amount float32 ){
	// logic to make payment
	fmt.Println("making payment using stripe" , amount)

 }
func main() {

	newPayment:=payment{}
	newPayment.makePayment(2500)

}