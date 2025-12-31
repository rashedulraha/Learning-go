package main

import (
	"fmt"
	"time"
)

// order struct


 type order struct{

	id string
	price float32
	status string
	createAt time.Time //! nanosecond precision
 }

func main() {

	 myOrder :=order {
		id:"1",
		price: 45.00,
		status: "pending",
		createAt: time.Now(),
		
	 }

	  myOrder2 := order{
			id:"2",
			price: 50.55,
			status: "delivered",
			createAt: time.Now(),
		}


		myOrder.status = "paid"

	 		fmt.Println("myOrder struct :" ,myOrder)
	 		fmt.Println("MyOrder2 struct :" ,myOrder2)

}