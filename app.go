/* Firewall Request Management Webapp using Gorilla & HTML Templates */
package main

import (
	//Consider changing this if the project is renamed.
	"github.com/f47h3r/stoplight/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", handlers.Index).Methods("GET")
	rtr.HandleFunc("/index", handlers.Index).Methods("GET")
	rtr.HandleFunc("/about", handlers.Index).Methods("GET")
	rtr.HandleFunc("/req", handlers.Req)
	rtr.HandleFunc("/status", handlers.Status).Methods("GET")
	rtr.HandleFunc("/status/{fwRequestId}", handlers.StatusById).Methods("GET")
	rtr.HandleFunc("/approve", handlers.ApprovePage).Methods("GET")
	rtr.HandleFunc("/approve/{fwRequestId}/{queueName}", handlers.ApproveRequest).Methods("GET")
	rtr.HandleFunc("/audit", handlers.Audit).Methods("GET")
	rtr.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	rtr.NotFoundHandler = http.HandlerFunc(handlers.ErrorPage)
	http.Handle("/", rtr)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
