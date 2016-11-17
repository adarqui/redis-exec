// from https://github.com/adarqui/redis-transfer

package main

import (
    goredis "gopkg.in/redis.v4"
)

type Redis_Server struct {
    client *goredis.Client
    host   string
    port   int
    db     int
    pass   string
}
