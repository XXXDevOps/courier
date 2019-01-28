package models

import (
	"strconv"
	"strings"
)

type Transceiver struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Sender Transceiver
type Receivers []Transceiver
type Cc []Transceiver

func (rs Receivers) ToString() string {
	addresses := []string{}
	splitter := ";"
	for _, v := range rs {
		addresses = append(addresses, v.Address)
	}
	return strings.Join(addresses, splitter)
}

func (rs Receivers) Addresses() []string {
	addresses := []string{}
	for _, v := range rs {
		addresses = append(addresses, v.Address)
	}
	return addresses
}
func (rs Receivers) Phone() []int {
	phones := []int{}
	for _, v := range rs {
		p, err := strconv.Atoi(v.Address)
		if err == nil {
			phones = append(phones, p)
		}
	}
	return phones
}
func (cc Cc) Addresses() []string {
	addresses := []string{}
	for _, v := range cc {
		addresses = append(addresses, v.Address)
	}
	return addresses
}
