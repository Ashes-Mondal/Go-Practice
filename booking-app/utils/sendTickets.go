package utils

import (
	myType "booking-app/structs"
	"fmt"
	"sync"
	"time"
)
var WG = sync.WaitGroup{}

func SendTicket( obj *myType.UserData) (string){
	time.Sleep(20*time.Second)
	result := fmt.Sprintf("Ticket sent %v",*obj)
	fmt.Println(result)
	WG.Done()
	return result
	
}