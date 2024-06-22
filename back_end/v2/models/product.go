package models

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Introduction string  `json:"introduction"`
	Note         string  `json:"note"`
	ProducerId   int     `json:"producer_id"`
	Msg          string  `json:"msg"`
}
