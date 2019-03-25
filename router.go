package bancho

import (
	"net/http"
	"os"

	"git.iptq.io/nso/bancho/handlers"

	gorrilaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	header := w.Header()

	header.Set("cho-protocol", "19")
	header.Set("connection", "keep-alive")
	header.Set("keep-alive", "timeout=5, max=100")
	header.Set("content-type", "text/html; charset=utf-8")

	if r.Header.Get("osu-token") == "" && r.Header.Get("user-agent") == "osu!" {
		handlers.LoginHandler(w, r)
		return
	}
}

func handler() (router http.Handler) {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	router = gorrilaHandlers.LoggingHandler(os.Stdout, r)
	return
}
