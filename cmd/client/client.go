package main

import (
	"context"
	api "diplom/api/proto"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := api.NewApiClient(conn)

	res, err := client.GetHardwareValue(context.Background(), &api.HardwareRequest{HarwareId: 1, Token: ""})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Params)
	res2, err := client.UpdateParamValue(context.Background(), &api.UpdateRequest{HardwareId: 1, Token: "", Params: []*api.UpdateParams{
		&api.UpdateParams{ParamId: 1, ParamValue: 1},
	}})
	fmt.Println(res2)
}
