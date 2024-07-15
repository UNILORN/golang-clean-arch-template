package schema

import (
	"fmt"
	"log"
	"strconv"

	"github.com/sqldef/sqldef"
	"github.com/sqldef/sqldef/database"
	"github.com/sqldef/sqldef/database/mysql"
	"github.com/sqldef/sqldef/parser"
	"github.com/sqldef/sqldef/schema"
	"github.com/urfave/cli"

	"gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/config"
)

// urfave/cli経由で実行する
func Migrate(cCtx *cli.Context) error {
	schemaFile := cCtx.Args().Get(0)
	// データベースへの接続情報を設定
	dbCfg := config.GetConfig().DB

	port, _ := strconv.Atoi(dbCfg.Port)
	db, err := mysql.NewDatabase(database.Config{
		Host:     dbCfg.Host,
		Port:     port,
		User:     dbCfg.User,
		Password: dbCfg.Password,
		DbName:   dbCfg.Name,
	})
	if err != nil {
		log.Fatal(err)
	}
	desiredDDLs, err := sqldef.ReadFile(schemaFile)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %s", err)
	}

	dryRun := true
	if cCtx.Args().Get(1) == "apply" {
		dryRun = false
	}
	options := &sqldef.Options{
		DesiredDDLs:     desiredDDLs,
		DryRun:          dryRun,
		EnableDropTable: true,
	}

	sp := database.NewParser(parser.ParserModeMysql)
	sqldef.Run(schema.GeneratorModeMysql, db, sp, options)
	return nil
}
