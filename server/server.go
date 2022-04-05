package server

import (
	"flag"
	"log"
	"net/http"

	"drones.com/server/auth"
	"github.com/gorilla/mux"
	"github.com/ieee0824/getenv"
)

type APIs struct {
	DronesApi DroneApi
}

func StartServer(apis APIs) {
	serverPort := getenv.String("HTTP_SERVER_PORT", "9090")
	port := flag.String("port", serverPort, "Port to listen on")
	flag.Parse()

	r := mux.NewRouter().PathPrefix("/api/").Subrouter()
	r = r.StrictSlash(true)

	dronesRouter := r.PathPrefix("/drone").Subrouter()
	dronesRouter.Use(auth.Auth)
	dronesRouter.StrictSlash(true)
	dronesRouter.HandleFunc("/", apis.DronesApi.Create).Methods("POST")

	start(*port, r)
}

func start(port string, r http.Handler) {
	loggingRouter := loggingHandler(r)
	log.Fatal(http.ListenAndServe(":"+port, loggingRouter))
}

func loggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}
