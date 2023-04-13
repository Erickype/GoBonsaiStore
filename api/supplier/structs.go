package supplier

type Supplier struct {
	Id        int    `json:"id"`
	CountryId int    `json:"countryId"`
	Name      string `json:"name"`
	Contact   string `json:"contact"`
	Address   string `json:"address"`
}

type Response struct {
	Status string `json:"status"`
}
