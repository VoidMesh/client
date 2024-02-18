package game

import (
	"context"
	"log"
	"time"

	"github.com/VoidMesh/backend/src/api/v1/account"
	"github.com/VoidMesh/backend/src/api/v1/character"
	"github.com/VoidMesh/backend/src/api/v1/resource"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Game struct {
	Client    *grpc.ClientConn
	Services  Services
	Account   *account.Account
	Character *character.Character
}

type Services struct {
	Account   account.AccountSvcClient
	Character character.CharacterSvcClient
	Resource  resource.ResourceSvcClient
}

func NewGame() *Game {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*2)
	conn, err := grpc.DialContext(ctx, "localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer cancel()

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &Game{
		Client: conn,
		Services: Services{
			Account:   account.NewAccountSvcClient(conn),
			Character: character.NewCharacterSvcClient(conn),
			Resource:  resource.NewResourceSvcClient(conn),
		},
	}
}
