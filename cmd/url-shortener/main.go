package main

import (
	"log"
	"os"

	"github.com/idkwhyureadthis/url-shortener/shortener/internal/db"
	"github.com/idkwhyureadthis/url-shortener/shortener/internal/handler"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	dbConn := os.Getenv("DB_CONN")
	redirectUrl := os.Getenv("REDIRECT_SERVICE_URL")
	err := db.SetupMigrations(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	h := handler.New(dbConn, redirectUrl)

	log.Fatal(h.Start(port))
}
