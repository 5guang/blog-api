package dao

import (
	"blog/pkg/logging"
	"github.com/jinzhu/gorm"
)

type User struct {
	Model
	// 当类型为零值时gorm默认不使用零值 例：username传空 默认sql语句就没有username=""条件 导致没有user查密码就可以登录
	// 使用指针类型可以避免
	Username *string `gorm:"type:varchar(20);"`
	Password *string `gorm:"type:varchar(255);"`
	NickName string `gorm:"type:varchar(15);"`
	Email string `gorm:"type:varchar(50);"`
	Salt string `gorm:"type:varchar(255);"`
	Roles    []Role `gorm:"many2many:user_role;association_autoupdate:false;"`
}

var user User

func CheckAuth(username string) (bol bool) {
	var auth User
	 DB.Select("id").Where(&User{Username:&username}).First(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}

func IsExistByUsername(username string) bool  {
	user = User{}
	err := DB.Where("username=? and deleted_on = 0", username).First(&user).Error
	// 如果没有记录 gorm默认会返回record not found
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Error(err)
		return false
	}
	if user.ID > 0 {
		return true
	}
	return false
}

func GetSalt() string  {
	return user.Salt
}

func GetHashPassword() *string  {
	return user.Password
}

func GetNickname() string  {
	return user.NickName
}

func AddUser(username, password, nickname, email, salt string) error {
	err := DB.Create(&User{
		Username: &username,
		Password: &password,
		NickName: nickname,
		Email: email,
		Salt:     salt,
	}).Error
	if err != nil {
		return err
	}
	return nil
}