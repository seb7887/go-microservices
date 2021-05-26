package redis

import "github.com/gomodule/redigo/redis"

func Get(key string) ([]byte, error) {
	conn := Pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	return data, err
}

func Set(key string, value []byte) error {
	conn := Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	return err
}
