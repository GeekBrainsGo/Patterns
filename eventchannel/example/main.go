package main

import (
	"fmt"
	"log"

	ec "lesson3/eventchannel"
)

func main() {
	proc := func(msg string) error {
		fmt.Printf("MESSAGE GOT: %s\n", msg)
		return nil
	}
	user1 := ec.NewUser("jkulvich", proc)
	user2 := ec.NewUser("vasya", proc)

	ch1 := ec.NewChannel("test")
	ch1.Subscribe(user1)
	ch1.Subscribe(user2)

	ch2 := ec.NewChannel("test2")
	ch2.Subscribe(user2)

	pub := ec.NewPublisher()
	pub.AddChannel(ch1)
	pub.AddChannel(ch2)

	if err := pub.Send("HELLO!", "test"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

	if err := pub.Send("HELLO FROM CH2", "test2"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

	if err := pub.Send("HELLO BROADCAST"); err != nil {
		log.Fatalf("can't send: %s", err)
	}
}
