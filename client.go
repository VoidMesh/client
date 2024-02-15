package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/VoidMesh/backend/src/api/v1" // Update this import path

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	greeter := pb.NewGreeterClient(conn)
	resource := pb.NewResourceServiceClient(conn)
	character := pb.NewCharacterServiceClient(conn)

	var command, subcommand string
	command = os.Args[1]
	subcommand = os.Args[2]

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	switch os.Args[1] {
	case "hello":
		name := os.Args[2]
		r, err := greeter.SayHello(ctx, &pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	case "resource":
		switch subcommand {
		case "read":
			name := os.Args[2]
			r, err := resource.ReadResource(ctx, &pb.ReadResourceRequest{Resource: &pb.Resource{Name: name}})
			if err != nil {
				log.Fatalf("could not get resource: %v", err)
			}
			log.Printf("Resource: %s", r.GetResource())
		case "list":
			r, err := resource.ListResources(ctx, &pb.ListResourcesRequest{})
			if err != nil {
				log.Fatalf("could not list resources: %v", err)
			}
			log.Printf("Resources: %s", r.Resources)
		}
	case "character":
		switch subcommand {
		case "create":
			name := os.Args[3]
			r, err := character.CreateCharacter(ctx, &pb.CreateCharacterRequest{Character: &pb.Character{Name: name}})
			if err != nil {
				log.Fatalf("could not get character: %v", err)
			}
			log.Printf("Character: %s", r.GetCharacter())
		case "read":
			name := os.Args[3]
			r, err := character.ReadCharacter(ctx, &pb.ReadCharacterRequest{Character: &pb.Character{Name: name}})
			if err != nil {
				log.Fatalf("could not get character: %v", err)
			}
			log.Printf("Character: %s", r.GetCharacter())
		}
	default:
		log.Fatalf("Invalid command: %s", command)
	}

	defer cancel()
}
