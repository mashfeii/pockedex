package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/mashfeii/pokedexcli/internal/pockecache"
)

const base_time = time.Millisecond * 50

var cache = pockecache.NewCache(base_time)

type Config struct {
	Next     string
	Previous string
}

type Response struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"Previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocations(conf *Config, forward bool) {
	URL := conf.Next
	if !forward {
		URL = conf.Previous
	}

	cache_resp, cache_exists := cache.Get(URL)

	new_slice := Response{}
	var json_err error

	if cache_exists {
		json_err = json.Unmarshal(cache_resp, &new_slice)
	} else {
		resp, err := http.Get(URL)
		if err != nil {
			log.Fatal(err)
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
		json_err = json.Unmarshal(body, &new_slice)

		cache.Add(URL, body)
	}

	if json_err != nil {
		log.Fatal(json_err)
	}

	if new_slice.Next != nil && *new_slice.Next != conf.Next {
		conf.Next = *new_slice.Next
	}
	if new_slice.Previous != nil && *new_slice.Previous != conf.Previous {
		conf.Previous = *new_slice.Previous
	}

	for _, loc := range new_slice.Results {
		fmt.Printf("location: %s\n", loc.Name)
	}
}
