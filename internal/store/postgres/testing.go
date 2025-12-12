package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}
		db.Close()
	}

	// config := &config.Store{}
	// config.DatabaseURL = databaseURL
	// s := New(config)
	// if err := s.Open(); err != nil {
	// 	t.Fatal(err)
	// }

	// return s, func(tables ...string) {
	// 	if len(tables) > 0 {
	// 		if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
	// 			t.Fatal(err)
	// 		}
	// 	}
	// 	s.Close()
	// }
}
