package database

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type Postgres struct {
	database Client
}

func NewConnect(ctx context.Context, connStr string) (*Postgres, error) {
	db, err := pgxpool.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}
	return &Postgres{database: db}, nil
}

func (c *Postgres) CreateUser() {}
func (c *Postgres) DeleteUser() {}
func (c *Postgres) SearchUser() {}

func (c *Postgres) CreateEnvironment()         {}
func (c *Postgres) DeleteEnvironment()         {}
func (c *Postgres) SearchEnvironment()         {}
func (c *Postgres) SearchByUserInEnvironment() {}
func (c *Postgres) SearchByEnvInEnvironment()  {}
