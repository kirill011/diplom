package main

import (
	"context"
	api "diplom/api/proto"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		errorLog.Fatalf("Client: %v\n", err)
	}
	client := api.NewApiClient(conn)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "token", "")

	res, err := client.Registration(ctx, &api.RegistrationRequest{Login: "Yasha", Password: "Lava"})
	if err != nil {
		errorLog.Printf("Client: %v\n", err)
	}
	fmt.Println(res.MessageId)

	req := &api.HardwareIdRequest{Token: "WWFzaGE6TGF2YQ=="}
	res2, err := client.GetHardwareId(metadata.AppendToOutgoingContext(context.Background(), "token", "WWFzaGE6TGF2YQ=="), req)

	req2 := &api.HardwareRequest{HarwareId: 1, Token: "WWFzaGE6TGF2YQ=="}
	res3, err := client.GetHardwareValue(metadata.AppendToOutgoingContext(context.Background(), "token", "WWFzaGE6TGF2YQ=="), req2)
	fmt.Println(res3)
	fmt.Println(res2)
}
