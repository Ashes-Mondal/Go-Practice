package main

import (
	myType "booking-app/structs"
	utils "booking-app/utils"
	"fmt"
	"math/rand"
)

func main(){
	const conferenceName = "Go practice"	
	var ticketCount uint = 0
	ticketCount =  uint(rand.Uint32()) % 50 + 1
	var bookings []myType.UserData = make([]myType.UserData, 0) 
	
	fmt.Println("Welcome to",conferenceName, "Conference")
	fmt.Printf("Total tickets availlable = %d\n",ticketCount)
	
	count:=uint(0) //Eqivalent to "var count uint = 0"
	for ticketCount>0 {
		fmt.Printf("How many tickets(%v): ",ticketCount)
		if utils.ValidateInt(fmt.Scan(&count)) && count<=ticketCount{
			usrData := takeUserInput(ticketCount)
			bookings = append(bookings,usrData)
			book(uint(count),&ticketCount)
			utils.WG.Add(1)
			go utils.SendTicket(&usrData)
		}
	}
	fmt.Println("Sorry no more tickets left!")
	utils.WG.Wait()
	fmt.Println("Following are the bookings:")
	for idx,val := range bookings{
		fmt.Printf("%d := %v\n",idx,val)
	}
}
