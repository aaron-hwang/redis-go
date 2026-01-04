package server

type RedisServer struct {
}

func NewServer() *RedisServer {
	return &RedisServer{}
}

// Spins up a new Redis Server and to listen
func (rs RedisServer) run() {

}

// Spins up a new redis server instance, and listens for a client.
func (rs RedisServer) Start() {
	go rs.run()
}
