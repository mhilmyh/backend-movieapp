package cache

import "errors"

type Redis struct {}

func (c *Redis) Get(key string) (interface{}, error) {
	return nil, errors.New("not implemented yet")
}

func (c *Redis) Set(key string, i interface{}) error {
	return errors.New("not implemented yet")
}

func NewRedisConnection() *Redis {
	return &Redis{}
}