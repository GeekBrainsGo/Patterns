package main

import "fmt"

func main() {

	bot, err := NewClient("jfsjkd", "bot")
	user, err := NewClient("kdsakfjb", "user")
	fmt.Printf("bot=%+v\nuser=%+v\nerr=%v", bot, user, err)

}
