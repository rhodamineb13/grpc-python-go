package main

import (
	"log"

	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption

	conn, err := grpc.NewClient("localhost:50051", opts...)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	//client := request.NewUserFeaturesClient(conn)

}
