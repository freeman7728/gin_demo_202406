package models

type Detail struct {
	Id         int    `json:"id"`
	ProductId  int    `json:"product_id"`
	Quantity   int    `json:"quantity"`
	TotalPrice int    `json:"total_price"`
	Note       string `json:"note"`
	ListId     int    `json:"list_id"`
	Msg        string `json:"msg"`
}