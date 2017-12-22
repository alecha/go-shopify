package shopify

import "fmt"

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
