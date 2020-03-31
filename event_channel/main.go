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
	ch3 := NewChannel()
	pub := NewPublisher()
	pub.AddChannel("test", ch1)
	pub.AddChannel("test2", ch2)
	pub.AddChannel("testDelete", ch3)
	pub.ShowChannels()
	pub.DeleteChannel("testDelete")
	pub.ShowChannels()

	if err := pub.Send("HELLO!", "test"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

	if err := pub.Send("HELLO FROM CH2", "test2"); err != nil {
		log.Fatalf("can't send: %s", err)
	}
	if err := pub.Send("Привет во все каналы!", ""); err != nil {
		log.Fatalf("can't send: %s", err)
	}
	if err := pub.SendAll("Привет во все каналы 2!"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

}
