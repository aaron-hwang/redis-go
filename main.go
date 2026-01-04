package main

import (
	server "redis-go/pkg/redis-server"
)

func main() {
	server := server.NewServer()
	server.Start()
}
