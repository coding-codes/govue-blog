package model

import (
	"fmt"
	"github.com/coding-codes/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	"time"
)

var DB *sqlx.DB
var RedisPool *redis.Pool

func init() {
	DB = sqlx.MustConnect(utils.DBInfo.Mode, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4",
		utils.DBInfo.User,
		utils.DBInfo.Password,
		utils.DBInfo.Host,
		utils.DBInfo.Port,
		utils.DBInfo.DBName,
	))

	DB.SetMaxIdleConns(100)
	DB.SetMaxIdleConns(10)
	// 创建 redis 连接池
	redisAddr := fmt.Sprintf("%s:%s", utils.RedisInfo.Host, utils.RedisInfo.Post)
	RedisPool = &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial(
				"tcp",
				redisAddr, redis.DialPassword(utils.RedisInfo.Password),
				redis.DialDatabase(utils.RedisInfo.DB),
				redis.DialConnectTimeout(time.Second*5))
		},
		MaxIdle:     20,
		MaxActive:   1000,
		IdleTimeout: time.Second * 100,
		//  应用发起连接之前做服务健康检查，检查失败直接断开连接
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
