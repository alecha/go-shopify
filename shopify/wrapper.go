package shopify

import (
	"encoding/json"
	"fmt"
)

//GetOrders returns all the orders
func (shopify *Shopify) GetOrders() ([]Order, []error) {
	var orders OrdersResponse
	response, errors := shopify.Get("orders")
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

//GetProducts returns all the orders
func (shopify *Shopify) GetProducts() ([]Product, []error) {
	var products ProductsResponse
	response, errors := shopify.Get("products")
	if err := unmarshal(response, errors, &products); len(err) > 0 {
		return nil, err
	}
	return products.Products, nil
}

//GetProduct returns all the orders
func (shopify *Shopify) GetProduct(productID int64) (*Product, []error) {
	var product ProductResponse
	response, errors := shopify.Get(fmt.Sprintf("products/%v", productID))
	if err := unmarshal(response, errors, &product); len(err) > 0 {
		return nil, err
	}
	return &product.Product, nil
}

//GetProductImages returns all the orders
func (shopify *Shopify) GetProductImages(productID int64) ([]ProductImage, []error) {
	var images ImagesResponse
	response, errors := shopify.Get(fmt.Sprintf("products/%v/images", productID))
	if err := unmarshal(response, errors, &images); len(err) > 0 {
		return nil, err
	}
	return images.Images, nil
}

//GetProductVariants returns all the product variants
func (shopify *Shopify) GetProductVariants(productID int64) ([]Variant, []error) {
	var variants VariantsResponse
	response, errors := shopify.Get(fmt.Sprintf("products/%v/variants", productID))
	if err := unmarshal(response, errors, &variants); len(err) > 0 {
		return nil, err
	}
	return variants.Variants, nil
}

func unmarshal(responseData []byte, responseErrors []error, output interface{}) []error {
	if len(responseErrors) > 0 {
		return responseErrors
	}
	if err := json.Unmarshal(responseData, output); err != nil {
		return []error{err}
	}
	return nil
}
