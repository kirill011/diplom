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
	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DB"))
	if err != nil {
		errorLog.Printf("Interceptor: %v\n", err)
		return nil, errors.New("Unable to connect to database")
	}
	messageId := uuid.NewV4().String()
	infoLog.Printf("%s : request start. MessageId : %v\n", info.FullMethod, messageId)

	rows, err := dbPool.Query(context.Background(), "select token from public.users where us.token = $1;")
	if err != nil {
		errorLog.Printf("Interceptor: %v\n", err)
		return nil, errors.New("SQL query select execution error")
	}
	if info.FullMethod != "Registration" {
		for rows.Next() {
			meta, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return nil, errors.New("could not grab metadata from context")
			}
			// Set ping-counts into the current ping value
			meta.Set("MessageId", messageId)
			// Metadata is sent on its own, so we need to send the header. There is also something called Trailer
			grpc.SendHeader(ctx, meta)
			// Last but super important, execute the handler so that the actualy gRPC request is also performed
			return handler(ctx, req)
		}
	} else {
		meta, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("could not grab metadata from context")
		}
		// Set ping-counts into the current ping value
		meta.Set("MessageId", messageId)
		// Metadata is sent on its own, so we need to send the header. There is also something called Trailer
		grpc.SendHeader(ctx, meta)
		// Last but super important, execute the handler so that the actualy gRPC request is also performed
		return handler(ctx, req)
	}
	return nil, errors.New("Not authenticated")
}
