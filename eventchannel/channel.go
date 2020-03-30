package eventchannel

import "fmt"

type Subscribers map[string]Subscriber

type Channel struct {
	id          string
	subscribers Subscribers
}

func NewChannel(id string) *Channel {
	return &Channel{
		id:          id,
		subscribers: Subscribers{},
	}
}

func (ch *Channel) GetID() string {
	return ch.id
}

func (ch *Channel) Send(msg string) {
	for _, sub := range ch.subscribers {
		sub.OnReceive(msg)
	}
}

func (ch *Channel) Subscribe(sub Subscriber) {
	ch.subscribers[sub.GetID()] = sub
}

func (ch *Channel) UnSubscribe(sub Subscriber) error {
	id := sub.GetID()
	if _, ok := ch.subscribers[id]; ok {
		delete(ch.subscribers, id)
		return nil
	}
	return fmt.Errorf("can't find user %s", id)
}

func (ch *Channel) UnSubscribeAll() error {
	for _, sub := range ch.subscribers {
		return ch.UnSubscribe(sub)
	}
	return nil
}
