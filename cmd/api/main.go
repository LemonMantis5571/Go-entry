package main

import (
	"fmt"
	"go-course/internal/handlers"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error(err)
	}

}
