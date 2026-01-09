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
}