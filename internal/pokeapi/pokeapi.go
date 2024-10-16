package pokeapi

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

func GetMapLocations(url string) (LocationAreaResp, error) {
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

type PokemonInAreaResp struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetPokemonInArea(area string) (PokemonInAreaResp, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + area
	// Make the get request
	res, err := http.Get(url)
	if err != nil {
		return PokemonInAreaResp{}, fmt.Errorf("network error occured: %v", err)
	}
	// Check response status code
	if res.StatusCode != http.StatusOK {
		return PokemonInAreaResp{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	defer res.Body.Close()

	// Decode the response
	var pokemonArea PokemonInAreaResp
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&pokemonArea); err != nil {
		return PokemonInAreaResp{}, fmt.Errorf("failed to decode response: %v", err)
	}

	return pokemonArea, nil
}

func GetPokemon(pokemonString string) (Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonString

	// Make the get request
	res, err := http.Get(url)
	if err != nil {
		return Pokemon{}, fmt.Errorf("network error occured: %v", err)
	}
	// Check response status code
	if res.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	defer res.Body.Close()

	// Decode the response
	var pokemon Pokemon
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&pokemon); err != nil {
		return Pokemon{}, fmt.Errorf("failed to decode response: %v", err)
	}

	return pokemon, nil
}
