package shopify

import (
	"encoding/json"
	"fmt"
	"log"
)

//GetOrders returns all the orders
func (shopify *Shopify) GetOrders() ([]Order, []error) {
	response, errors := shopify.Get("orders")
	if len(errors) > 0 {
		return nil, errors
	}

	var orders OrdersResponse
	if err := json.Unmarshal(response, &orders); err != nil {
		return nil, []error{err}
	}
	return orders.Orders, nil
}

//GetOrder returns a order given its id
func (shopify *Shopify) GetOrder(orderID int64) (*Order, []error) {
	response, errors := shopify.Get(fmt.Sprintf("orders/%v", orderID))
	if len(errors) > 0 {
		return nil, errors
	}

	var orderResponse OrderResponse
	if err := json.Unmarshal(response, &orderResponse); err != nil {
		return nil, []error{err}
	}
	return &orderResponse.Order, nil
}

//GetOrderTransactions returns the order's transactions
func (shopify *Shopify) GetOrderTransactions(orderID int64) ([]Transaction, []error) {
	response, errors := shopify.Get(fmt.Sprintf("orders/%v/transactions", orderID))
	if len(errors) > 0 {
		return nil, errors
	}

	log.Print(string(response))

	var transactionsResponse TransactionsResponse
	if err := json.Unmarshal(response, &transactionsResponse); err != nil {
		return nil, []error{err}
	}
	return transactionsResponse.Transactions, nil
}

//GetOrderTransactionsCount returns the order's transactions count
func (shopify *Shopify) GetOrderTransactionsCount(orderID int64) (int, []error) {
	response, errors := shopify.Get(fmt.Sprintf("orders/%v/transactions/count", orderID))
	if len(errors) > 0 {
		return 0, errors
	}

	log.Print(string(response))

	var count CountResponse
	if err := json.Unmarshal(response, &count); err != nil {
		return 0, []error{err}
	}
	return count.Count, nil
}

//GetOrdersCount returns all the products
func (shopify *Shopify) GetOrdersCount() (int, []error) {
	response, errors := shopify.Get("orders/count")
	if len(errors) > 0 {
		return 0, errors
	}

	var ordersCount CountResponse
	if err := json.Unmarshal(response, &ordersCount); err != nil {
		return 0, []error{err}
	}
	return ordersCount.Count, nil
}
