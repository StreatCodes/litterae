package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/backend"
)

type Backend struct {
	users map[string]*User
}

func (be *Backend) Login(_ *imap.ConnInfo, username, password string) (backend.User, error) {
	fmt.Printf("%s logged in with %s password", username, password)
	user, ok := be.users[username]

	if ok && user.password == password {
		return user, nil
	}

	return nil, errors.New("Bad username or password")
}

func NewBackend() *Backend {
	user := &User{username: "mort@streats.dev", password: "password"}

	body := "From: contact@example.org\r\n" +
		"To: contact@example.org\r\n" +
		"Subject: A little message, just for you\r\n" +
		"Date: Wed, 11 May 2016 14:31:59 +0000\r\n" +
		"Message-ID: <0000000@localhost/>\r\n" +
		"Content-Type: text/plain\r\n" +
		"\r\n" +
		"Hi there :)"

	user.mailboxes = map[string]*Mailbox{
		"INBOX": {
			name: "INBOX",
			user: user,
			Messages: []*Message{
				{
					Uid:   6,
					Date:  time.Now(),
					Flags: []string{"\\Seen"},
					Size:  uint32(len(body)),
					Body:  []byte(body),
				},
			},
		},
		"ayeee": {
			name: "ayee",
			user: user,
			Messages: []*Message{
				{
					Uid:   7,
					Date:  time.Now(),
					Flags: []string{"\\Seen"},
					Size:  uint32(len(body)),
					Body:  []byte(body),
				},
			},
		},
	}

	return &Backend{
		users: map[string]*User{user.username: user},
	}
}