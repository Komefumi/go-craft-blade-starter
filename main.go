package main

import (
	"craft-blade/config"
	"craft-blade/server"
)

func main() {
	defer config.CloseDB()
	server.SetupServer()
}
