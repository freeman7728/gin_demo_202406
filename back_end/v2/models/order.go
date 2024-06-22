package models

type Order struct {
	Id         int     `json:"id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
	Time       string  `json:"time"`
	Note       string  `json:"note"`
	EmployeeId int     `json:"employee_id"`
	Msg        string  `json:"msg"`
}
