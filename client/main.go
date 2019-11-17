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
	conn, err := grpc.Dial("192.168.0.94:7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := birthday.Server{}

	response, err := c.CheckBirthday(context.Background(), &birthday.Date{Day: 17, Month: 11, Year: 1999})
	if err != nil {
		log.Fatalf("error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %v", response.Age)
}
