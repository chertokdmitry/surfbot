package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func GetConn() *grpc.ClientConn {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial("127.0.0.1:5300", opts...)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	return conn
}
