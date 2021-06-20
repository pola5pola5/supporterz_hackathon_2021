package main

import (
	"flag"

	"github.com/littletake/supporterz_hackathon_2021/pkg/server"
)

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":1323", "tcp host:port to connect")
	flag.Parse()
}

func main() {
	server.Serve(addr)
}
