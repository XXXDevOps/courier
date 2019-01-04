package models

type Receiver struct {
	Transceiver
}

type Sender struct {
	Transceiver
}

type Transceiver struct {
	Name string	`json name`
	Address string	`json address`
}

type Senders []Sender
type Receivers []Receiver