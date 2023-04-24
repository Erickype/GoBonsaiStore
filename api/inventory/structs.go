package inventory

type Product struct {
	Id    int     `json:"id"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}
