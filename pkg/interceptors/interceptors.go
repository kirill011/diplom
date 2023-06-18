package inter

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ServerAuthentication(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (response interface{}, err error) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	messageId := uuid.NewV4().String()
	infoLog.Printf("%s : request start. MessageId : %v\n", info.FullMethod, messageId)

	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DB"))
	if err != nil {
		errorLog.Printf("Interceptor: %v\n", err)
		return nil, errors.New("Unable to connect to database")
	}

	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		errorLog.Printf("Interceptor: %v\n", "could not grab metadata from context")
		return nil, errors.New("could not grab metadata from context")
	}

	token := meta.Get("token")
	if len(token) != 0 {
		infoLog.Printf("Interceptor token: %v\n", token[0])
	} else {
		errorLog.Printf("Interceptor: %v\n", "could not grab token from metadata")
		return nil, errors.New("could not grab token from metadata")
	}

	if token[0] == "" && info.FullMethod == "/api.api/Registration" {
		meta.Append("MessageId", messageId)
		// Metadata is sent on its own, so we need to send the header. There is also something called Trailer
		ctx = metadata.NewIncomingContext(ctx, meta)
		// Last but super important, execute the handler so that the actualy gRPC request is also performed
		return handler(ctx, req)
	}

	rows, err := dbPool.Query(context.Background(), "select token from public.users where token = $1;", token[0])
	if err != nil {
		errorLog.Printf("Interceptor: %v\n", err)
		return nil, errors.New("SQL query select execution error")
	}
	for rows.Next() {
		// Set ping-counts into the current ping value
		meta.Append("MessageId", messageId)
		// Metadata is sent on its own, so we need to send the header. There is also something called Trailer
		ctx = metadata.NewIncomingContext(ctx, meta)
		// Last but super important, execute the handler so that the actualy gRPC request is also performed
		return handler(ctx, req)
	}

	errorLog.Printf("Interceptor: %v\n", "Not authenticated")
	return nil, errors.New("Not authenticated")
}
