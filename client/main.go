package main

import (
	"fmt"
	"log"

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
	if response.Status {
		log.Printf("Congratulations for your Birthday! You're %v years old.", response.Age)
	} else {
		fmt.Printf("Sadly it's not your Bithday. You're current age is %v", response.Age)
	}
}
