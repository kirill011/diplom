package main

import (
	"context"
	api "diplom/api/proto"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := api.NewApiClient(conn)

	res, err := client.GetHardwareValue(context.Background(), &api.HardwareRequest{HarwareId: 1, Token: ""})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Params)
	par := make([]*api.UpdateParams, 1, 1)
	par[0] = &api.UpdateParams{ParamId: 1, ParamValue: 1}
	res2, err := client.UpdateParamValue(context.Background(), &api.UpdateRequest{HardwareId: 1, Token: "", Params: []*api.UpdateParams{
		&api.UpdateParams{ParamId: 1, ParamValue: 1},
	}})
	fmt.Println(res2)
}
