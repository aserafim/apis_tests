package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Page struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Data       []struct {
		Name           string   `json:"name"`
		NativeName     string   `json:"nativeName"`
		TopLevelDomain []string `json:"topLevelDomain"`
		Alpha2Code     string   `json:"alpha2Code"`
		NumericCode    string   `json:"numericCode"`
		Alpha3Code     string   `json:"alpha3Code"`
		Currencies     []string `json:"currencies"`
		CallingCodes   []string `json:"callingCodes"`
		Capital        string   `json:"capital"`
		AltSpellings   []string `json:"altSpellings"`
		Relevance      string   `json:"relevance"`
		Region         string   `json:"region"`
		Subregion      string   `json:"subregion"`
		Language       []string `json:"language"`
		Languages      []string `json:"languages"`
		Translations   struct {
			De string `json:"de"`
			Es string `json:"es"`
			Fr string `json:"fr"`
			It string `json:"it"`
			Ja string `json:"ja"`
			Nl string `json:"nl"`
			Hr string `json:"hr"`
		} `json:"translations"`
		Population int       `json:"population"`
		Latlng     []float32 `json:"latlng"`
		Demonym    string    `json:"demonym"`
		Borders    []string  `json:"borders"`
		Area       int       `json:"area"`
		Gini       float64   `json:"gini,omitempty"`
		Timezones  []string  `json:"timezones"`
	} `json:"data"`
}

type Country struct {
	Name           string   `json:"name"`
	NativeName     string   `json:"nativeName"`
	TopLevelDomain []string `json:"topLevelDomain"`
	Alpha2Code     string   `json:"alpha2Code"`
	NumericCode    string   `json:"numericCode"`
	Alpha3Code     string   `json:"alpha3Code"`
	Currencies     []string `json:"currencies"`
	CallingCodes   []string `json:"callingCodes"`
	Capital        string   `json:"capital"`
	AltSpellings   []string `json:"altSpellings"`
	Relevance      string   `json:"relevance"`
	Region         string   `json:"region"`
	Subregion      string   `json:"subregion"`
	Language       []string `json:"language"`
	Languages      []string `json:"languages"`
	Translations   struct {
		De string `json:"de"`
		Es string `json:"es"`
		Fr string `json:"fr"`
		It string `json:"it"`
		Ja string `json:"ja"`
		Nl string `json:"nl"`
		Hr string `json:"hr"`
	} `json:"translations"`
	Population int       `json:"population"`
	Latlng     []float32 `json:"latlng"`
	Demonym    string    `json:"demonym"`
	Borders    []string  `json:"borders"`
	Area       int       `json:"area"`
	Gini       float64   `json:"gini"`
	Timezones  []string  `json:"timezones"`
}

func main() {
	var region, keyword string
	var found_countries []string
	fmt.Print("Informe uma região para consulta: ")
	fmt.Scan(&region)

	fmt.Print("Informe a palavra para busca: ")
	fmt.Scan(&keyword)

	//	fmt.Println(region, keyword)

	link := "https://jsonmock.hackerrank.com/api/countries/search?region=" + region + "&name=" + keyword + ""
	var total_pages int
	total_pages = 1
	for i := 1; i <= total_pages; i++ {
		resp, err := http.Get(link)
		if err != nil {
			log.Fatal("Erro ao realizar requisição: ", err)
		}
		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		var page Page
		err = decoder.Decode(&page)
		if err != nil {
			log.Fatal("Erro ao decodificar JSON: ", err)
		}

		total_pages = page.TotalPages

		for _, ctr := range page.Data {
			found_countries = append(found_countries, ctr.Name+", "+strconv.Itoa(ctr.Population))
		}
	}

	fmt.Println(found_countries)
}
