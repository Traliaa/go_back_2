package main

import (
	"context"
	"fmt"

	"github.com/Traliaa/go_back_2/internal/app"
	"github.com/Traliaa/go_back_2/internal/pkg/database"
)

func main() {
	ctx := context.Background()
	db, err := database.NewConnect(ctx, "")
	if err != nil {
		fmt.Println(err)
	}
	app.NewApp(db)
}
