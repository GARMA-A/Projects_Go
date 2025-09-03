package store

import (
	"database/sql"
	"fmt"
	"io/fs"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("db : open %w", err)
	}
	fmt.Println("Connected to the database successfully")
	return db, nil
}

func MigrateFs(db *sql.DB, migrationFS fs.FS, dir string) error {
	goose.SetBaseFS(migrationFS)
	defer func() {
		goose.SetBaseFS(nil)
	}()
	return Migrate(db, dir)
}

func Migrate(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("db : migrate : set dialect %w", err)
	}
	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("db : migrate : up %w", err)
	}
	return nil
}
