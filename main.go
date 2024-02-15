package main

import (
	"context"
	"log"
	"os"
	"time"

	apb "github.com/VoidMesh/backend/src/api/v1/account"
	cpb "github.com/VoidMesh/backend/src/api/v1/character"
	rpb "github.com/VoidMesh/backend/src/api/v1/resource"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	account := apb.NewAccountSvcClient(conn)
	character := cpb.NewCharacterSvcClient(conn)
	resource := rpb.NewResourceSvcClient(conn)

	var command, subcommand string
	command = os.Args[1]
	subcommand = os.Args[2]

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	switch os.Args[1] {
	case "resource":
		switch subcommand {
		case "get":
			name := os.Args[3]
			resp, err := resource.Get(ctx, &rpb.GetRequest{Resource: &rpb.Resource{Name: name}})
			if err != nil {
				log.Fatalf("could not get resource: %v", err)
			}
			log.Printf("Resource: %s", resp.GetResource())
		case "list":
			resp, err := resource.List(ctx, &rpb.ListRequest{})
			if err != nil {
				log.Fatalf("could not list resources: %v", err)
			}
			log.Printf("Resources: %s", resp.Resources)
		}

	case "character":
		switch subcommand {
		case "create":
			name := os.Args[3]
			resp, err := character.Create(ctx, &cpb.CreateRequest{Character: &cpb.Character{Name: name}})
			if err != nil {
				log.Fatalf("could not get character: %v", err)
			}
			log.Printf("Character: %s", resp.GetCharacter())
		case "read":
			name := os.Args[3]
			resp, err := character.Read(ctx, &cpb.ReadRequest{Character: &cpb.Character{Name: name}})
			if err != nil {
				log.Fatalf("could not get character: %v", err)
			}
			log.Printf("Character: %s", resp.GetCharacter())
		}

	case "account":
		switch subcommand {
		case "create":
			email := os.Args[3]
			resp, err := account.Create(ctx, &apb.CreateRequest{Email: email})
			if err != nil {
				log.Fatalf("could not create: %v", err)
			}
			log.Printf("Account: %s", resp.GetAccount())
		case "auth":
			email := os.Args[3]
			resp, err := account.Authenticate(ctx, &apb.AuthenticateRequest{Email: email})
			if err != nil {
				log.Fatalf("could not auth: %v", err)
			}
			log.Printf("Account: %s", resp.GetAccount())
		case "get":
			email := os.Args[3]
			resp, err := account.Get(ctx, &apb.GetRequest{Account: &apb.Account{Email: email}})
			if err != nil {
				log.Fatalf("could not get account: %v", err)
			}
			log.Printf("Account: %s", resp.GetAccount())
		}
	default:
		log.Fatalf("Invalid command: %s", command)
	}

	defer cancel()
}
