package main

import (
	api "diplom/api/proto"
	inter "diplom/pkg/interceptors"
	serv "diplom/pkg/serv"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lmicroseconds)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lmicroseconds)

	s := grpc.NewServer(grpc.UnaryInterceptor(inter.ServerAuthentication))
	api.RegisterApiServer(s, &serv.ApiServ{})

	l, err := net.Listen("tcp", ":50800")
	if err != nil {
		errorLog.Fatalf("Server: %v\n", err)
	}
	infoLog.Print("Server: server starting")
	if err := s.Serve(l); err != nil {
		errorLog.Fatalf("Server: %v\n", err)
	}
}
