package dal

import (
	"TikTokLite_v2/user_follow/user/dal/db"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"index:idx_name,type:char(10);"`
	Password string `gorm:"type:char(32);"`
}

//IUserRepository 对用户表操作接口
type IUserRepository interface {
	Insert(ctx context.Context, u *User) error
	GetById(ctx context.Context, u *User, id uint) error
	GetByName(ctx context.Context, u *User, username string) error
	Update(ctx context.Context, u *User) error
	IsExists(ctx context.Context, username string) error
}

//UserManagerRepository 实现了IUserRepository接口
type UserManagerRepository struct {
	db *gorm.DB
}

func NewUserManagerRepository() *UserManagerRepository {
	return &UserManagerRepository{db.DB}
}

//Insert 插入一个user实例
func (r *UserManagerRepository) Insert(ctx context.Context, u *User) error {
	session := r.db.WithContext(ctx)
	return session.Create(u).Error
}

//GetById 根据id获取实例
func (r *UserManagerRepository) GetById(ctx context.Context, u *User, id uint) error {
	session := r.db.WithContext(ctx)
	return session.Where("id=?", id).First(u).Error
}

//GetByName 根据username获取实例
func (r *UserManagerRepository) GetByName(ctx context.Context, u *User, username string) error {
	session := r.db.WithContext(ctx)
	return session.Where("name=?", username).First(u).Error
}

//Update 更新
func (r *UserManagerRepository) Update(ctx context.Context, u *User) error {
	session := r.db.WithContext(ctx)
	return session.Save(u).Error
}

//IsExists 判断username是否已存在表中，存在的话返回error
func (r *UserManagerRepository) IsExists(ctx context.Context, username string) error {
	session := r.db.WithContext(ctx)
	err := session.Where("name=?", username).Take(&User{}).Error
	if err == nil {
		return errors.New(fmt.Sprintf("The user already exists with UserName:%s", username))
	} else {
		return nil
	}
}
