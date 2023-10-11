package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func ReadJson() (*[]Country, error) {
	countries := []Country{}
	fileContent, err := os.Open("data/countries.json")
	if err != nil {
		log.Fatal(err)
		return &countries, err
	}
	defer fileContent.Close()
	byteCountries, err := ioutil.ReadAll(fileContent)
	if err != nil {
		return &countries, err
	}
	err = json.Unmarshal(byteCountries, &countries)
	if err != nil {
		return &countries, err
	}
	return &countries, nil
}
