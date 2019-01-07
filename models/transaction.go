package models

type Transaction struct {
	Sender Sender
	Receivers Receivers
	Message Message
	Cc Cc
}

func NewTransaction(sender Sender, receivers Receivers, message Message, cc Cc)*Transaction{
	return &Transaction{sender, receivers, message, cc}
}