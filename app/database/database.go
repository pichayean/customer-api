package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // The database driver in use.
)

type Config struct {
	User       string
	Password   string
	Host       string
	Port       int
	Name       string
	DisableTLS bool
}

func Open(cfg Config) (*sqlx.DB, error) {
	sslmode := "require"
	if cfg.DisableTLS {
		sslmode = "disable"
	}
	var dataSoruce = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, sslmode)
	return sqlx.Connect("postgres", dataSoruce)
}

func NewTable(db *sqlx.DB, schema string) (int64, error) {
	return db.MustExec(schema).RowsAffected()
}

func DropTable(db *sqlx.DB, tableName string) (int64, error) {
	schema := fmt.Sprintf("DROP TABLE %s;", tableName)
	return db.MustExec(schema).RowsAffected()
}
