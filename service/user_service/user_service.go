package user_service

import (
	"blog/dao"
	"blog/pkg/e"
	"blog/pkg/setting"
	"blog/pkg/util"
)

type User struct {
	Username string
	Password string
	NickName string
	Email string
}

func (u *User) Register() int {
	var (
		bol      bool
		err      error
		salt     string
		password string
	)
	bol = dao.IsExistByUsername(u.Username)
	if bol == true {
		return e.ERROR_USER_ALREADY_EXIST
	}
	salt, err = util.Salt(setting.AppSetting.SaltLocalSecret)
	if err != nil {
		return e.ERROR
	}
	password, err = util.PasswordHash(u.Password, salt)
	if err != nil {
		return e.ERROR
	}
	err = dao.AddUser(u.Username, password, u.NickName, u.Email, salt)
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

func (u *User) Login() int {

	var (
		bol      bool
		err      error
		salt string
	)
	bol = dao.IsExistByUsername(u.Username)
	if bol == false {
		return e.ERROR_USER_NOT_EXIST
	}
	salt = dao.GetSalt()
	hash := dao.GetHashPassword()
	bol, err = util.PasswordVerify(*hash, u.Password, salt)
	if err != nil {
		return e.ERROR
	}
	if bol == false {
		return e.ERROR_USER_WRONG_PASSWORD
	}
	return e.SUCCESS
}
