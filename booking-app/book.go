package main

import (
	"fmt"
)


func book(ticketCount uint,available *uint)(string){
	if(ticketCount<=*available){
		msg := fmt.Sprintf("Successfully booked %d tickets",ticketCount)
		*available-=ticketCount
		return msg
	}else{
		return "Failed to book"
	}
}