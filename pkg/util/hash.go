package util

import (
	"blog/pkg/setting"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"github.com/unknwon/com"
	"io"
	mt "math/rand"
	"strconv"
	"strings"
)


//加密密码
func PasswordHash(pass, salt string) (string, error) {

	saltSecret, err := saltSecret()
	if err != nil {
		return "", err
	}

	interaction := randInt(1, 10)

	hash, err := hash(pass, saltSecret, salt, int64(interaction))
	if err != nil {
		return "", err
	}
	interactionString := strconv.Itoa(interaction)
	delimiter := setting.AppSetting.Delimiter
	password := saltSecret + delimiter + interactionString + delimiter + hash + delimiter// + salt

	return password, nil

}

//校验密码是否有效
func PasswordVerify(hashing string, plainPass string, salt string) (bool, error) {
	data := trimSaltHash(hashing)

	interaction, _ := strconv.ParseInt(data["interaction_string"], 10, 64)

	has, err := hash(plainPass, data["salt_secret"], salt, int64(interaction))
	if err != nil {
		return false, err
	}
	delimiter := setting.AppSetting.Delimiter
	if (data["salt_secret"] + delimiter + data["interaction_string"] + delimiter + has + delimiter) == hashing {
		return true, nil
	} else {
		return false, nil
	}

}

func hash(pass string, saltSecret string, salt string, interaction int64) (string, error) {
	var passSalt = saltSecret + pass + salt + saltSecret + pass + salt + pass + pass + salt
	var i int
	saltLocalSecret := setting.AppSetting.SaltLocalSecret
	hashPass := saltLocalSecret
	hashStart := sha512.New()
	hashCenter := sha256.New()
	hashOutput := sha256.New224()
	stretchingPassword, err := com.StrTo(setting.AppSetting.StretchingPassword).Int()
	if err != nil {
		return "",err
	}
	i = 0
	for i <= stretchingPassword {
		i = i + 1
		_, err := hashStart.Write([]byte(passSalt + hashPass))
		if err != nil {
			return "", err
		}
		hashPass = hex.EncodeToString(hashStart.Sum(nil))
	}

	i = 0
	for int64(i) <= interaction {
		i = i + 1
		hashPass = hashPass + hashPass
	}

	i = 0
	for i <= stretchingPassword {
		i = i + 1
		_, err := hashCenter.Write([]byte(hashPass + saltSecret))
		if err != nil {
			return "", err
		}
		hashPass = hex.EncodeToString(hashCenter.Sum(nil))
	}
	if _, err := hashOutput.Write([]byte(hashPass + saltLocalSecret)); err != nil {
		return "", err
	}
	hashPass = hex.EncodeToString(hashOutput.Sum(nil))
	return hashPass, nil
}

func trimSaltHash(hash string) map[string]string {
	delimiter := setting.AppSetting.Delimiter
	str := strings.Split(hash, delimiter)

	return map[string]string{
		"salt_secret":       str[0],
		"interaction_string": str[1],
		"hash":              str[2],
		//"salt":              str[3],
	}
}
func Salt(secret string) (string, error) {
	saltSize, err := com.StrTo(setting.AppSetting.SaltSecret).Int()
	if err != nil {
		return "",err
	}
	ss, err := saltSecret()
	if err != nil {
		return "",err
	}
	secret += ss
	buf := make([]byte, saltSize, saltSize+md5.Size)
	_, err = io.ReadFull(rand.Reader, buf)
	if err != nil {
		return "", err
	}

	hash := md5.New()
	hash.Write(buf)
	hash.Write([]byte(secret))
	return hex.EncodeToString(hash.Sum(buf)), nil
}

func saltSecret() (string, error) {
	rb := make([]byte, randInt(10, 100))
	_, err := rand.Read(rb)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(rb), nil
}

func randInt(min int, max int) int {
	return min + mt.Intn(max-min)
}
