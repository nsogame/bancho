package bancho

import (
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

type BanchoServer struct {
	config *Config
	db     *gorm.DB
	rds    *redis.Client
	router http.Handler
}

func NewInstance(config *Config) (bancho *BanchoServer, err error) {
	// db
	db, err := gorm.Open(config.DbProvider, config.DbConnection)
	if err != nil {
		return
	}

	// redis
	rds := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPass,
		DB:       config.RedisDB,
	})

	bancho = &BanchoServer{
		config: config,
		db:     db,
		rds:    rds,
	}

	// router
	router := bancho.Handlers()
	bancho.router = router

	return
}

func (bancho *BanchoServer) close() {
	bancho.db.Close()
}

func (bancho *BanchoServer) Run() {
	defer bancho.close()
	log.Println("starting...")
	server := &http.Server{
		Handler: bancho.router,
		Addr:    bancho.config.BindAddr,

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
