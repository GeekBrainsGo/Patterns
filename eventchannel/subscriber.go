package eventchannel

type Subscriber interface {
	OnReceive(msg string) error
	GetID() string
}

type SubscriberDefault struct{}

func (SubscriberDefault) OnReceive(string) error {
	panic("not implemented")
}
func (SubscriberDefault) GetID(string) {
	panic("not implemented")
}
