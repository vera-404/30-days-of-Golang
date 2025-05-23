package main 

import (
	"fmt"
	"time"
)
//A function example of switch case
// it is evaluated from top to bottom and stops when a case succeeds
func main (){
	fmt.Println("When's Saturday")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

}