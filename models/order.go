package models

type Order struct {
	Item        string `json:"item"`
	Quantity    int    `json:"quantity"`
	DeliverType string `json:"deliverType"`
}
