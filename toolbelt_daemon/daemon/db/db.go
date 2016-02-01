package db

import "github.com/garyburd/redigo/redis"

var (
	conn redis.Conn
	connectionErr error
)

func Connect() {
	conn, connectionErr = redis.Dial("tcp", "redis:6379")
}

func Close() error {
	return conn.Close()
}

func Set(key string, value string) (reply interface{}, err error) {
	return conn.Do("SET", key, value)
}

func Get(key string) (string, error) {
	response, _ := redis.Values(conn.Do("GET", key))
	var value string

	_, err := redis.Scan(response, value)

	return value, err
}
