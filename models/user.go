package models

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

type User struct {
	Model
	Username string `form:"username"`
	Password string `form:"password"`
	Nickname string `form:"nickname"`
	Realname string `form:"realname"`
	Avatar string `form:"avatar"`
	AuthKey string
	PasswordToken string
	Email string
	Status int
	Sex int
	LoginAt time.Time
	Token    string
}

func NewUser() (user *User) {
	return &User{}
}

func (u *User) Login() (findOne2 *User) {
	var findOne User
	Db.Where("username=?", u.Username).First(&findOne)
	h := md5.New()
	h.Write([]byte(u.Password))

	if findOne.Password == hex.EncodeToString(h.Sum(nil)) {
		findOne2 = &findOne
		return
	}

	return nil
}
