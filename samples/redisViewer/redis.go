//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func newRedisConn(host, port, password string) (redis.Conn, error) {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return nil, err
	}
	if password != "" {
		if _, err = conn.Do("AUTH", password); err != nil {
			conn.Close()
			return nil, err
		}
	}
	return conn, nil
}
