package main

import (
	"database/sql"
	"db_lab01/models"
	"db_lab01/routers"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)

const (
	host           = "localhost"
	dbPort         = 5432
	user           = "postgres"
	password       = "newpassword"
	dbname         = "labs_db"
	migrationsUp   = "migrations/up.sql"
	migrationsDown = "migrations/down.sql"
	driver         = "postgres"
	maxConnections = 50
	apiPort        = ":5000"
)

func closeDb(db *os.File) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func applyMigrations(filename string, db *sql.DB) {
	migrationFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer closeDb(migrationFile)

	st, err := migrationFile.Stat()
	if err != nil {
		log.Fatal(err)
	}
	stArray := make([]byte, st.Size())

	_, err = migrationFile.Read(stArray)
	if err != nil {
		log.Fatal(err)
	}

	databaseScheme := string(stArray)
	_, err = db.Exec(databaseScheme)
	if err != nil {
		log.Fatal(err)
	}
}

func initDatabase(database *models.Database, withClear bool) *models.Env {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		database.Host, database.Port, database.User, database.Password, database.Name)
	db, err := sql.Open(database.Driver, connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	if withClear {
		applyMigrations(database.MigrationsDownFileName, db)
	}

	applyMigrations(database.MigrationsUpFileName, db)
	env := &models.Env{
		DB: db,
	}

	db.SetMaxOpenConns(maxConnections)
	log.Println("Successfully initialized database")
	return env
}

func main() {
	database := &models.Database{
		Host:                   host,
		Port:                   dbPort,
		User:                   user,
		Password:               password,
		Name:                   dbname,
		Driver:                 driver,
		MaxConnections:         maxConnections,
		MigrationsUpFileName:   migrationsUp,
		MigrationsDownFileName: migrationsDown,
	}
	env := initDatabase(database, true)

	router := routers.InitRouter(env)
	err := fasthttp.ListenAndServe(apiPort, router.HandleRequest)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening at localhost" + apiPort)
}
