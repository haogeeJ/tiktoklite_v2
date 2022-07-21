package dal

import (
	"TikTokLite_v2/favorite_comment/dal/db"
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	VideoID int64 `gorm:"index"`
	UserID  int64
	Content string `gorm:"index:idx_content,type:varchar(255);"`
}

func GetCommentRes(ctx context.Context, videoID int64) (rows *sql.Rows, err error) {
	session := db.DB.WithContext(ctx)
	//从原生sql改为gorm，方便修改迭代
	rows, err = session.Model(Comment{}).
		Select("comments.id", "comments.content", "comments.created_at", "users.id", "users.name").
		Joins("left join users on users.id=comments.user_id").
		Where("comments.deleted_at is null and comments.video_id=?", videoID).
		Order("comments.id desc").
		Rows()
	if err != nil {
		return nil, err
	}
	return rows, err
}

func (c *Comment) Create(ctx context.Context) error {
	session := db.DB.WithContext(ctx)
	err := session.Create(&c).Error
	if err != nil {
		return err
	}
	IncrCommentRedis(ctx, c.VideoID)
	return nil
}

func (c *Comment) Delete(ctx context.Context) error {
	session := db.DB.WithContext(ctx)
	err := session.Delete(&c).Error
	if err != nil {
		return err
	}
	DecrCommentRedis(ctx, c.VideoID)
	return nil
}

func (c *Comment) DeleteByUser(ctx context.Context) error {
	session := db.DB.WithContext(ctx)
	DecrCommentRedis(ctx, c.VideoID)
	return session.Where("id=? AND user_id=?", c.ID, c.UserID).Delete(&Comment{}).Error
	//return errors.New("invalid delete")
}

// GetCommentNum 获取评论数
func GetCommentNum(ctx context.Context, videoID int64) (count int64) {
	session := db.DB.WithContext(ctx)
	session.Model(&Comment{}).Where("video_id = ?", videoID).Count(&count)
	return
}
