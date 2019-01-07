package webserver

const DataField string = "data"
const ErrorField string = "error"

type V1 struct {
	Data interface{} `json:"data"`
	During float64 `json:"during"`
}

type V1Error struct {
	Code int	`json:"code"`
	Error string	`json:"error"`
}