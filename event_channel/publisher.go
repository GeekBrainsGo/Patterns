package main

import (
	"fmt"
)

type Channels map[string]*Channel

type Publisher struct {
	channels Channels
}

func NewPublisher() *Publisher {
	return &Publisher{
		channels: Channels{},
	}
}

func (p *Publisher) AddChannel(name string, channel *Channel) {
	p.channels[name] = channel

}

//TODO: Удаление канала и получение списка каналов
func (p *Publisher) DeleteChannel(name string) {
	delete(p.channels, name)

}
func (p *Publisher) ShowChannels() {
	channels := []string{}
	for ch := range p.channels {
		channels = append(channels, ch)
	}
	fmt.Printf("Каналы: %s\n", channels)
}

//TODO: Если 0 имён каналов - отправлять всем
func (p *Publisher) Send(msg string, channels ...string) error {
	for _, ch := range channels {
		if ch == "" {
			for _, channel := range p.channels {
				channel.Send(msg)
			}
		} else {
			channel, ok := p.channels[ch]
			if !ok {
				return fmt.Errorf("channel %s can't be found", ch)
			}
			channel.Send(msg)
		}
	}
	return nil
}

func (p *Publisher) SendAll(msg string) error {
	for _, channel := range p.channels {
		channel.Send(msg)
	}
	return nil
}
