package client

import (
	"fmt"

	"github.com/Jopoleon/ActantTest/server"
	"google.golang.org/grpc"
)

func NewMetricsClient(host, port string) server.ActantTestClient {

	target := fmt.Sprintf("%s:%s", host, port)
	conn, e := grpc.Dial(target, grpc.WithInsecure())

	if e != nil {
		panic(e)
	}
	defer conn.Close()

	client := server.NewActantTestClient(conn)
	return client
}
