package dal

import (
	"TikTokLite_v2/favorite_comment/dal/db"
	"context"
	"database/sql"
	"errors"
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	VideoID int64 `gorm:"index:idx_video_user_id"`
	UserID  int64 `gorm:"index:idx_video_user_id"`
}

// Create 插入新纪录
func (f *Favorite) Create(ctx context.Context) error {
	session := db.DB.WithContext(ctx)
	return session.Create(&f).Error
}

// UniqueInsert 判断是否已经点赞，若未点赞，进行点赞并redis计数
func (f *Favorite) UniqueInsert(ctx context.Context) error {
	session := db.DB.WithContext(ctx)
	var FirstRes Favorite
	_ = session.Model(&Favorite{}).
		Where("video_id = ? and user_id = ?", f.VideoID, f.UserID).
		First(&FirstRes).Error
	if FirstRes.ID != 0 {
		return errors.New("repeat favorite")
	}
	err := f.Create(ctx)
	if err != nil {
		return err
	}
	IncrFavoriteRedis(ctx, f.VideoID)
	return nil
}

// Delete 根据userId和VideoId删除记录
func (f *Favorite) Delete(ctx context.Context) error {
	session := db.DB.WithContext(ctx)
	err := session.Where("user_id=? AND video_id=?", f.UserID, f.VideoID).
		Unscoped().Delete(&Favorite{}).Error
	if err != nil {
		return err
	}
	DecrFavoriteRedis(ctx, f.VideoID)
	return nil
}

// GetFavoriteNum count获取点赞数
func GetFavoriteNum(ctx context.Context, videoID int64) (count int64) {
	session := db.DB.WithContext(ctx)
	session.Model(&Favorite{}).Where("video_id = ?", videoID).Count(&count)
	return
}

// GetUserFavoriteNum 获取用户点赞视频总数
func GetUserFavoriteNum(ctx context.Context, userID int64) (count int64) {
	session := db.DB.WithContext(ctx)
	session.Model(&Favorite{}).Where("user_id = ?", userID).Count(&count)
	return
}

// IsFavorite 判断是否已点赞
func IsFavorite(ctx context.Context, userId, videoId int64) bool {
	session := db.DB.WithContext(ctx)

	var count int64
	err := session.Model(&Favorite{}).Where("video_id = ? and user_id = ?", videoId, userId).Count(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}

// GetFavoriteRes 联查获取user喜欢的所有video相关信息
func GetFavoriteRes(ctx context.Context, userID int64) (rows *sql.Rows, err error) {
	//rows, err = db.DB.Raw("select favorites.video_id,videos.author_id,videos.play_url,videos.cover_url,videos.title "+
	//	"FROM favorites INNER JOIN videos On favorites.video_id = videos.id "+
	//	"WHERE favorites.deleted_at is null and favorites.user_id = ?", userID).Rows()
	session := db.DB.WithContext(ctx)
	rows, err = session.Model(Favorite{}).
		Select("favorites.video_id", "videos.author_id", "videos.play_url", "videos.cover_url",
			"videos.title", "users.name").
		Joins("join videos on favorites.video_id=videos.id").
		Joins("join users on users.id=videos.author_id").
		Where("favorites.deleted_at is null and favorites.user_id=?", userID).
		Rows()
	if err != nil {
		return nil, err
	}
	return
}
