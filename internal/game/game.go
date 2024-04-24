package game

import (
	"net/http"
	"os"

	accountv1 "github.com/VoidMesh/backend/api/gen/go/account/v1"
	"github.com/VoidMesh/backend/api/gen/go/account/v1/accountv1connect"
	characterv1 "github.com/VoidMesh/backend/api/gen/go/character/v1"
	"github.com/VoidMesh/backend/api/gen/go/character/v1/characterv1connect"
)

type Game struct {
	Services Services

	Account   *accountv1.Account
	Character *characterv1.Character
}

type Services struct {
	Account   accountv1connect.AccountServiceClient
	Character characterv1connect.CharacterServiceClient
}

func NewGame() *Game {
	return &Game{
		Services: Services{
			Account:   accountv1connect.NewAccountServiceClient(http.DefaultClient, os.Getenv("API_BACKEND_URL")),
			Character: characterv1connect.NewCharacterServiceClient(http.DefaultClient, os.Getenv("API_BACKEND_URL")),
		},
	}
}
