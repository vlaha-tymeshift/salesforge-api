package psql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"salesforge-api/internal/config"
)

func New(conf config.MySqlConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Pass,
		conf.Db,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create db: %v", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database unreachable: %v", err)
	}
	return db, nil
}
