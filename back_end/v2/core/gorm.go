package core

import (
	"database_lesson/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GormInit() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Log.Warningln("未配置mysql,取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()
	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.Log.Fatalf("[%s] mysql连接失败", dsn)
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour * 4)
	if global.Config.Mysql.Log_level == "dev" {
		global.Log.Infof("mysql连接成功 [%s]\n", dsn)
	}
	return db
}
