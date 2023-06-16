package main

import (
	"context"
	api "diplom/api/proto"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		errorLog.Fatalf("Client: %v\n", err)
	}
	client := api.NewApiClient(conn)

	res, err := client.GetHardwareValue(context.Background(), &api.HardwareRequest{HarwareId: 1, Token: ""})
	if err != nil {
		errorLog.Fatalf("Client: %v\n", err)
	}
	fmt.Println(res.Params)
	req := &api.UpdateRequest{HardwareId: 1, Token: "", Params: []*api.UpdateParams{
		&api.UpdateParams{ParamId: 1, ParamValue: 1}}}
	res2, err := client.UpdateParamValue(context.Background(), req)
	fmt.Println(res2)
}
