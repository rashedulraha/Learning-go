package main

import "fmt"

//!   my type string
// type myType string
	type orderStatus string


	 const(

		received orderStatus = "Received"
			Confirmed = "Confirmed"
			Prepared ="Prepared"
			Delivered ="Delivered"
	 )


 func changeOrderStatus(status orderStatus){
	fmt.Println("changing order status to" ,status)
 }

func main() {
changeOrderStatus(Delivered)
}