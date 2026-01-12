package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {

	// u1 := User{name: "Rashedul", age: 45}

	// fmt.Println(u1)

	u:=User {name: "Rashedul",age:45}

	fmt.Println(u.Intro())

}


// value receiver means this method a copy of the user

func (u User) Intro() string{
	return  fmt.Sprintf("Hi I'm %s", u.name)
}



