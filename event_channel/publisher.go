package main

import (
	"errors"
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

func (p *Publisher) GetChannels() []string {
	keys := make([]string, len(p.channels))

	i := 0
	for k := range p.channels {
		keys[i] = k
		i++
	}
	return keys
}

func (p *Publisher) DelChannel(name string) error {
	if p.channels[name] == nil {
		return errors.New("Channel not found: " + name)
	}
	delete(p.channels, name)
	return nil
}

func (p *Publisher) Send(msg string, channels ...string) error {
	if len(channels) == 0 {
		for _, channel := range p.channels {
			channel.Send(msg)
		}
	} else {
		for _, ch := range channels {
			channel, ok := p.channels[ch]
			if !ok {
				return fmt.Errorf("channel %s can't be found", ch)
			}
			channel.Send(msg)
		}
	}
	return nil
}
