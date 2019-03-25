package bancho

import (
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
)

type BanchoServer struct {
	config *Config
	db     *gorm.DB
	router http.Handler
}

func NewInstance(config *Config) (bancho *BanchoServer, err error) {
	// db
	db, err := gorm.Open(config.DbProvider, config.DbConnection)
	if err != nil {
		return
	}

	// router
	router := handler()

	bancho = &BanchoServer{
		config: config,
		db:     db,
		router: router,
	}
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
