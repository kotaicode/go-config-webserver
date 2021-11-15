package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler).Methods("GET")
	r.HandleFunc("/data", getConfigData).Methods("GET")
	staticFiles := http.Dir("./assets")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFiles))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	return r
}
func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	r := newRouter()
	log.Info().
		Int("Port", 8080).
		Msg("Starting server")
	http.ListenAndServe(":8080", r)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
