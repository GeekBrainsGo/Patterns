package eventchannel

import (
	"reflect"
	"testing"
)

const (
	chanTest  string = "test_ch"
	chanTest2 string = "test_ch2"
)

func TestAddChannel(t *testing.T) {
	ch := NewChannel(chanTest)
	pub := NewPublisher()
	pub.AddChannel(ch)
	if _, ok := pub.channels[chanTest]; !ok {
		t.Errorf("expected exists channel with key=%s, but not", chanTest)
	}
}

func TestDelChannel(t *testing.T) {
	ch := NewChannel(chanTest)
	pub := NewPublisher()
	pub.AddChannel(ch)
	if _, ok := pub.channels[chanTest]; !ok {
		t.Errorf("expected exists channel with key=%s, but not", chanTest)
	}
	pub.DelChannel(chanTest)
	if _, ok := pub.channels[chanTest]; ok {
		t.Errorf("expected not exists channel with key=%s, but not", chanTest)
	}
}

func TestSend(t *testing.T) {
	ch := NewChannel(chanTest)
	pub := NewPublisher()
	pub.AddChannel(ch)
	if _, ok := pub.channels[chanTest]; !ok {
		t.Errorf("expected exists channel with key=%s, but not", chanTest)
	}
	if err := pub.Send("HELLO!", chanTest); err != nil {
		t.Errorf("expected success send message, but error %s", err)
	}
	// wrong chan send
	if err := pub.Send("HELLO!", "testWrong"); err == nil {
		t.Error("expected error send message, but not")
	}

}

func TestListChannels(t *testing.T) {
	ch := NewChannel(chanTest)
	ch2 := NewChannel(chanTest2)
	pub := NewPublisher()
	pub.AddChannel(ch)
	pub.AddChannel(ch2)
	expList := []string{chanTest, chanTest2}
	list := pub.ListChannels()
	if !reflect.DeepEqual(expList, list) {
		t.Errorf("expected list=%v, but got %v", expList, list)
	}

}
