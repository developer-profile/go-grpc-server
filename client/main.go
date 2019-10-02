package main

import (
	"log"

	"github.com/daniel-vera-g/golang-grpc-server/api/birthday"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	// Initiate a connection with the server
	conn, err := grpc.Dial("localhost:7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := birthday.Server{}

	response, err := c.CheckBirthday(context.Background(), &birthday.Date{Day: 30, Month: 12, Year: 1999})
	if err != nil {
		log.Fatalf("error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Age)
}
