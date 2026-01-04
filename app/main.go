package main

import (
	server "redis-go/app/pkg/server"
)

func main() {
	server := server.NewServer()
	server.Start()
}
