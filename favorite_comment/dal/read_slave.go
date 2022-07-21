package dal

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"type:char(10)"`
}
type Video struct {
	gorm.Model
	AuthorId int64  `gorm:"index:idx_author_id"`
	Title    string `gorm:"type:varchar(255)" ,json:"title"`
	PlayUrl  string `gorm:"type:varchar(255)" ,json:"play_url"`
	CoverUrl string `gorm:"type:varchar(255)" ,json:"cover_url"`
}
