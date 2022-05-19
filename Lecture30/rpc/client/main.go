package main

import (
	"L30/news"
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := news.NewNewsServiceClient(conn)
	message := news.MessageClient{
		Title: "https://hacker-news.firebaseio.com",
	}
	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("response from Server %s", response.Title)
}
