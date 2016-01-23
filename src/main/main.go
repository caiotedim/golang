package main

import (
    "webapp"
    "flag"
)

var (
	port *int
	bind *string
)


func main() {
    webapp.Server(bind, port)
}

func init() {
	bind = flag.String("b", "127.0.1", "bind address")
    port = flag.Int("p", 8080, "port")
    flag.Parse()
}