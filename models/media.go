package models

type Media interface {
	Send(Senders,Receivers,Message)error
	Callback()
	ErrorCallback()
}