package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/cristianchaparroa/humanity/backend/pkg/websocket"
	_ "github.com/lib/pq"
)

// Server defines the methods to server the application
type Server interface {
	SetupDB()
	SetupRepositories()
	SetupRoutes()
	Run()
	Close()
}

// ChatServer tis the implementation of Server interface
type ChatServer struct {
	db *sql.DB
}

// NewChatServer returns a pointer to ChatServer
func NewChatServer() *ChatServer {
	return &ChatServer{}
}

// SetupDB is charge to initialize the database connection
func (s *ChatServer) SetupDB() {

	user := os.Getenv("USER_DB")
	pass := os.Getenv("PASSWORD_DB")
	dbName := os.Getenv("NAME_DB")
	host := os.Getenv("HOST_DB")

	datasource := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, pass, host, dbName)
	fmt.Println(datasource)

	db, err := sql.Open("postgres", datasource)

	if err != nil {
		panic(err)
	}

	s.db = db
}

// SetupRepositories ...
func (s *ChatServer) SetupRepositories() {}

// SetupRoutes setup the endpoints availables in the backend
func (s *ChatServer) SetupRoutes() {

	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		RoomHandler(pool, w, r)
	})

}

// Run start the server
func (s *ChatServer) Run() {
	http.ListenAndServe(":8080", nil)
}

// Close all resources opened in the server
func (s *ChatServer) Close() {
	s.db.Close()
}
