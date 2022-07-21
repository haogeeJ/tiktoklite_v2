package dal

import (
	"TikTokLite_v2/util/trace_id_log/loggers"
	"TikTokLite_v2/video/dal/db"
	"TikTokLite_v2/video/dal/redb"
)

func Init() {
	loggers.InitLogger()
	db.MysqlInit()
	_ = db.DB.AutoMigrate(&Video{})
	redb.RedisInit()
}
