package main

import (
	"flag"
	"time"

	"github.com/hprose/hprose-golang/rpc"
)

func say(in []byte) ([]byte, error) {
	args := &BenchmarkMessage{}
	args.Unmarshal(in)
	args.Field1 = "OK"
	args.Field2 = 100
	if *delay > 0 {
		time.Sleep(*delay)
	}
	return args.Marshal()
}

var (
	host  = flag.String("s", "127.0.0.1:8972", "listened ip and port")
	delay = flag.Duration("delay", 0, "delay to mock business processing")
)

func main() {
	flag.Parse()
	server := rpc.NewTCPServer("tcp://" + *host)
	server.AddFunction("say", say, rpc.Options{Simple: true})
	server.Start()
}
