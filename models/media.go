package models

type Media interface {
	Send(*Transaction)error
	GetName()string
	Init()error
	Close()error
	Redial()error
}