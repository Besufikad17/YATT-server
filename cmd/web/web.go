package web

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Application struct {
	Ctx    context.Context
	DBConn *pgxpool.Pool
}

func Serve() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}

	ctx := context.Background()

	conn, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error connecting to database "+os.Getenv("DATABASE_URL"), err.Error())
	}

	app := &Application{
		DBConn: conn,
	}

	log.Println("Starting server on :4000")
	err = http.ListenAndServe(":4000", app.routes())
	if err != nil {
		log.Fatal("Error running server on :4000", err.Error())
	}
}
