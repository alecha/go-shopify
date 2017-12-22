package shopify

import (
	"fmt"
)

//CreateDiscountCode creates a discount code
func (shop *Shopify) CreateDiscountCode(priceRule string, discountCode DiscountCode) (*DiscountCode, []error) {
	var discountCodeResponse DiscountCodeResponse
	response, errors := shop.Post(fmt.Sprintf("price_rules/%v/discount_codes", priceRule), map[string]interface{}{
		"discount_code": discountCode,
	})
	if err := unmarshal(response, errors, &discountCodeResponse); len(err) > 0 {
		return nil, err
	}
	return &discountCodeResponse.DiscountCode, nil
}
