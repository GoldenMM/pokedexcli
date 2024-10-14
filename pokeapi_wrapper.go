package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationAreaResp struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getMapLocations(url string) (LocationAreaResp, error) {
	// Make the get request
	res, err := http.Get(url)
	if err != nil {
		return LocationAreaResp{}, fmt.Errorf("network error occured: %v", err)
	}
	// Check response status code
	if res.StatusCode != http.StatusOK {
		return LocationAreaResp{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	defer res.Body.Close()

	// Decode the response
	var locationArea LocationAreaResp
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationArea); err != nil {
		return LocationAreaResp{}, fmt.Errorf("failed to decode response: %v", err)
	}

	return locationArea, nil
}
