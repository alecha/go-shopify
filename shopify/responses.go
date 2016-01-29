package shopify

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

//CountResponse is a response to counts endpoint
type CountResponse struct {
	Count int `json:"count"`
}
