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
	ctx := metadata.AppendToOutgoingContext(context.Background(), "token", "Q3ZiOjU0Mw==")

	var mas []*api.UpdateParams
	mas[0] = &api.UpdateParams{ParamId: 62, ParamValue: 1}
	res, err := client.UpdateParamValue(ctx, &api.UpdateRequest{HardwareId: 60, Token: "U3RyZWtvemE6MTIzNDU2Nw==", Params: mas})
	if err != nil {
		errorLog.Printf("Client: %v\n", err)
	}
	fmt.Println(res)

	req := &api.HardwareIdRequest{Token: "Q3ZiOjU0Mw=="}
	res2, err := client.GetHardwareId(metadata.AppendToOutgoingContext(context.Background(), "token", "Q3ZiOjU0Mw=="), req)
	if err != nil {
		errorLog.Printf("Client: %v\n", err)
	}
	fmt.Println(res2)

	req2 := &api.HardwareRequest{HarwareId: 1, Token: "WWFzaGE6TGF2YQ=="}
	res3, err := client.GetHardwareValue(metadata.AppendToOutgoingContext(context.Background(), "token", "WWFzaGE6TGF2YQ=="), req2)
	if err != nil {
		errorLog.Printf("Client: %v\n", err)
	}
	fmt.Println(res3)
}
