package service

import (
	"fmt"
	"github.com/coding-codes/model"
	"github.com/coding-codes/utils"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	ID           int    `json:"id" db:"id" form:"id"`
	Username     string `json:"username" db:"username" form:"username"`
	Password     string `json:"password" db:"password" form:"password"`
	Introduction string `json:"introduction" db:"introduction" form:"introduction"`
	Avatar       string `json:"avatar" db:"avatar" form:"avatar"`
	Nickname     string `json:"nickname" db:"nickname" form:"nickname"`
	About        string `json:"about" db:"about" form:"about"`
}

var db = model.DB
var jwtSecret = []byte(utils.AppInfo.JwtSecret)

type CustomClaims struct {
	username string
	jwt.StandardClaims
}

func (u User) CheckAuth() bool {
	var count int
	// 把 SQL 查询的结果赋给 count ,匹配为 1 ，不匹配为 0；
	if e := db.Get(&count, "select count(1) from blog_user where username=? and password=?", u.Username, utils.EncodeMD5(u.Password)); e != nil {
		return false
	}
	return count > 0
}

func (u User) GenToken() (string, error) {
	claims := CustomClaims{u.Username, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * time.Duration(utils.AppInfo.TokenTimeout)).Unix(),
		Id:        fmt.Sprintf("%v", time.Now().UnixNano()),
	}}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, e := tokenClaims.SignedString(jwtSecret)
	return token, e
}

func GetUser() (User, error) {
	var user User
	e := db.Get(&user, "select avatar,introduction,nickname from blog_user")
	return user, e
}

func GetAbout() (string, error) {
	var s string
	e := db.Get(&s, "select about from blog_user where username='admin'")
	return s, e
}

func (u User) EditAbout() error {
	_, e := db.Exec("update blog_user set about=? where username='admin'", u.About)
	return e
}

func (u User) ResetPassword() error {
	_, e := db.Exec("update blog_user set password=? where username='admin'", utils.EncodeMD5(u.Password))
	return e
}

func (u User) EditUser() error {
	_, e := db.Exec("update blog_user set avatar=?,introduction=?,nickname=?", u.Avatar, u.Introduction, u.Nickname)
	return e
}

func SetAvatarUrl(url string) error {
	_, e := db.Exec("update blog_user set avatar=? where username='admin'", url)
	return e
}
