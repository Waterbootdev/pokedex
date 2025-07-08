package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (Resource, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Resource{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Resource{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Resource{}, err
	}

	locationsResp := Resource{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return Resource{}, err
	}

	return locationsResp, nil
}
