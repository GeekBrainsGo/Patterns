package eventchannel

import (
	"strconv"
	"testing"
)

func TestOnReceive(t *testing.T) {
	type tCase struct {
		title    string
		user     string
		proc     func(string) error
		msg      string
		hasError bool
	}
	cases := []tCase{
		tCase{
			"simple proc OK",
			"user1",
			func(msg string) error {
				return nil
			},
			"blablabla",
			false,
		},
		tCase{
			"proc error",
			"user1",
			func(msg string) error {
				_, err := strconv.Atoi(msg)
				return err
			},
			"blablabla",
			true,
		},
	}
	for _, tcase := range cases {
		t.Run(tcase.title, func(t *testing.T) {
			u := NewUser(tcase.user, tcase.proc)
			if err := u.OnReceive(tcase.msg); (err != nil) != tcase.hasError {
				t.Errorf("expected hasError=%t, but not, error %v", tcase.hasError, err)
			}
		})
	}
}

func TestGetID(t *testing.T) {
	const (
		userID string = "u1"
	)
	u := NewUser(userID, func(msg string) error {
		return nil
	})
	id := u.GetID()
	if id != userID {
		t.Errorf("expected userID=%s, but got %s", userID, id)
	}
}
