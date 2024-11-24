package models

type Message struct {
	ID          string  `json:"id"`
	Coordinates string  `json:"coordinates"`
	Fulness     float32 `json:"fulness"`
}
