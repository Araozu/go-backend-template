package src

import (
	"acide/src/database"
	"embed"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"net/http"
	"os"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Server struct {
	port int
	db   *gorm.DB
}

// Creates a new Server
func NewServer(static *embed.FS) *http.Server {
	portEnv := os.Getenv("PORT")
	port, err := strconv.Atoi(portEnv)
	if err != nil {
		panic(fmt.Sprintf("Error converting PORT env variable to int.\n%s\n%s", portEnv, err))
	}

	NewServer := &Server{
		port: port,
		db:   database.New(),
	}

	// TODO: Register DB schemas here
	// login.SetupSchema(NewServer.db)

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(static),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
