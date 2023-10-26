package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type inputAge struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type inputGender struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

type inputNationality struct {
	Name    string `json:"name"`
	Country []struct {
		CountryID   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}

func (h *Handler) getAge(name string) (int, error) {
	var input inputAge
	resp, err := http.Get(fmt.Sprintf("https://api.agify.io/?name=%s", name))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&input); err != nil {
		return 0, err
	}

	return input.Age, nil
}

func (h *Handler) getGender(name string) (string, error) {
	var input inputGender
	resp, err := http.Get(fmt.Sprintf("https://api.genderize.io/?name=%s", name))

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&input); err != nil {
		return "", err
	}

	return input.Gender, nil
}

func (h *Handler) getNationality(name string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.nationalize.io/?name=%s", name))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var input inputNationality

	if err = json.NewDecoder(resp.Body).Decode(&input); err != nil {
		return "", err
	}

	return input.Country[0].CountryID, nil
}
