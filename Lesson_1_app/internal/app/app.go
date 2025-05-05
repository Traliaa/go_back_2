package app

import (
	"github.com/Traliaa/go_back_2/api/server"
	"github.com/Traliaa/go_back_2/internal/app/environment"
	"github.com/Traliaa/go_back_2/internal/app/users"
)

type ImplementationUsers interface {
	Create()
	Search()
	Delete()
}
type ImplementationEnvironments interface {
	Create()
	Search()
	Delete()
	SearchByUser()
	SearchByEnv()
}

type ImplementationServer interface {
}
type app struct {
	Users        ImplementationUsers
	Environments ImplementationEnvironments
	Server       ImplementationServer
}

func NewApp(dbUser users.ImplementationDatabase, dbEnv environment.ImplementationDatabase) *app {

	return &app{
		Users:        users.NewUserRepositories(dbUser),
		Environments: environment.NewEnvironmentRepositories(dbEnv),
		Server:       server.NewServer(),
	}
}
