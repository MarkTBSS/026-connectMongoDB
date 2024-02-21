package main

import (
	"context"
	"log"
	"os"

	"github.com/MarkTBSS/go-connectMongoDB/config"
	"github.com/MarkTBSS/go-connectMongoDB/pkg/database"
)

func main() {
	ctx := context.Background()
	_ = ctx
	// Initialize config
	cfg := config.LoadConfig(func() string {
		// [0]main.go, [1]./env/xxx/file.ext
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		return os.Args[1]
	}())
	// Check Config Object
	log.Println(cfg)
	// Database connection
	db := database.DatabaseConnect(ctx, &cfg)
	//defer db.Disconnect(ctx)
	log.Println(db)
}
