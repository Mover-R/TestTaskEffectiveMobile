package model

type PersonRequest struct {
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Patronymic *string `json:"patronymic,omitempty"`
}

type Person struct {
	UserID     int        `json:"user_id"`
	Name       string     `json:"name"`
	Surname    string     `json:"surname"`
	Patronymic *string    `json:"patronymic,omitempty"`
	Age        int        `json:"age"`
	Gender     string     `json:"gender"`
	Country    CountryInf `json:"country"`
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

type Filter struct {
	Name       *string `json:"name,omitempty"`
	Surname    *string `json:"surname,omitempty"`
	Patronymic *string `json:"patronymic,omitempty"`
	AgeMin     *int    `json:"age_min,omitempty"`
	AgeMax     *int    `json:"age_max,omitempty"`
	Gender     *string `json:"gender,omitempty"`
	CountryID  *string `json:"country_id,omitempty"`
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
}
