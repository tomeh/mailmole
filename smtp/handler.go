package smtp

import "net/mail"

type Handler interface {
	HandleMessage(*mail.Message)
}

func NewMessageHandler(h func(*mail.Message)) Handler {
	return &MessageHandler{
		HandlerFunc: h,
	}
}

type MessageHandler struct {
	HandlerFunc HandlerFunc
}

func (h *MessageHandler) HandleMessage(msg *mail.Message) {
	h.HandlerFunc(msg)
}

type HandlerFunc func(*mail.Message)
