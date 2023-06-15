package main

import (
	api "diplom/api/proto"
	serv "diplom/pkg/serv"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	api.RegisterApiServer(s, &serv.ApiServ{})

	l, err := net.Listen("tcp", "77.222.42.182:8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Starting server")
	go func() {
		if err := s.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()
	select {}
}
