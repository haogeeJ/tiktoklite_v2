package dal

import (
	"TikTokLite_v2/favorite_comment/dal/db"
	"TikTokLite_v2/favorite_comment/dal/redb"
	"TikTokLite_v2/util/trace_id_log/loggers"
)

func Init() {
	loggers.InitLogger()
	db.MysqlInit()
	_ = db.DB.AutoMigrate(&Favorite{})
	_ = db.DB.AutoMigrate(&Comment{})
	_ = db.DB.AutoMigrate(&User{})
	_ = db.DB.AutoMigrate(&Video{})
	redb.RedisInit()
}
