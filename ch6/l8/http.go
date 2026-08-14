package main

import (
	"net/http"
)

func deleteUser(baseURL, id, apiKey string) error {
	fullURL := baseURL + "/" + id

	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-API-Key", apiKey)

	c := http.DefaultClient
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode > 299 {
		return err
	}
	return nil
}
