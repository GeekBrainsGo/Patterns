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
func (p *Publisher) RemoveChannel(channel string) error {
	if _, ok := p.channels[channel]; ok {
		delete(p.channels, channel)
		return nil
	}
	return fmt.Errorf("channel %s not found", channel)
}

//TODO: Если 0 имён каналов - отправлять всем
func (p *Publisher) Send(msg string, channels ...string) error {
	if len(channels) == 0 {
		for _, channel := range p.channels {
			channel.Send(msg)
		}
	}
	for _, ch := range channels {
		channel, ok := p.channels[ch]
		if !ok {
			return fmt.Errorf("channel %s can't be found", ch)
		}
		channel.Send(msg)
	}
	return nil
}
