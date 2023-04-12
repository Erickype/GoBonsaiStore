package product

type Product struct {
	Id                int    `json:"id"`
	CategoryId        int    `json:"categoryId"`
	Name              string `json:"name"`
	ManufacturingDate string `json:"manufacturingDate"`
}

type Response struct {
	Status string `json:"status"`
}
