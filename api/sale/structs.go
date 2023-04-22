package sale

type Detail struct {
	Id              int `json:"id"`
	ProductId       int `json:"productId"`
	SaleId          int `json:"saleId"`
	ProductQuantity int `json:"productQuantity"`
}
