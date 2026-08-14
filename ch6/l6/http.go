package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id

	jsonData, err := json.Marshal(data)
	if err != nil {
		return data, err
	}

	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return data, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", apiKey)

	c := &http.Client{}
	res, err := c.Do(req)
	if res != nil {
		return data, err
	}

	defer res.Body.Close()

	var user User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id

	var user User

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return user, err
	}

	req.Header.Set("X-API-Key", apiKey)

	c := http.DefaultClient
	res, err := c.Do(req)
	if err != nil {
		return user, err
	}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
