package main

import (
	"context"

	"gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/config"
	"gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/infrastructure/mysql/db"
	"gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/server"
)

// @title Team D API Server
// @version バージョン(1.0)
// @description 説明
// @host localhost:8080
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := config.GetConfig()
	db.NewMainDB(conf.DB)
	db.NewReadDB(conf.ReadDB)

	server.Run(ctx, conf)
}
