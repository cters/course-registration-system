package initialize

import (
	"database/sql"
	"fmt"

	"github.com/QuanCters/backend/global"
	"github.com/QuanCters/backend/internal/errors"
	_ "github.com/lib/pq"
)

func InitPostgresql() {
	m := global.Config.Postgresql
	dsn := "postgres://%s:%s@%s:%v/%s?sslmode=%s"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname, m.SSLMode)
	db, err := sql.Open("postgres", s)
	errors.Must(global.Logger, err, "InitPostgresql initialization error")
	global.Logger.Info("Initializing PostgresQL Successfully")
	global.Pdbc = db
}

