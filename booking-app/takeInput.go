package main

import (
	myType "booking-app/structs"
	"fmt"
)

func takeUserInput( ticketCount uint)(myType.UserData){
	var temp map[string]string = make(map[string]string,1)
	var name string

	fmt.Printf("Enter fname: ")
	fmt.Scanln(&name)
	temp["firstName"] = name

	fmt.Printf("Enter lname: ")
	fmt.Scanln(&name)
	temp["lastName"] = name


	newRec  := new(myType.UserData)
	newRec.Make(temp,ticketCount)
	return *newRec
}