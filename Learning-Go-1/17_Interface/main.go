package main

import "fmt"

type payment struct{
	gateway stripe
}


//  inter face
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


//   fake payment gateway 

 type fakePayment struct{}

  func (f fakePayment) pay(amount float32){
		fmt.Println("making payment using testing purpose",amount)
	}


func main() {
	 stripePaymentGW:=stripe{}


	newPayment:=payment{
		gateway: stripePaymentGW,
	}
	newPayment.makePayment(4500)

}