package dao

import (
	"database_lesson/app/employee/repository/db/model"
)

func migration() {
	err := _db.Set(`gorm:table_options`, "charset=utf8mb4").
		AutoMigrate(&model.Employee{})
	if err != nil {
		return
	}
}
