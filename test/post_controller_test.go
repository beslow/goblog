package test

import (
	"database/sql"
	"embed"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	txdb "github.com/DATA-DOG/go-txdb"
	_ "github.com/GoAdminGroup/go-admin/adapter/gin"              // web framework adapter
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql" // sql driver
	_ "github.com/GoAdminGroup/themes/adminlte"                   // ui theme
	"github.com/beslow/goblog/initialize"
	"github.com/beslow/goblog/router"
	"github.com/romanyx/polluter"
	"github.com/stretchr/testify/assert"
)

//go:generate cp -r ../views .
//go:generate cp -r ../static .
//go:embed views/* static/**
var f embed.FS

func TestPostIndex(t *testing.T) {
	t.Parallel()
	defer exampleSuite(t)()

	initialize.InitMySQL()
	router.TemplateFs = f
	app, _ := router.SetRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/blog", nil)
	app.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func init() {
	txdb.Register("mysqltx", "mysql", "root:root@tcp(127.0.0.1:3306)/test_goblog")
}

func exampleSuite(t *testing.T) func() error {
	db, cleanup := prepareMySQLDB(t)
	seed, err := os.Open("seed.yaml")
	if err != nil {
		t.Fatalf("failed to open seed file: %s", err)
	}
	defer seed.Close()
	p := polluter.New(polluter.MySQLEngine(db))
	if err := p.Pollute(seed); err != nil {
		t.Fatalf("failed to pollute: %s", err)
	}
	return cleanup
}

func prepareMySQLDB(t *testing.T) (db *sql.DB, cleanup func() error) {
	cName := fmt.Sprintf("connection_%d", time.Now().UnixNano())
	db, err := sql.Open("mysqltx", cName)

	if err != nil {
		t.Fatalf("open mysqltx connection: %s", err)
	}

	return db, db.Close
}
