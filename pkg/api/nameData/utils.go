package namedata

import (
	"TestTaskEffectiveMobile/internal/names/model"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ageQuery     = `https://api.agify.io/?name=%s`
	genderQuery  = `https://api.genderize.io/?name=%s`
	countryQuery = `https://api.nationalize.io/?name=%s`
)

func GetAge(name string) (int, error) {
	query := fmt.Sprintf(ageQuery, name)
	resp, err := http.Get(query)
	if err != nil {
		return -1, fmt.Errorf("api.namedata.GetAge: %w", err)
	}
	defer resp.Body.Close()

	var data model.AgeResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return -1, fmt.Errorf("api.namedata.GetAge: %w", err)
	}

	return data.Age, nil
}

func GetGender(name string) (string, error) {
	query := fmt.Sprintf(genderQuery, name)
	resp, err := http.Get(query)
	if err != nil {
		return "", fmt.Errorf("api.namedata.GetGender: %w", err)
	}
	defer resp.Body.Close()

	var data model.GenderResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", fmt.Errorf("api.namedata.GetGender: %w", err)
	}

	return data.Gender, nil

}

func GetCountry(name string) ([]model.CountryInf, error) {
	query := fmt.Sprintf(countryQuery, name)
	resp, err := http.Get(query)
	if err != nil {
		return []model.CountryInf{}, fmt.Errorf("api.namedata.GetCountry: %w", err)
	}
	defer resp.Body.Close()

	var data model.CountryResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return []model.CountryInf{}, fmt.Errorf("api.namedata.GetCountry: %w", err)
	}

	return data.Country, nil
}
