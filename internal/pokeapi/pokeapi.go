package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiurl = "https://pokeapi.co/api/v2/"

func PokeDo(endpoint string, obj interface{}) error {

	req, err := http.NewRequest(http.MethodGet, apiurl+endpoint, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &obj)
}

func PokeResource(endpoint string, offset int, limit int) (result Resource, err error) {

	err = PokeDo(fmt.Sprintf("%s?offset=%d&limit=%d", endpoint, offset, limit),
		&result)
	return
}

func PokeLocationAreaResource() (Resource, error) {
	return PokeResource("location-area", 0, 9999)
}
