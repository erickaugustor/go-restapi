package main

import (
	"github.com/erickaugustor/go-restapi/server"
)

func main() {
	server := server.NewServer()

	server.Run()
}
