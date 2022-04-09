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
	DronesApi      DroneApi
	MedicationApi  MedicationApi
	DroneActionApi DroneActionApi
}

func StartServer(apis APIs) {
	serverPort := getenv.String("HTTP_SERVER_PORT", "9090")
	port := flag.String("port", serverPort, "Port to listen on")
	flag.Parse()

	r := mux.NewRouter().PathPrefix("/api/").Subrouter()
	r = r.StrictSlash(true)

	droneRouter := r.PathPrefix("/drone").Subrouter()
	droneRouter.Use(auth.Auth)
	droneRouter.StrictSlash(true)
	droneRouter.HandleFunc("/", apis.DronesApi.Create).Methods("POST")
	droneRouter.HandleFunc("/load", apis.DroneActionApi.Load).Methods("POST")
	droneRouter.HandleFunc("/medication", apis.DroneActionApi.GetLoadedMedicationItems).Methods("GET")
	droneRouter.HandleFunc("/available", apis.DroneActionApi.GetAvailableDrones).Methods("GET")

	medicationRouter := r.PathPrefix("/medication").Subrouter()
	medicationRouter.Use(auth.Auth)
	medicationRouter.StrictSlash(true)
	medicationRouter.HandleFunc("/", apis.MedicationApi.Create).Methods("POST")

	start(*port, r)
}

func start(port string, r http.Handler) {
	loggingRouter := loggingHandler(r)
	log.Println("listening on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, loggingRouter))
}

func loggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}
