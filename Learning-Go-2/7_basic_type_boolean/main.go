package main

import "fmt"

func main() {

	isLogin := true
	isAdmin := false
	hasSubscription := true

	canOpenDashboard := isLogin && isAdmin

	canDeletePost := isAdmin || (isLogin && hasSubscription)
	fmt.Println(canOpenDashboard, canDeletePost)

	age:= 12


	isAdult:= age >18

	fmt.Println(isAdult)

}