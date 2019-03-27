package bancho

import (
	"net/http"
	"os"

	gorrilaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func (bancho *BanchoServer) IndexHandler(w http.ResponseWriter, r *http.Request) {
	header := w.Header()

	header.Set("cho-protocol", "19")
	header.Set("connection", "keep-alive")
	header.Set("keep-alive", "timeout=5, max=100")
	header.Set("content-type", "text/html; charset=utf-8")

	if r.Header.Get("osu-token") == "" && r.Header.Get("user-agent") == "osu!" {
		bancho.LoginHandler(w, r)
		return
	} else if r.Header.Get("osu-token") != "" {

	}
}

func (bancho *BanchoServer) Handlers() (router http.Handler) {
	r := mux.NewRouter()
	r.HandleFunc("/", bancho.IndexHandler)
	router = gorrilaHandlers.LoggingHandler(os.Stdout, r)
	return
}
