package eventchannel

type User struct {
	SubscriberDefault
	Username string
	proc     func(string) error
}

func NewUser(username string, proc func(string) error) *User {
	return &User{
		Username: username,
		proc:     proc,
	}
}

func (u *User) OnReceive(msg string) error {
	//fmt.Printf("MESSAGE GOT: %s: %s\n", u.Username, msg)
	return u.proc(msg)
}
func (u *User) GetID() string {
	return u.Username
}
