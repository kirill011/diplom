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
	//ctx := metadata.AppendToOutgoingContext(context.Background(), "token", "")

	//res, err := client.Registration(ctx, &api.RegistrationRequest{Login: "Yasha", Password: "Lava"})
	//if err != nil {
	//	errorLog.Fatalf("Client: %v\n", err)
	//}
	//fmt.Println(res.MessageId)

	req := &api.UpdateRequest{HardwareId: 1, Token: "", Params: []*api.UpdateParams{
		&api.UpdateParams{ParamId: 1, ParamValue: 1}}}
	res2, err := client.UpdateParamValue(metadata.AppendToOutgoingContext(context.Background(), "token", "WWFzaGE6TGF2YQ=="), req)
	fmt.Println(res2)
}
