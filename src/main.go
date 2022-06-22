package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"ldap-integration/src/ldap"
	"ldap-integration/src/server"
	"ldap-integration/src/user"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

func main() {
	db, err := connectDatabase()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	m, err := migrate.NewWithDatabaseInstance("file://./database/migrations", "postgres", driver)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	ldapService := ldap.NewLDAPService()
	userRepo := user.NewRepo(db)
	userService := user.NewService(userRepo, ldapService)

	ctx := context.Background()

	handler := server.NewHTTPServer(ctx, userService)

	fmt.Println("App listening on localhost:3000")
	http.ListenAndServe("127.0.0.1:3000", handler)
}

func connectDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgresql://postgres:postgres@localhost:5432/main?sslmode=disable")

	if err != nil {
		return db, err
	}

	db.SetConnMaxLifetime(time.Minute * time.Duration(20))
	db.SetMaxIdleConns(30)
	db.SetMaxOpenConns(50)

	var count int
	err = db.QueryRow("SELECT 1 AS count;").Scan(&count)
	if err != nil {
		return db, err
	}
	if count != 1 {
		return nil, errors.New("failed to test database connection")
	}
	return db, err
}
