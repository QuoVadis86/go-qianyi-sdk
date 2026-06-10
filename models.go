package qianyi

// Shop represents a store/platform account connected to QERP.
type Shop struct {
	ShopID            int64  `json:"shopId"`
	Name              string `json:"name"`
	Platform          string `json:"platform"`
	Status            string `json:"status"`
	SiteCode          string `json:"siteCode,omitempty"`
	AuthExpiredStatus string `json:"authExpiredStatus"`
	CreateTime        int64  `json:"createTime"`
	OnlineShopID      string `json:"onlineShopId,omitempty"`
	Currency          string `json:"currency,omitempty"`
	TimeZoneID        string `json:"timeZoneId,omitempty"`
}

// Warehouse represents a physical or virtual warehouse in QERP.
type Warehouse struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Kind         string `json:"kind"`
	ProviderName string `json:"providerName"`
	Code         string `json:"code,omitempty"`
	CodeName     string `json:"codeName,omitempty"`
	Country      string `json:"country"`
	TimezoneID   string `json:"timezoneId"`
	Status       string `json:"status,omitempty"`
}

// Sku represents a stock keeping unit (product) in QERP.
type Sku struct {
	Sku            string  `json:"sku"`
	Title          string  `json:"title"`
	Barcode        string  `json:"barcode,omitempty"`
	Type           string  `json:"type"`
	IsAssembly     int     `json:"isAssembly,omitempty"`
	PicURL         string  `json:"picUrl,omitempty"`
	SaleStatus     string  `json:"saleStatus,omitempty"`
	Weight         float64 `json:"weight"`
	NetWeight      float64 `json:"netWeight"`
	WeightUnit     string  `json:"weightUnit"`
	Length         float64 `json:"length"`
	Width          float64 `json:"width"`
	Height         float64 `json:"height"`
	DimensionUnit  string  `json:"dimensionUnit"`
	Enable         int     `json:"enable"`
	Price          float64 `json:"price,omitempty"`
	Brand          string  `json:"brand,omitempty"`
	Unit           string  `json:"unit,omitempty"`
	Color          string  `json:"color,omitempty"`
	Size           string  `json:"size,omitempty"`
	Description    string  `json:"description,omitempty"`
	CategoryName1  string  `json:"categoryName1,omitempty"`
	CategoryName2  string  `json:"categoryName2,omitempty"`
	CategoryName3  string  `json:"categoryName3,omitempty"`
	PurchaseCost   float64 `json:"purchaseCost,omitempty"`
	RemarkName     string  `json:"remarkName,omitempty"`
}

// Buyer holds the receiver/shipping address for an order.
type Buyer struct {
	BuyerID      string `json:"buyerId,omitempty"`
	ReceiverName string `json:"receiverName"`
	Phone        string `json:"phone,omitempty"`
	Email        string `json:"email,omitempty"`
	Country      string `json:"country"`
	Province     string `json:"province"`
	City         string `json:"city"`
	District     string `json:"district,omitempty"`
	PostCode     string `json:"postCode"`
	Address1     string `json:"address1"`
	Address2     string `json:"address2,omitempty"`
}

// OrderSku represents a product line item within a sales order.
type OrderSku struct {
	Sku               string  `json:"sku"`
	PayAmount         float64 `json:"payAmount,omitempty"`
	PaymentPrice      float64 `json:"paymentPrice,omitempty"`
	Quantity          int     `json:"quantity"`
	ShippingPrice     float64 `json:"shippingPrice,omitempty"`
	PromotionDiscount float64 `json:"promotionDiscount,omitempty"`
}

// Order represents a sales order in QERP.
type Order struct {
	OrderNumber       string     `json:"orderNumber"`
	OnlineOrderNumber string     `json:"onlineOrderNumber,omitempty"`
	ParentOrderNumber string     `json:"parentOrderNumber,omitempty"`
	Shop              string     `json:"shop"`
	Warehouse         string     `json:"warehouse,omitempty"`
	Status            string     `json:"status"`
	WMSStatus         string     `json:"wmsStatus,omitempty"`
	Currency          string     `json:"currency"`
	TotalAmount       float64    `json:"totalAmount"`
	Freight           float64    `json:"freight,omitempty"`
	Platform          string     `json:"platform"`
	Carrier           string     `json:"carrier,omitempty"`
	TrackingNumber    string     `json:"trackingNumber,omitempty"`
	PayTime           int64      `json:"payTime,omitempty"`
	ShippingTime      int64      `json:"shippingTime,omitempty"`
	CreateTime        int64      `json:"createTime"`
	UpdateTime        int64      `json:"updateTime"`
	Buyer             *Buyer     `json:"buyer,omitempty"`
	SkuList           []OrderSku `json:"skuList,omitempty"`
}

// ReturnOrder represents a refund/return order in QERP.
type ReturnOrder struct {
	ReturnNumber  string      `json:"returnNumber"`
	OrderNumber   string      `json:"orderNumber,omitempty"`
	Warehouse     string      `json:"warehouse"`
	Status        string      `json:"status"`
	Shop          string      `json:"shop"`
	CreateTime    int64       `json:"createTime"`
	UpdateTime    int64       `json:"updateTime"`
	Reason        string      `json:"reason,omitempty"`
	Carrier       string      `json:"carrier,omitempty"`
	CustomNumber  string      `json:"customNumber,omitempty"`
	Type          string      `json:"type"`
	ReturnSkuList []ReturnSku `json:"returnSkuList,omitempty"`
}

// ReturnSku represents a product line item within a return order.
type ReturnSku struct {
	Sku      string `json:"sku"`
	Quantity int    `json:"quantity"`
	Remark   string `json:"remark,omitempty"`
}

// SkuInventory represents the current inventory state for a SKU in a warehouse.
type SkuInventory struct {
	Sku              string `json:"sku"`
	SkuName          string `json:"skuName,omitempty"`
	Warehouse        string `json:"warehouse"`
	WarehouseCode    string `json:"warehouseCode,omitempty"`
	Total            int    `json:"total"`
	Available        int    `json:"available"`
	Allocated        int    `json:"allocated"`
	Unavailable      int    `json:"unavailable,omitempty"`
	ShippingQuantity int    `json:"shippingQuantity,omitempty"`
}
