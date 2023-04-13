package client

type Client struct {
	Id        int    `json:"id"`
	CountryId int    `json:"countryId"`
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Birthdate string `json:"birthdate"`
	Email     string `json:"email"`
}
