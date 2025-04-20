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
		ID        int    `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Avatar    string `json:"avatar"`
	} `json:"data"`
	Support struct {
		URL  string `json:"url"`
		Text string `json:"text"`
	} `json:"support"`
}

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

func main() {

	/*
		var pages []Page
		if err := decoder.Decode(&pages); err != nil {
			log.Fatal("Erro ao decodificar JSON: ", err)
		}
	*/
	var total_pages int
	total_pages = 1
	for page := 1; page <= total_pages; page++ {

		link := "https://reqres.in/api/users?page=" + strconv.Itoa(page)

		resp, err := http.Get(link)
		if err != nil {
			log.Fatal("Erro ao realizar requisição: ", err)
		}
		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)

		var current_page Page
		//var users []User
		err = decoder.Decode(&current_page)
		total_pages = current_page.TotalPages
		if err != nil {
			log.Fatal("Erro ao processar resposta: ", err)
		}
		fmt.Printf("Page number: %d\n", page)
		for _, user := range current_page.Data {
			fmt.Printf("Id: %d - Name: %s\n", user.ID, user.FirstName)
		}

	}

}
