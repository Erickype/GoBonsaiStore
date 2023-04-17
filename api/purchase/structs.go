package purchase

type Purchase struct {
	PurchaseId int    `json:"purchaseId"`
	SupplierId int    `json:"supplierId"`
	Date       string `json:"date"`
	Cancel     bool   `json:"cancel"`
}
