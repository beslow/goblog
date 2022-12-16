package test

import (
	"database/sql"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/beslow/goblog/initialize"
	"github.com/go-testfixtures/testfixtures/v3"
)

//go:generate cp -r ../views .
//go:generate cp -r ../static .
//go:embed views/* static/**
var f embed.FS

var (
	db       *sql.DB
	fixtures *testfixtures.Loader
)

func TestMain(m *testing.M) {
	var err error

	// Open connection to the test database.
	// Do NOT import fixtures in a production database!
	// Existing data would be deleted.
	dir, _ := os.Getwd()

	if !strings.HasSuffix(dir, "test") {
		panic("wrong position")
	}

	cf := config.ReadFromYaml(filepath.Join(dir, "config.yml"))

	db, err = sql.Open("mysql", cf.Databases["default"].GetDSN())
	if err != nil {
		fmt.Printf("open mysql fail: %#v\n", err)
		os.Exit(1)
	}

	fixtures, err = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory("fixtures"),
	)
	if err != nil {
		fmt.Printf("new test fixtures fail: %#v\n", err)
		os.Exit(1)
	}

	os.Setenv("GO_TEST", "1")

	os.Exit(m.Run())
}

func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		fmt.Printf("fixtures load data fail: %#v\n", err)
		os.Exit(1)
	}

	initialize.InitMySQL()

	conn := initialize.RedisPool.Get()
	if _, err := conn.Do("flushall"); err != nil {
		fmt.Printf("flush redis fail: %#v\n", err)
		os.Exit(1)
	}
}
