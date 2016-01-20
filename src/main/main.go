package main

import (
    "webapp"
    "flag"
)

var (
	port *int
)


func main() {
    webapp.Server(port)
}

func init() {
    port = flag.Int("p", 8080, "port")
    flag.Parse()
}