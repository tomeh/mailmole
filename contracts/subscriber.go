package contracts

import "net/mail"

type Sub interface {
	Pub(message *mail.Message)
}
