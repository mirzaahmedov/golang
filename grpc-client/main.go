package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mirzaahmedov/golang/api"
	"google.golang.org/grpc"
)

var address string = "localhost:8080"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := api.NewContactProviderClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.Get(ctx, &api.Id{
		Id: 11,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}
