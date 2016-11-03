package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UrbanDictionartyResponse struct {
	Tags       []string `json:"tags"`
	ResultType string   `json:"result_type"`
	List       []struct {
		Definition  string `json:"definition"`
		Permalink   string `json:"permalink"`
		ThumbsUp    int    `json:"thumbs_up"`
		Author      string `json:"author"`
		Word        string `json:"word"`
		Defid       int    `json:"defid"`
		CurrentVote string `json:"current_vote"`
		Example     string `json:"example"`
		ThumbsDown  int    `json:"thumbs_down"`
	} `json:"list"`
	Sounds []string `json:"sounds"`
}

//http://localhost:8080/word/?udword=slack
func wordHandler(w http.ResponseWriter, r *http.Request) {
	var word = r.URL.Query().Get("udword")
	url := fmt.Sprintf("http://api.urbandictionary.com/v0/define?term=%s", word)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()

	var record UrbanDictionartyResponse

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	fmt.Fprintf(w, "Tags = %s", record.Tags[0])

	// p, _ := loadPage(title)
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/word/", wordHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
