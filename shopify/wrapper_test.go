package shopify

var shop *Shopify

func init() {
	//init shop instance
	shop = New("", "", "")
}
