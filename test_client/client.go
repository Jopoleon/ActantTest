package main

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/Jopoleon/ActantTest/server"
	"google.golang.org/grpc"
)

func NewMetricsClient(host, port string) server.ActantTestClient {

	target := fmt.Sprintf("%s:%s", host, port)
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		logrus.Fatal(err)
	}

	client := server.NewActantTestClient(conn)
	return client
}
