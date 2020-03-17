package tools

import (
	"github.com/coding-codes/model"
	"github.com/coding-codes/utils"
)

type Options struct {
	Timeout bool
}

var defaultOptions = Options{
	Timeout: false,
}

type Option func(*Options)

func newOptions(opts ...Option) Options {
	opt := defaultOptions

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func SetTimeout(timeout bool) Option {
	return func(o *Options) {
		o.Timeout = timeout
	}
}

func SetKey(key string, value interface{}, opts ...Option) error {
	conn := model.RedisPool.Get()
	defer func() {
		if e := conn.Close(); e != nil {
			return
		}
	}()

	options := newOptions(opts...)
	if options.Timeout {
		_, err := conn.Do("SET", key, value, "EX", utils.RedisInfo.CacheTime)
		return err
	}
	_, err := conn.Do("SET", key, value)
	return err
}

func GetKey(key string) (data interface{}, err error) {
	conn := model.RedisPool.Get()
	defer func() {
		if e := conn.Close(); e != nil {
			return
		}
	}()

	data, err = conn.Do("GET", key)
	return
}

func DelKey(key string) error {
	conn := model.RedisPool.Get()
	defer func() {
		if e := conn.Close(); e != nil {
			return
		}
	}()

	_, err := conn.Do("DEL", key)
	return err
}
