package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bukhavtsov/museum/back-end/db"
	"github.com/bukhavtsov/museum/back-end/pkg/api"
	"github.com/bukhavtsov/museum/back-end/pkg/data"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	serverEndpoint = os.Getenv("SERVER_ENDPOINT")

	host     = os.Getenv("DB_USERS_HOST")
	port     = os.Getenv("DB_USERS_PORT")
	user     = os.Getenv("DB_USERS_USER")
	dbname   = os.Getenv("DB_USERS_DBNAME")
	password = os.Getenv("DB_USERS_PASSWORD")
	sslmode  = os.Getenv("DB_USERS_SSL")
)

func init() {
	if serverEndpoint == "" {
		serverEndpoint = "localhost:8080"
	}
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "1001"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "postgres"
	}
	if password == "" {
		password = "postgres"
	}
	if sslmode == "" {
		sslmode = "disable"
	}
}

func main() {
	r := mux.NewRouter()
	conn := db.GetConnection(host, port, user, dbname, password, sslmode)
	defer conn.Close()
	api.ServerArtifactResource(r, data.NewArtifactData(conn))
	handler := cors.AllowAll().Handler(r)
	log.Println("serving server at ", serverEndpoint)
	log.Fatal(http.ListenAndServe(serverEndpoint, handler))
}
