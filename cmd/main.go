package main

import (
	api "diplom/api/proto"
	serv "diplom/pkg/serv"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	api.RegisterApiServer(s, &serv.ApiServ{})

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Starting server")
	go func() {
		if err := s.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()
	select {}
}
