package purchase

type Purchase struct {
	PurchaseId int    `json:"purchaseId"`
	SupplierId int    `json:"supplierId"`
	Date       string `json:"date"`
	Cancel     bool   `json:"cancel"`
}

type Detail struct {
	Id              int     `json:"id"`
	ProductId       int     `json:"productId"`
	PurchaseId      int     `json:"purchaseId"`
	SupplierId      int     `json:"supplierId"`
	ProductQuantity int     `json:"productQuantity"`
	ProductPrice    float64 `json:"productPrice"`
}
