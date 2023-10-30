package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"test/mpa_example/config"
	"test/mpa_example/data"
	"test/mpa_example/handlers"
)

var port = ":8080"

func AppRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		// AllowedOrigins: []string{"https://localhost:*", "http://*"},
		AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
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
	r.Post("/country/{ccn}/like", handlers.Repo.CountryLike)
	r.Post("/countries-json", handlers.Repo.CountriesWithJson)
	r.Get("/countries", handlers.Repo.Countries)

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
