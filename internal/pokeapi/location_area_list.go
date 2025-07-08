package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (*Resource, error) {

	resource := &Resource{}

	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return resource, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return resource, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return resource, err
	}

	err = json.Unmarshal(dat, resource)
	if err != nil {
		return resource, err
	}

	return resource, nil
}
