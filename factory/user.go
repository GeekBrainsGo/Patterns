package main

type User struct {
	ClientDefault
	Token string
}

func NewUser(token string) *User {
	// Проверка прав пользователя
	return &User{Token: token}
}

func (u *User) Send(msg string) error {
	// Отправляем в сокет и пишем в БД
	return nil
}

func (u *User) OnReceive(msg string) error {
	// Уведомляем пользователя о сообщении
	return nil
}

func (u *User) Leave() error {
	// Покидает чат
	return nil
}