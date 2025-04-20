package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatal("Erro na requisição")
	}
	defer resp.Body.Close()

	/*
		ret, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Erro ao ler resposta")
		}
		fmt.Println(string(ret))*/

	decoder := json.NewDecoder(resp.Body)
	var posts []Post
	if err := decoder.Decode(&posts); err != nil {
		log.Fatal("Erro ao decodificar JSON:", err)
	}

	for i, post := range posts {
		fmt.Printf("%d - %s\n", post.ID, post.Title)
		if i >= 9 {
			break
		}
	}
}
