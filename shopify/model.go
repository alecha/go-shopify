package shopify

import "time"

//BillingAddress is a billing address
type BillingAddress struct {
	Address1     string `json:"address1"`
	Address2     string `json:"address2"`
	City         string `json:"city"`
	Company      string `json:"company"`
	Country      string `json:"country"`
	FirstName    string `json:"first_name"`
	ID           int64  `json:"id"`
	LastName     string `json:"last_name"`
	Phone        string `json:"phone"`
	Province     string `json:"province"`
	Zip          string `json:"zip"`
	Name         string `json:"name"`
	ProvinceCode string `json:"province_code"`
	CountryCode  string `json:"country_code"`
	Default      bool   `json:"default"`
}

//ClientDetails are details about a client
type ClientDetails struct {
	AcceptLanguage *string `json:"accept_language"` //TODO check
	BrowserHeight  int     `json:"browser_height"`
	BrowserIP      string  `json:"browser_ip"`
	BrowserWidth   int     `json:"browser_width"`
	SessionHash    *string `json:"session_hash"`
	UserAgent      *string `json:"user_agent"`
}

//Customer is a customer
type Customer struct {
	AcceptsMarketing bool      `json:"accepts_marketing"`
	CreatedAt        time.Time `json:"created_at"`
	Email            string    `json:"email"`
	ID               int64     `json:"id"`
	FirstName        string    `json:"first_name"`
	Note             string    `json:"note"`
	LastName         string    `json:"last_name"`
	OrdersCount      int       `json:"orders_count"`
	State            string    `json:"state"`
	TotalSpent       string    `json:"total_spent"`
	UpdatedAt        time.Time
	Tags             string `json:"tags"`
}

//Discount is a discount
type Discount struct {
	ID                 int64     `json:"id"`
	DiscountType       string    `json:"discount_type"`
	Code               string    `json:"code"`
	Value              string    `json:"value"`
	EndsAt             time.Time `json:"ends_at"`
	StartsAt           time.Time `json:"starts_at"`
	Status             string    `json:"status"`
	MinimumOrderAmount string    `json:"minimum_order_amount"`
	UsageLimit         int       `json:"usage_limit"`
	AppliesToID        int64     `json:"applies_to_id"`
	AppliesOnce        bool      `json:"applies_once"`
	AppliesToResource  string    `json:"applies_to_resource"`
	TimesUsed          int       `json:"times_used"`
}

//DiscountCode is a discount code
type DiscountCode struct {
	Amount string `json:"amount"`
	Code   string `json:"code"`
	Type   string `json:"type"`
}

//Fulfillment is a fulfillment
type Fulfillment struct {
	ID              int64     `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	OrderID         int64     `json:"order_id"`
	Status          string    `json:"status"`
	TrackingCompany string    `json:"tracking_company"`
	TrackingNumber  string    `json:"tracking_number"`
	UpdatedAt       time.Time `json:"updated_at"`
}

//LineItem is an order line item
type LineItem struct {
	FulfillableQuantity int       `json:"fulfillable_quantity"`
	FulfillmentService  *string   `json:"fulfillment_service"`
	FulfillmentStatus   *string   `json:"fulfillment_status"`
	Grams               int       `json:"grams"`
	ID                  int64     `json:"id"`
	Price               string    `json:"price"` //e.g. 199.99
	ProductID           int64     `json:"product_id"`
	Quantity            int       `json:"id"`
	RequiresShipping    bool      `json:"requires_shipping"`
	SKU                 string    `json:"sku"`
	Title               string    `json:"title"`
	VariantID           int64     `json:"variant_id"`
	VariantTitle        string    `json:"variant_title"`
	Vendor              string    `json:"vendor"`
	GiftCard            *bool     `json:"gift_card"`
	Taxable             bool      `json:"taxable"`
	TaxLines            []TaxLine `json:"tax_line"`
	TotalDiscount       string    `json:"total_discount"`
}

//NoteAttribute is a note attribute
type NoteAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

//Order is a product
type Order struct {
	BillingAddress        BillingAddress  `json:"billing_address"`
	BrowserIP             string          `json:"browser_ip"`
	BuyerAcceptsMarketing bool            `json:"buyer_accepts_marketing"`
	CancelReason          *string         `json:"cancel_reason"`
	CancelledAt           *time.Time      `json:"cancelled_at"`
	ClientDetails         ClientDetails   `json:"client_details"`
	ClosedAt              *time.Time      `json:"closed_at"`
	CreatedAt             time.Time       `json:"created_at"`
	Currency              string          `json:"currency"`
	Customer              *Customer       `json:"customer"`
	DiscountCodes         []DiscountCode  `json:"discount_codes"`
	Email                 string          `json:"email"`
	FinancialStatus       string          `json:"financial_status"`
	Fulfillments          []Fulfillment   `json:"fulfillments"`
	FulfillmentStatus     string          `json:"fulfillment_status"`
	Tags                  string          `json:"tags"`
	ID                    int64           `json:"id"`
	LandingSite           string          `json:"landing_site"`
	LineItems             []LineItem      `json:"line_items"`
	Name                  string          `json:"name"`
	Note                  *string         `json:"note"`
	NoteAttributes        []NoteAttribute `json:"note_attributes"`
	Number                int64           `json:"number"`
	OrderNumber           int64           `json:"order_number"`
	PaymentGatewayNames   []string        `json:"payment_gateway_names"`
	ProcessedAt           time.Time       `json:"processed_at"`
	ProcessingMethod      string          `json:"processing_method"`
	ReferringSite         string          `json:"referring_site"`
	Refunds               []Refund        `json:"refunds"`
	ShippingAddress       ShippingAddress `json:"shipping_address"`
	ShippingLines         []ShippingLine  `json:"shipping_lines"`
	SourceName            string          `json:"source_name"`
	SubtotalPrice         string          `json:"subtotal_price"`
	TaxLines              []TaxLine       `json:"tax_lines"`
	TaxesIncluded         bool            `json:"taxes_included"`
	TotalDiscounts        string          `json:"total_discounts"`
	TotalPrice            string          `json:"total_price"`
	TotalTax              string          `json:"total_tax"`
	TotalWeight           float64         `json:"total_weight"`
	UpdatedAt             time.Time       `json:"updatedAt"`
}

//PaymentDetails are the details about a payment
type PaymentDetails struct {
	AvsResultCode     *string `json:"avs_result_code"`
	CreditCardBin     *string `json:"credit_card_bin"`
	CvvResultCode     *string `json:"cvv_result_code"`
	CreditCardNumber  string  `json:"credit_card_number"`
	CreditCardCompany string  `json:"credit_card_company"`
}

//Product is a product
type Product struct {
	BodyHTML                       string                   `json:"body_html"`
	CreatedAt                      time.Time                `json:"created_at"`
	Handle                         string                   `json:"handle"`
	ID                             int64                    `json:"id"`
	Images                         []ProductImage           `json:"images"`
	Options                        []map[string]interface{} `json:"options"`
	ProductType                    string                   `json:"product_type"`
	PublishedAt                    *time.Time               `json:"published_at"`
	PublishedScope                 string                   `json:"published_scope"`
	Tags                           string                   `json:"tags"`
	TemplateSuffix                 string                   `json:"template_suffix"`
	Title                          string                   `json:"title"`
	MetafieldsGlobalTitleTag       string                   `json:"metafields_global_title_tag"`
	MetafieldsGlobalDescriptionTag string                   `json:"metafields_global_description_tag"`
	UpdatedAt                      time.Time                `json:"updatedAt"`
	Variants                       []Variant                `json:"variants"`
	Vendor                         string                   `json:"vendor"`
}

//ProductImage is a product's image
type ProductImage struct {
	CreatedAt  time.Time `json:"created_at"`
	ID         int64     `json:"id"`
	Position   int       `json:"position"`
	ProductID  int64     `json:"product_id"`
	VariantIDs []int64   `json:"variant_ids"`
	Src        string    `json:"src"`
	UpdatedAt  time.Time `json:"id"`
}

//Refund is a refund
type Refund struct {
	CreatedAt       time.Time        `json:"created_at"`
	ID              int64            `json:"id"`
	Note            string           `json:"note"`
	RefundLineItems []RefundLineItem `json:"refund_line_items"`
	Restock         bool             `json:"restock"`
	Transactions    []Transaction    `json:"transactions"`
	UserID          int64            `json:"user_id"`
}

//RefundLineItem is a refund line item
type RefundLineItem struct {
	ID         int64    `json:"id"`
	LineItem   LineItem `json:"line_item"`
	LineItemID int64    `json:"line_item_id"`
	Quantity   int      `json:"quantity"`
}

//ShippingAddress is a billing address
type ShippingAddress struct {
	Address1     string  `json:"address1"`
	Address2     string  `json:"address2"`
	City         string  `json:"city"`
	Company      string  `json:"company"`
	Country      string  `json:"country"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Phone        string  `json:"phone"`
	Province     string  `json:"province"`
	Zip          string  `json:"zip"`
	Name         string  `json:"name"`
	ProvinceCode string  `json:"province_code"`
	CountryCode  string  `json:"country_code"`
}

//ShippingLine is a shipping line
type ShippingLine struct {
	Code     string    `json:"code"`
	Price    string    `json:"price"`
	Source   string    `json:"source"`
	Title    string    `json:"title"`
	TaxLines []TaxLine `json:"tax_lines"`
}

//TaxLine is a tax line
type TaxLine struct {
	Title string  `json:"title"`
	Price float64 `json:"price"`
	Rate  float64 `json:"rate"`
}

//Transaction is a transaction
type Transaction struct {
	Amount         string         `json:"amount"`
	Authorization  *string        `json:"authorization"`
	CreatedAt      time.Time      `json:"created_at"`
	DeviceID       *string        `json:"device_id"`
	Gateway        string         `json:"gateway"`
	SourceName     string         `json:"source_name"`
	PaymentDetails PaymentDetails `json:"payment_details"`
	ID             int64          `json:"id"`
	Kind           string         `json:"kind"`
	OrderID        int64          `json:"orderId"`
	Receipt        string         `json:"receipt"`
	ErrorCode      string         `json:"error_code"`
	Status         string         `json:"status"`
	Test           bool           `json:"test"`
	UserID         *int64         `json:"userId"`
	Currency       string         `json:"currency"`
}

//Variant is a product's variant
type Variant struct {
	BarCode             string    `json:"bar_code"`
	CompareAtPrice      string    `json:"compare_at_price"`
	CreatedAt           time.Time `json:"created_at"`
	FulfillmentService  string    `json:"fulfillment_service"`
	Grams               float64   `json:"grams"`
	Weight              float64   `json:"weight"`
	WeightUnit          string    `json:"weight_unit"`
	ID                  int64     `json:"id"`
	InventoryManagement string    `json:"inventory_management"`
	InventoryPolicy     string    `json:"inventory_policy"`
	InventoryQuantity   int       `json:"inventory_quantity"`
	Option1             string    `json:"option1"`
	Option2             string    `json:"option2"`
	Option3             string    `json:"option3"`
	Position            int       `json:"position"`
	Price               string    `json:"price"`
	ProductID           int64     `json:"product_id"`
	RequiresShipping    bool      `json:"requires_shipping"`
	SKU                 string    `json:"sku"`
	Taxable             bool      `json:"taxable"`
	Title               string    `json:"title"`
	UpdatedAt           time.Time `json:"updated_at"`
}
