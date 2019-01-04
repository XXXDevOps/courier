package models

type Transaction struct {
	Senders Senders
	Receivers Receivers
	Message Message
	Media Media
}

func (t *Transaction)Run()error{
	return t.Media.Send(t.Senders, t.Receivers, t.Message)
}