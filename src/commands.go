package src

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
curl -s "https://pokeapi.co/api/v2/location-area/?limit=20&offset=20" | jq
This is how it works, just use the offset to paginate the content
*/

type pokeAPI struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	commands := *cfg.commands

	for key, values := range commands {
		fmt.Printf("%s: %s\n", key, values.description)
	}
	return nil
}

func apiCall(cfg *config, jump string) error {

	url := ""

	if jump == "next" {
		url = cfg.next
	}

	if jump == "prev" {
		url = cfg.prev
	}

	res, err := http.Get(url)

	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	res.Body.Close()

	if res.StatusCode >= 299 {
		return errors.New(fmt.Sprintf("\nstatus code %d", res.StatusCode))
	}

	apiJson := pokeAPI{}

	if err := json.Unmarshal(body, &apiJson); err != nil {
		return err
	}

	for _, location := range apiJson.Results {
		fmt.Println(location.Name)
	}

	cfg.next = apiJson.Next
	cfg.prev = apiJson.Previous

	return nil
}

func commandMap(cfg *config) error {

	return apiCall(cfg, "next")

}

func commadMapBack(cfg *config) error {

	if cfg.prev == "" {
		return errors.New("you're on the first page")
	}

	return apiCall(cfg, "prev")
}
