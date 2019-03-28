package bancho

import (
	"log"
	"net/http"
	"time"

	"github.com/nsogame/common"
)

type BanchoServer struct {
	config *Config
	db     *common.DB
	rds    *common.RedisAPI
	router http.Handler
}

func NewInstance(config *Config) (bancho *BanchoServer, err error) {
	// db
	db, err := common.ConnectDB(config.DbProvider, config.DbConnection)
	if err != nil {
		return
	}

	// redis
	rds := common.NewRedis(config.RedisAddr, config.RedisPass, config.RedisDB)

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
