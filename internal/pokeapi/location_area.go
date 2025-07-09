package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) CachedLocationArea(currentLocationAreasURL string) (*LocationArea, bool, error) {

	resource := &LocationArea{}

	dat, ok := c.cache.Get(currentLocationAreasURL)
	if ok {
		err := json.Unmarshal(dat, resource)
		if err == nil {
			return resource, ok, err
		} else {
			return resource, ok, errors.New("cached data is no resource")
		}
	} else {

		return resource, ok, nil
	}
}

func (c *Client) LocationArea(currentLocationAreasURL string) (*LocationArea, error) {

	resource, ok, err := c.CachedLocationArea(currentLocationAreasURL)

	if err != nil {
		return resource, err
	}

	if ok {
		return resource, nil
	}

	req, err := http.NewRequest("GET", currentLocationAreasURL, nil)

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

	c.cache.Add(currentLocationAreasURL, dat)

	err = json.Unmarshal(dat, resource)
	if err != nil {
		return resource, err
	}

	return resource, err
}
