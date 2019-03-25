package bancho

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello!")
	})
	return
}
