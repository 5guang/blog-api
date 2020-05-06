package user_service

import (
	"blog/dao"
	"blog/models/response"
	"blog/pkg/e"
	"blog/pkg/logging"
	"blog/pkg/setting"
	"blog/pkg/util"
)

type User struct {
	Username string
	Password string
	AdminPassword string
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

func (u *User) Login() (int, response.ResLoginData) {

	var (
		bol      bool
		adminBol bool
		// 默认是超级管理员
		isAdmin = 0
		err      error
		salt string
		adminPassword string
		resLogin = response.ResLoginData{}
	)
	bol = dao.IsExistByUsername(u.Username)
	if bol == false {
		return e.ERROR_USER_NOT_EXIST, resLogin
	}
	salt = dao.GetSalt()
	hash := dao.GetHashPassword()
	bol, err = util.PasswordVerify(*hash, u.Password, salt)
	if err != nil {
		return e.ERROR, resLogin
	}
	if bol == false {
		return e.ERROR_USER_WRONG_PASSWORD,resLogin
	}
	if u.AdminPassword != "" {
		// 校验超级管理员权限
		salt, err = util.Salt(setting.AppSetting.SaltLocalSecret)
		if err != nil {
			logging.Warn(err)
			return e.ERROR, resLogin
		}
		adminPassword, err = util.PasswordHash(setting.AppSetting.AdminPassword, salt)
		if err != nil {
			logging.Warn(err)
			return e.ERROR, resLogin
		}
		adminBol, err = util.PasswordVerify(adminPassword, u.AdminPassword, salt)
		if err != nil {
			logging.Warn(err)
			return e.ERROR, resLogin
		}
		if adminBol == false {
			return e.ERROR_USER_WRONG_ADMIN_PASSWORD,resLogin
		}
		isAdmin = 1
	}
	nickname := dao.GetNickname()
	resLogin = response.ResLoginData{
		Nickname: nickname,
		IsAdmin: isAdmin,
	}
	return e.SUCCESS, resLogin
}
