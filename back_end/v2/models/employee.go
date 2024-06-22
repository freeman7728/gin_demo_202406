package models

type Employee struct {
	ID        int     `json:"id"`
	Account   string  `json:"account"`
	Password  string  `json:"password"`
	Name      string  `json:"name"`
	Tel       string  `json:"tel"`
	Salary    float64 `json:"salary"`
	Note      string  `json:"note"`
	Level     int     `json:"level"`
	Msg       string  `json:"msg"`
	Email     string  `json:"email"`
	EmailAuth bool    `json:"email_auth"`
}
