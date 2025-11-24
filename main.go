package main

import (
	"ToDO/config"
	"ToDO/pkg/task"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	router.Mount("/api/v1/tasks", task.Routes(configuration))
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./vue/index.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		t.Execute(w, nil)
	})
	return router
}

func main() {
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error:", err)
	}

	router := Routes(configuration)

	log.Println("Serving on :8080")
	http.ListenAndServe(":8080", router)
}
