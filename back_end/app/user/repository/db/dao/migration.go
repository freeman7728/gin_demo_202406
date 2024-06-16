package dao

import (
	"database_lesson/app/user/repository/db/model"
)

func migration() {
	err := _db.Set(`gorm:table_options`, "charset=utf8mb4").
		AutoMigrate(&model.User{})
	if err != nil {
		return
	}
}
