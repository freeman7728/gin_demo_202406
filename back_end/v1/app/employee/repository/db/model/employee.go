package model

type Employee struct {
	Id       int
	Name     string
	Account  string
	Password string
	Tel      string
	Salary   float64
	Note     string
	Level    int32
}
type Login struct {
	Id       int    `gorm:"column:id"`
	Account  string `gorm:"column:account"`
	Password string `gorm:"column:password"`
	Level    int32  `gorm:"column:level"`
	Token    string `gorm:"column:token"`
}
