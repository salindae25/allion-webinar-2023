package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"test/mpa_example/config"
	"test/mpa_example/data"
	"test/mpa_example/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Health(w http.ResponseWriter, r *http.Request) {
	msg, _ := json.Marshal(map[string]string{"message": "Server is running.."})
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	w.Write(msg)
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	firstContries := (*m.App.Data)[:10]
	render.RenderPage(w, "home", firstContries)
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, "about", nil)
}

func (m *Repository) Country(w http.ResponseWriter, r *http.Request) {
	ccn := chi.URLParam(r, "ccn")
	if ccn == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Incorrect url parameter"))
		return
	}
	var country data.Country
	for _, c := range *m.App.Data {
		if c.Ccn3 == ccn {
			country = c
		}
	}
	if country.Name.Common == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Incorrect url parameter"))
		return
	}
	render.RenderPage(w, "country", country)
}

var elementsPerPage = 10

func (m *Repository) Countries(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	count := 0
	var err error
	if page != "" {
		count, err = strconv.Atoi(page)
		if err != nil {
			http.Error(w, "requested content is not present", http.StatusInternalServerError)
			return
		}
	}
	contentType := r.Header.Get("accept")
	switch contentType {
	case "application/json":
		CountriesInJson(w, count, *m.App.Data)
		return
	default:
		CountriesPage(w, count, *m.App.Data)
		return
	}
}

type CountryRequest struct {
	Page int `json:"page"`
}

func (m *Repository) CountriesWithJson(w http.ResponseWriter, r *http.Request) {
	var err error
	var reqBody CountryRequest

	err = json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	CountriesInJson(w, reqBody.Page, *m.App.Data)
}

func CountriesPage(w http.ResponseWriter, count int, data []data.Country) {
	countries, pageRange := getCountriesBasedOnStep(count, data)

	if countries == nil {
		http.Error(w, "requested content is not present", http.StatusInternalServerError)
		return
	}
	// w.Write([]byte(fmt.Sprintf("page 4 : %v", pageRange)))
	// return
	render.RenderPage(w, "countries", map[string]interface{}{"pages": pageRange, "data": countries})
}

func CountriesInJson(w http.ResponseWriter, count int, originalData []data.Country) {
	countries, pageRange := getCountriesBasedOnStep(count, originalData)
	if countries == nil {
		http.Error(w, "requested content is not present", http.StatusInternalServerError)
		return
	}
	respObject := struct {
		Data      []data.Country `json:"data"`
		Page      int            `json:"page"`
		NoOfPages int            `json:"noOfPages"`
	}{
		Data:      countries,
		Page:      count,
		NoOfPages: len(pageRange),
	}
	msg, err := json.Marshal(respObject)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(msg)
}

func getCountriesBasedOnStep(step int, data []data.Country) ([]data.Country, []int) {
	start := step * elementsPerPage
	end := start + elementsPerPage
	totalCountries := len(data)
	if start > totalCountries {
		return nil, nil
	}
	if end > totalCountries {
		end = totalCountries - 1
	}
	if end < start {
		return nil, nil
	}
	firstContries := data[start:end]
	noOfPages := totalCountries / elementsPerPage
	pageRange := []int{}
	for i := 0; i < noOfPages; i++ {
		pageRange = append(pageRange, i)
	}
	return firstContries, pageRange
}
