package game

import (
	"context"
	"log"
	"time"

	"github.com/VoidMesh/backend/src/api/v1/account"
	"github.com/VoidMesh/backend/src/api/v1/character"
	"github.com/VoidMesh/backend/src/api/v1/inventory"
	"github.com/VoidMesh/backend/src/api/v1/resource"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Game struct {
	Services Services

	Account   *account.Account
	Character *character.Character
	Inventory *inventory.Inventory
}

type Services struct {
	Client *grpc.ClientConn

	Account   account.AccountSvcClient
	Character character.CharacterSvcClient
	Inventory inventory.InventorySvcClient
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
		Services: Services{
			Client:    conn,
			Account:   account.NewAccountSvcClient(conn),
			Character: character.NewCharacterSvcClient(conn),
			Inventory: inventory.NewInventorySvcClient(conn),
			Resource:  resource.NewResourceSvcClient(conn),
		},
	}
}
