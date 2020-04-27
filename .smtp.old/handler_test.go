package _smtp_old

import (
	"io/ioutil"
	"net/mail"
	"testing"
)

func TestNewMessageHandler(t *testing.T) {
	NewMessageHandler(func(msg *mail.Message) {})
}

func TestMessageHandler_HandleMessage(t *testing.T) {
	body := "FooBar"
	h := &MessageHandler{
		HandlerFunc: func(msg *mail.Message) {

			b, err := ioutil.ReadAll(msg.Body)

			if err != nil {
				t.Error(err)
			}

			if string(b) != body {
				t.Errorf("expected %s, got %s", body, string(b))
			}
		},
	}
	h.HandleMessage(stubMessage(body))
}
