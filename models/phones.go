package models

type PhoneNum struct {
	Mobile     int `json:"mobile"`
	Nationcode int `json:"nationcode"`
}

type Phones struct {
	Phones []PhoneNum `json:"phone"`
}
