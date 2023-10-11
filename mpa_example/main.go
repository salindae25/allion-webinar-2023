package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"test/mpa_example/config"
	"test/mpa_example/data"
	"test/mpa_example/handlers"
)

var port = ":8080"

func AppRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("method is not valid"))
	})
	r.Get("/health", handlers.Repo.Health)
	r.Get("/", handlers.Repo.Home)
	r.Get("/about", handlers.Repo.About)
	r.Get("/country/{ccn}", handlers.Repo.Country)
	r.Post("/countries-json", handlers.Repo.CountriesWithJson)
	r.Get("/countries/{page}/page", handlers.Repo.Countries)

	fileServer := http.FileServer(http.Dir("./static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return r
}

func main() {
	countries, err := data.ReadJson()
	if err != nil {
		log.Fatal(err)
	}
	app := config.AppConfig{
		Data: countries,
	}
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	r := AppRouter()
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server started at port:%s", port)
}
