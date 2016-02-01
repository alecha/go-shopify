package shopify

import "fmt"

var emptyBody = make(map[string]string)

//GetOrders returns all the orders
func (shopify *Shopify) GetOrders(parameters map[string]string) ([]Order, []error) {
	var orders OrdersResponse
	response, errors := shopify.GetWithParameters("orders", parameters)
	if err := unmarshal(response, errors, &orders); len(err) > 0 {
		return nil, err
	}
	return orders.Orders, nil
}

//GetOrder returns a order given its id
func (shopify *Shopify) GetOrder(orderID int64) (*Order, []error) {
	var orderResponse OrderResponse
	response, errors := shopify.Get(fmt.Sprintf("orders/%v", orderID))
	if err := unmarshal(response, errors, &orderResponse); len(err) > 0 {
		return nil, err
	}
	return &orderResponse.Order, nil
}

//CloseOrder closes an order
func (shopify *Shopify) CloseOrder(orderID int64) (*Order, []error) {
	var orderResponse OrderResponse
	response, errors := shopify.Post(fmt.Sprintf("orders/%v/close", orderID), emptyBody)
	if err := unmarshal(response, errors, &orderResponse); len(err) > 0 {
		return nil, err
	}
	return &orderResponse.Order, nil
}

//OpenOrder re-opens an order
func (shopify *Shopify) OpenOrder(orderID int64) (*Order, []error) {
	var orderResponse OrderResponse
	response, errors := shopify.Post(fmt.Sprintf("orders/%v/order", orderID), emptyBody)
	if err := unmarshal(response, errors, &orderResponse); len(err) > 0 {
		return nil, err
	}
	return &orderResponse.Order, nil
}

//CancelOrder cancel an order
func (shopify *Shopify) CancelOrder(orderID int64) (*Order, []error) {
	var orderResponse OrderResponse
	response, errors := shopify.Post(fmt.Sprintf("orders/%v/cancel", orderID), emptyBody)
	if err := unmarshal(response, errors, &orderResponse); len(err) > 0 {
		return nil, err
	}
	return &orderResponse.Order, nil
}

//CreateOrder creates an order
func (shopify *Shopify) CreateOrder(order map[string]interface{}) (*Order, []error) {
	var orderResponse OrderResponse
	response, errors := shopify.Post("orders", order)
	if err := unmarshal(response, errors, &orderResponse); len(err) > 0 {
		return nil, err
	}
	return &orderResponse.Order, nil
}

//EditOrder edits an existing
func (shopify *Shopify) EditOrder(orderID int64, order map[string]interface{}) (*Order, []error) {
	var orderResponse OrderResponse
	order["id"] = orderID
	response, errors := shopify.Post(fmt.Sprintf("orders/%v", orderID), order)
	if err := unmarshal(response, errors, &orderResponse); len(err) > 0 {
		return nil, err
	}
	return &orderResponse.Order, nil
}

//DeleteOrder cancel an order
func (shopify *Shopify) DeleteOrder(orderID int64) (*Order, []error) {
	var orderResponse OrderResponse
	response, errors := shopify.Delete(fmt.Sprintf("orders/%v", orderID))
	if err := unmarshal(response, errors, &orderResponse); len(err) > 0 {
		return nil, err
	}
	return &orderResponse.Order, nil
}

//GetOrderTransactions returns the order's transactions
func (shopify *Shopify) GetOrderTransactions(orderID int64) ([]Transaction, []error) {
	var transactionsResponse TransactionsResponse
	response, errors := shopify.Get(fmt.Sprintf("orders/%v/transactions", orderID))
	if err := unmarshal(response, errors, &transactionsResponse); len(err) > 0 {
		return nil, err
	}
	return transactionsResponse.Transactions, nil
}

//GetOrderTransactionsCount returns the order's transactions count
func (shopify *Shopify) GetOrderTransactionsCount(orderID int64) (int, []error) {
	var count CountResponse
	response, errors := shopify.Get(fmt.Sprintf("orders/%v/transactions/count", orderID))
	if err := unmarshal(response, errors, &count); len(err) > 0 {
		return 0, err
	}
	return count.Count, nil
}

//GetOrdersCount returns all the products
func (shopify *Shopify) GetOrdersCount() (int, []error) {
	var ordersCount CountResponse
	response, errors := shopify.Get("orders/count")
	if err := unmarshal(response, errors, &ordersCount); len(err) > 0 {
		return 0, err
	}
	return ordersCount.Count, nil
}
