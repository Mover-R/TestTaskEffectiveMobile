package model

type PersonRequest struct {
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Patronymic *string `json:"patronymic,omitempty"`
}

type Person struct {
	UserID     int          `json:"user_id"`
	Name       string       `json:"name"`
	Surname    string       `json:"surname"`
	Patronymic *string      `json:"patronymic,omitempty"`
	Age        int          `json:"age"`
	Gender     string       `json:"gender"`
	Country    []CountryInf `json:"country"`
}

type AgeResponse struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type GenderResponse struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float32 `json:"probability"`
}

type CountryResponse struct {
	Count   int          `json:"count"`
	Name    string       `json:"name"`
	Country []CountryInf `json:"country"`
}

type CountryInf struct {
	CountryID   string  `json:"country_id"`
	Probability float32 `json:"probability"`
}
