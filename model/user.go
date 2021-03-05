package model

import (
	"database/sql"
	"time"
)

type User struct {
	UserId           int16        `gorm:"column:user_id;primary_key"`
	UserName         string       `gorm:"column:user_name"`
	UserNickname     string       `gorm:"column:user_nickname"`
	UserPassword     string       `gorm:"column:user_password"`
	UserAvatar       string       `gorm:"column:user_avatar"`
	UserEmail        string       `gorm:"column:user_email"`
	UserTel          string       `gorm:"column:user_tel"`
	UserLoginIP      string       `gorm:"column:user_login_ip"`
	UserLoginArea    string       `gorm:"column:user_login_area"`
	UserLoginSystem  string       `gorm:"column:user_login_system"`  //ios,windows
	UserLoginDevice  string       `gorm:"column:user_login_device"`  //iphone,huawei
	UserLoginBrowser string       `gorm:"column:user_login_browser"` //Mobile Safari
	UserStatus       uint         `gorm:"column:user_status"`
	CreatedAt        time.Time    `gorm:"column:created_at"`
	UpdatedAt        time.Time    `gorm:"column:updated_at"`
	DeletedAt        sql.NullTime `gorm:"column:deleted_at"`
}
