package shopify

import "fmt"

///admin/orders/#{id}/refunds.json

//GetOrders returns all the orders
func (shopify *Shopify) GetOrderRefunds(orderID int64, parameters map[string]string) ([]Refund, []error) {
	var refunds RefundsResponse
	response, errors := shopify.Get(fmt.Sprintf("orders/%v/refunds", orderID))
	if err := unmarshal(response, errors, &refunds); len(err) > 0 {
		return nil, err
	}
	return refunds.Refunds, nil
}
