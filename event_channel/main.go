package main

import "log"

func main() {

	user1 := NewUser("jkulvich")
	user2 := NewUser("vasya")

	ch1 := NewChannel()
	ch1.Subscribe(user1)
	ch1.Subscribe(user2)

	ch2 := NewChannel()
	ch2.Subscribe(user2)

	pub := NewPublisher()
	log.Println(pub.GetChannels())
	pub.AddChannel("test", ch1)
	pub.AddChannel("test2", ch2)
	log.Println(pub.GetChannels())

	if err := pub.Send("HELLO!", "test"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

	if err := pub.Send("HELLO FROM CH2", "test2"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

	if err := pub.Send("HELLO FROM ALL"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

	if err := pub.DelChannel("test2"); err != nil {
		log.Fatalf("can't del channel: %s", err)
	}
	log.Println(pub.GetChannels())
}
