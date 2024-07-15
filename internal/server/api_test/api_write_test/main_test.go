//go:build integration_write

package api_write_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/testfixtures.v2"

	"gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/infrastructure/mysql/db"
	dbTest "gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/infrastructure/mysql/db/db_test"
	"gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/infrastructure/mysql/db/dbgen"
	"gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/presentation/settings"
	"gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/server/route"
)

var (
	fixtures *testfixtures.Context
	api      *gin.Engine
)

func TestMain(m *testing.M) {
	var err error

	// DBの立ち上げ
	resource, pool := dbTest.CreateContainer()
	defer dbTest.CloseContainer(resource, pool)

	// DBへ接続する
	dbCon := dbTest.ConnectDB(resource, pool)
	defer dbCon.Close()

	// テスト用DBをセットアップ
	dbTest.SetupTestDB("../../../infrastructure/mysql/db/schema/schema.sql")

	// テストデータの準備
	fixtures, err = testfixtures.NewFolder(
		dbCon,
		&testfixtures.MySQL{},
		"../../../infrastructure/mysql/fixtures",
	)
	if err != nil {
		panic(err)
	}

	q := dbgen.New(dbCon)
	db.SetQuery(q)
	db.SetReadQuery(q)
	db.SetDB(dbCon)
	db.SetReadDB(dbCon)

	api = settings.NewGinEngine()
	route.InitRoute(api)

	m.Run()
}

func resetTestData(t *testing.T) {
	t.Helper()
	if err := fixtures.Load(); err != nil {
		t.Fatal(err)
	}
}
