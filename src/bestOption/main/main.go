package main

import (
	"..//webapp"
	"flag"
	"log"
)

var (
	port *int
	bind *string
	key  *[]byte
)

func main() {
	log.Printf("=== Starting Server ===")
	log.Printf("Bind address: [%s]", *bind)
	log.Printf("Port address: [%d]", *port)
	webapp.Server(bind, port)
}

func init() {

	bind = flag.String("b", "127.0.0.1", "bind address")
	port = flag.Int("p", 8080, "port")
	flag.Parse()
}