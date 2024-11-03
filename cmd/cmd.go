package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/pradeepbepari/golang_microservices/database"
	"github.com/pradeepbepari/golang_microservices/pkg/config"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel/trace"
)

type Server struct {
	Config   *config.Config
	Tracer   *trace.TracerProvider
	Database chan *sql.DB
	Server   chan *gin.Engine
	Wg       *sync.WaitGroup
}

func NewCommand(s *Server) *cobra.Command {
	rootCommand := &cobra.Command{
		Use:   "myapp",
		Short: "test and run concurrency",
		Run: func(cmd *cobra.Command, args []string) {
			s.Wg.Add(2)
			go connectDB(s.Database, s.Config, s.Wg)
			go connecterver(s.Server, s.Wg)
		},
	}
	return rootCommand
}

func connectDB(db chan<- *sql.DB, config *config.Config, wg *sync.WaitGroup) {
	defer wg.Done()
	cfg := &mysql.Config{
		User:                 config.DBUser,
		Passwd:               config.DBPassword,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", config.DBHost, config.DBPort),
		DBName:               config.DBName,
		AllowNativePasswords: true,
	}
	connection, err := database.ConnectionDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	db <- connection
}

func connecterver(server chan<- *gin.Engine, wg *sync.WaitGroup) {
	router := gin.Default()
	server <- router
	defer wg.Done()
}
