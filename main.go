package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/k0kubun/pp"

	"github.com/Jopoleon/ActantTest/client"

	"github.com/Jopoleon/ActantTest/server"
	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = ":8899"
)

var (
	sendToPort = flag.String("send_port", "", "Port for sending received telemetry data")
)

func main() {
	flag.Parse()
	pp.Println(sendToPort)
	if sendToPort == nil || *sendToPort == "" {
		logrus.Error("you must define port for sending data with flag -send_port {0000}")
		return
	}

	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		logrus.Fatalf("failed to listen: %v\n", err)
	}

	grpcServer := grpc.NewServer()

	mc := client.NewMetricsClient("localhost", *sendToPort)

	server.RegisterActantTestServer(grpcServer, server.NewServer(mc))

	reflection.Register(grpcServer)

	fmt.Println("grpc server started on port: ", grpcPort)
	logrus.Fatal(grpcServer.Serve(listen))
}
