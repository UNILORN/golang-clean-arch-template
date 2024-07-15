package repository

import (
	"testing"

	"gopkg.in/testfixtures.v2"

	"gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/infrastructure/mysql/db"
	dbTest "gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/infrastructure/mysql/db/db_test"
	"gitlab.tokyo.optim.co.jp/hackathon/2024-summer/d-team/applications/backend-golang.git/internal/infrastructure/mysql/db/dbgen"
)

var (
	fixtures *testfixtures.Context
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
	dbTest.SetupTestDB("../db/schema/schema.sql")

	// テストデータの準備
	fixturePath := "../fixtures"
	fixtures, err = testfixtures.NewFolder(dbCon, &testfixtures.MySQL{}, fixturePath)
	if err != nil {
		panic(err)
	}

	q := dbgen.New(dbCon)
	db.SetQuery(q)
	db.SetDB(dbCon)

	// テスト実行
	m.Run()
}

func resetTestData(t *testing.T) {
	t.Helper()
	if err := fixtures.Load(); err != nil {
		t.Fatal(err)
	}
}
