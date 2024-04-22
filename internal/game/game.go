package game

import (
	"context"
	"log"
	"time"

	"github.com/VoidMesh/backend/pkg/api/account/v1"
	"github.com/VoidMesh/backend/pkg/api/character/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Game struct {
	Services Services

	Account   *account.Account
	Character *character.Character
}

type Services struct {
	Client *grpc.ClientConn

	Account   account.AccountSvcClient
	Character character.CharacterSvcClient
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
		},
	}
}
