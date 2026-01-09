package main

import (
	"errors"
	"fmt"
)


func withdraw(balance float64, amount float64) (float64, error) {
    if amount > balance {
        return 0, errors.New("Your accepting balance not fount")
    }
    
    newBalance := balance - amount
    return newBalance, nil 
}

func main() {
    currentBalance := 500.0

    withdrawAmount := 600.0
    remainingBalance, err := withdraw(currentBalance, withdrawAmount)


    if err != nil {
        fmt.Println("Transition:", err)
        return 
    }

    
    fmt.Printf("withdraw successfully: %.2f\n", remainingBalance)
}