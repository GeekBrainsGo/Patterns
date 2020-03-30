package eventchannel

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

func (p *Publisher) AddChannel(ch *Channel) {
	p.channels[ch.GetID()] = ch
}

func (p *Publisher) DelChannel(id string) {
	delete(p.channels, id)
}

// TODO: [x] Удаление канала и [x] получение списка каналов
func (p *Publisher) ListChannels() []string {
	list := make([]string, 0, len(p.channels))
	for chName := range p.channels {
		list = append(list, chName)
	}
	return list
}

// TODO: [x] Если 0 имён каналов - отправлять всем
func (p *Publisher) Send(msg string, channels ...string) error {
	if len(channels) == 0 {
		p.broadcast(msg)
		return nil
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

func (p *Publisher) broadcast(msg string) {
	for _, ch := range p.channels {
		ch.Send(msg)
	}
}
