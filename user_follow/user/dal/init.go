package dal

import (
	"TikTokLite_v2/user_follow/user/dal/db"
	"TikTokLite_v2/user_follow/user/dal/redb"
	"TikTokLite_v2/util/trace_id_log/loggers"
)

func Init() {
	loggers.InitLogger()
	db.MysqlInit()
	_ = db.DB.AutoMigrate(&User{})
	redb.RedisInit()
}
