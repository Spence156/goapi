package main

import (
	"fmt"
	"net/http"

	"github.com/Spence156/goapi/internal/handlers"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

var endpoint string = "localhost:8000"

func main() {

	log.SetReportCaller(true)        // Enables logger
	var r *chi.Mux = chi.NewRouter() // Return pointer to mux type (Used to setup the API)
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")

	fmt.Println(`
 ______     ______        ______     ______   __    
/\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \   
\ \ \__ \  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \  
 \ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\ 
  \/_____/   \/_____/      \/_/\/_/   \/_/     \/_/ `)

	err := http.ListenAndServe(endpoint, r)
	if err != nil {
		log.Error(err)
	}
}
