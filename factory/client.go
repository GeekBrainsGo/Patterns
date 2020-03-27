package main

// Общий интерфейс которому должны уовлетворять клиенты чата (боты, пользователи)
type ClientInterface interface {
	Send (msg string) error
	OnReceive (msg string) error // OnReceive, HandleReceive, ReceiveHandler
	Leave () error // Покинуть беседу / Есть у User, но нет у Bot
}

type ClientDefault struct {
	Token string
}

func (c *ClientDefault) Send(msg string) error {
	panic("not implemented")
}

func (c *ClientDefault) OnReceive(msg string) error {
	panic("not implemented")
}

func (c *ClientDefault) Leave() error {
	panic("not implemented")
}