package product

type ProductModel struct {
	ProductId    int64   `json:"productId"`
	Sku          string  `json:"sku"`
	Manufacturer string  `json:"manufacturer"`
	CategoryId   int64   `json:"categoryId"`
	Weight       float64 `json:"weight"`
	SomeOtherId  int64   `json:"someOtherId"`
}
