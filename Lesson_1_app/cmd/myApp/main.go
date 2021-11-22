package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Traliaa/go_back_2/api/server"
	"github.com/Traliaa/go_back_2/api/server/handler"
	"github.com/Traliaa/go_back_2/config"
	"github.com/Traliaa/go_back_2/version"
)

func main() {
	launchMode := config.ENV(os.Getenv("LAUNCH_MODE"))
	if len(launchMode) == 0 {
		launchMode = config.LocalEnv
	}
	log.Printf("LAUNCH MODE: %v", launchMode)

	cfg, err := config.Load(launchMode)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CONFIG: %+v", cfg)
	info := handler.VersionInfo{
		Version: version.Version,
		Commit:  version.Commit,
		Build:   version.Build,
	}
	srv := server.NewServer(info, cfg.Port)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		srv.Serve(ctx)
	}()

	osSigChan := make(chan os.Signal, 1)
	signal.Notify(osSigChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-osSigChan
	log.Println("OS interrupting signal has received")

	cancel()

}

//ctx := context.Background()
//db, err := database.NewConnect(ctx, "")
//if err != nil {
//	fmt.Println(err)
//}
//app.NewApp(db)
