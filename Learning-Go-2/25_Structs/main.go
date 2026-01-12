package main

import "fmt"

//  struct groupe related field  into one type
type User struct {

	Id int
	Name string 
	Email string 
	Age int 

}

func main() {

	u1:=User {
		Id:1,
		Name:"Rashedul",
		Email:"email@gmail.com",
		Age: 100,
	}
	fmt.Println(u1 , u1.Age,u1.Email)

	 u1.Age = 500
	 fmt.Println(u1)


	 u2 := User{ Id:2 , Name: "Check two" , Email: "check@gmail.com",Age: 200}
	 fmt.Println(u2)


}