package api

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/cristianchaparroa/humanity/backend/core/websocket"
	"github.com/cristianchaparroa/humanity/backend/initializer"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// Server defines the methods to server the application
type Server interface {
	SetupDB()
	SetupRoutes()
	Run()
	Close()
}

// ChatServer tis the implementation of Server interface
type ChatServer struct {
	Router *gin.Engine
	db     *sql.DB
	gormDB *gorm.DB
}

// NewChatServer returns a pointer to ChatServer
func NewChatServer() *ChatServer {

	r := gin.Default()
	return &ChatServer{Router: r}
}

// SetupDB is charge to initialize the database connection
func (s *ChatServer) SetupDB() {

	user := os.Getenv("USER_DB")
	pass := os.Getenv("PASSWORD_DB")
	dbName := os.Getenv("NAME_DB")
	host := os.Getenv("HOST_DB")

	datasource := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", user, pass, host, dbName)
	fmt.Println(datasource)

	//db, err := sql.Open("postgres", datasource)

	db, err := gorm.Open("postgres", datasource)

	if err != nil {
		panic(err)
	}

	s.gormDB = db

	im := initializer.NewInitialzerManager()
	im.Run(s.gormDB)

}

// SetupRoutes setup the endpoints availables in the backend
func (s *ChatServer) SetupRoutes() {

	pool := websocket.NewChatPool()
	go pool.Start()

	store := sessions.NewCookieStore([]byte("secret"))
	s.Router.Use(sessions.Sessions("mysession", store))
	s.Router.Use(CORS())

	s.Router.POST("/api/login", LoginHandler(s.gormDB))
	s.Router.GET("/api/logout", LogoutHandler)

	s.Router.GET("/ws/room", func(c *gin.Context) {

		RoomHandler(c, pool, c.Writer, c.Request)
	})
}

// Run start the server
func (s *ChatServer) Run() {
	s.Router.Run(":8080")
}

// Close all resources opened in the server
func (s *ChatServer) Close() {
	s.db.Close()
}
