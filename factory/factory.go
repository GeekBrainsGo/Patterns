package main

import "errors"

// Фабрика которая делегирует создание объекта
func NewClient(token string, typ string) (ClientInterface, error) {
	// Какие-то общие моменты
	switch typ {
	case "user":
		return NewUser(token), nil
	case "bot":
		return NewBot(token), nil
	}
	return nil, errors.New("unknown client type")
}
