package driver

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var dbconn = &DB{}

const maxOpenDBConn = 5
const maxIdleDBConn = 5
const maxDBLifeTime = 5 * time.Minute

func ConnectPostgres(DSN string) (*DB, error) {
	d, err := sql.Open("pgx", DSN)
	if err != nil {
		return nil, err
	}
	d.SetMaxOpenConns(maxOpenDBConn)
	d.SetMaxIdleConns(maxIdleDBConn)
	d.SetConnMaxLifetime(maxDBLifeTime)

	err = testDB(d)
	if err != nil {
		return nil, err
	}

	dbconn.SQL = d
	return dbconn, nil
}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	fmt.Println("Pinged database successfully")

	return nil
}
