package bancho

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Bancho struct {
	config *Config
	db     *gorm.DB
	router *mux.Router
}

func NewInstance(config *Config) (bancho *Bancho, err error) {
	// db
	db, err := gorm.Open(config.DbProvider, config.DbConnection)
	if err != nil {
		return
	}

	// router
	router := Router()

	bancho = &Bancho{
		config: config,
		db:     db,
		router: router,
	}
	return
}

func (bancho *Bancho) close() {
	bancho.db.Close()
}

func (bancho *Bancho) Run() {
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
