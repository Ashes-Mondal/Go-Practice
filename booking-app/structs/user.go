package structs

import "fmt"

type UserData struct{
	name map[string]string
	ticketsTaken uint
}

func (obj *UserData) Make(name map[string]string,ticketsTaken uint)  {
	obj.name = name
	obj.ticketsTaken = ticketsTaken
}

func (obj *UserData) Print(){
	fmt.Println(*obj)
}