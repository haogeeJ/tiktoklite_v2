package db

import (
	"TikTokLite_v2/common/gorm_tracing"
	"TikTokLite_v2/util/trace_id_log"
	"TikTokLite_v2/video/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func MysqlInit() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", setting.Conf.MysqlConfig.User, setting.Conf.MysqlConfig.Password, setting.Conf.MysqlConfig.Host, setting.Conf.MysqlConfig.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: trace_id_log.NewGormLogger(logger.Error), //输出sql语句日志
	})
	if err != nil {
		log.Println("err in MysqlInit:", err)
		return
	}
	DB = db
	DB.Use(&gorm_tracing.OpentracingPlugin{})
}
