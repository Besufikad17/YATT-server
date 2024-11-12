package internal

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Ctx    context.Context
	DBConn *pgxpool.Pool
}

func NewDBUtil(conn *pgxpool.Pool, ctx context.Context) *DB {
	return &DB{
		Ctx:    ctx,
		DBConn: conn,
	}
}

func (db *DB) Migrate() {
	sqlFiles, _ := os.ReadDir("sql/migrations")

	for _, file := range sqlFiles {
		query, err := os.ReadFile(filepath.Join("sql/migrations", file.Name()))
		if err != nil {
			log.Fatal("Erro reading "+file.Name()+"file", err.Error())
		}

		log.Println("Migration started")
		_, err = db.DBConn.Exec(db.Ctx, string(query))
		if err != nil {
			log.Fatal("Error migrating database from "+file.Name()+" file", err.Error())
		}
		log.Println("Migration ended")
	}
}
