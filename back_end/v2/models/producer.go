package models

type Producer struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	ShortName   string `json:"short_name"`
	Address     string `json:"address"`
	Tel         string `json:"tel"`
	Email       string `json:"email"`
	ContactName string `json:"contact_name"`
	ContactTel  string `json:"contact_tel"`
	Note        string `json:"note"`
	Msg         string `json:"msg"`
}
