package main

type Bot struct {
	ClientDefault
	Token string
}

func NewBot(token string) *Bot {
	// Проверка прав бота
	return &Bot{Token: token}
}

func(b *Bot) Send(msg string) error {
	// Отправляем в сокет и пишем в БД
	return nil
}

func(b *Bot) OnReceive(msg string) error {
	// Уведомляем пользователя о сообщении
	return nil
}