package shopify

//DiscountCodeResponse is a response to /discount_codes endpoint
type DiscountCodeResponse struct {
	DiscountCode DiscountCode `json:"discount_code"`
}

//OrdersResponse is a response to /orders endpoint
type OrdersResponse struct {
	Orders []Order `json:"orders"`
}

//OrderResponse is a response to /orders endpoint
type OrderResponse struct {
	Order Order `json:"order"`
}

//TransactionsResponse is a response to /orders/{id}/transactions
type TransactionsResponse struct {
	Transactions []Transaction `json:"transactions"`
}

type RefundsResponse struct {
	Refunds []Refund `json:"refunds"`
}

//CountResponse is a response to counts endpoint
type CountResponse struct {
	Count int `json:"count"`
}

//ProductsResponse is a response from products
type ProductsResponse struct {
	Products []Product `json:"products"`
}

//ProductResponse is a response for a product
type ProductResponse struct {
	Product Product `json:"product"`
}

//ImagesResponse is a response for product images
type ImagesResponse struct {
	Images []ProductImage `json:"images"`
}

//VariantsResponse is a response for product images
type VariantsResponse struct {
	Variants []Variant `json:"variants"`
}
