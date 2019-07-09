package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	handleRequest()
}

//routing
func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllRequests)

	log.Fatal(http.ListenAndServe(":9091", nil))
}

//home page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Hey There</h1> <p>Wow</p>`)
}

//handle all requests
func returnAllRequests(w http.ResponseWriter, r *http.Request) {

	_ = json.NewEncoder(w).Encode(articles)

}

//article struct
type Article struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Content     string `json:"Content"`
}

var articles = []Article{
	{Title: "Why you lose", Description: "This is an article about losing", Content: "A long long time ago..."},
	{Title: "Oh No", Description: "This is an article about everything is wrong", Content: "I wish I had a dream"},
}
