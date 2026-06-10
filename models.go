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
	ShopGroupVOList   []ShopGroup `json:"shopGroupVOList,omitempty"`
}

// ShopGroup represents a shop group.
type ShopGroup struct {
	ID            int64  `json:"id"`
	ShopGroupName string `json:"shopGroupName"`
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
	Sku                        string            `json:"sku"`
	Title                      string            `json:"title"`
	Barcode                    string            `json:"barcode,omitempty"`
	Type                       string            `json:"type"`
	IsAssembly                 int               `json:"isAssembly,omitempty"`
	PicURL                     string            `json:"picUrl,omitempty"`
	PicURLList                 []string          `json:"picUrlList,omitempty"`
	SaleStatus                 string            `json:"saleStatus,omitempty"`
	Weight                     float64           `json:"weight"`
	NetWeight                  float64           `json:"netWeight"`
	WeightUnit                 string            `json:"weightUnit"`
	Length                     float64           `json:"length"`
	Width                      float64           `json:"width"`
	Height                     float64           `json:"height"`
	DimensionUnit              string            `json:"dimensionUnit"`
	Enable                     int               `json:"enable"`
	Price                      float64           `json:"price,omitempty"`
	PriceUnit                  string            `json:"priceUnit,omitempty"`
	Brand                      string            `json:"brand,omitempty"`
	Unit                       string            `json:"unit,omitempty"`
	Color                      string            `json:"color,omitempty"`
	Size                       string            `json:"size,omitempty"`
	Description                string            `json:"description,omitempty"`
	DescriptionEn              string            `json:"descriptionEn,omitempty"`
	FunctionDescription        string            `json:"functionDescription,omitempty"`
	AbbrTitle                  string            `json:"abbrTitle,omitempty"`
	ItemPackage                int               `json:"itemPackage,omitempty"`
	PackingRate                int               `json:"packingRate,omitempty"`
	SingleBoxCode              string            `json:"singleBoxCode,omitempty"`
	SingleItemVolume           float64           `json:"singleItemVolume,omitempty"`
	CartonVolumeUnit           string            `json:"cartonVolumeUnit,omitempty"`
	CategoryName1              string            `json:"categoryName1,omitempty"`
	CategoryName2              string            `json:"categoryName2,omitempty"`
	CategoryName3              string            `json:"categoryName3,omitempty"`
	CustomProp1                string            `json:"customProp1,omitempty"`
	CustomProp2                string            `json:"customProp2,omitempty"`
	CustomProp3                string            `json:"customProp3,omitempty"`
	PurchaseCost               float64           `json:"purchaseCost,omitempty"`
	PurchaseCostUnit           string            `json:"purchaseCostUnit,omitempty"`
	RemarkName                 string            `json:"remarkName,omitempty"`
	SingleSkuList              []SubSkuDTO       `json:"singleSkuList,omitempty"`
	Parts                      []SubSkuDTO       `json:"parts,omitempty"`
	SkuCustomFieldValueVOList  []CustomField     `json:"skuCustomFieldValueVOList,omitempty"`
	SkuSupplierList            []SkuSupplierDTO  `json:"skuSupplierList,omitempty"`
	IsRelatedSingleCost        bool              `json:"isRelatedSingleCost,omitempty"`
	WarehouseNameList          []string          `json:"warehouseNameList,omitempty"`
	CreateAllWarehouseItems    bool              `json:"createAllWarehouseItems,omitempty"`
	DeveloperUserNameList      []string          `json:"developerUserNameList,omitempty"`
	// Customs/declaration fields
	ChineseCustomsDeclarationName  string  `json:"chineseCustomsDeclarationName,omitempty"`
	EnglishCustomsDeclarationName  string  `json:"englishCustomsDeclarationName,omitempty"`
	CdPriceMethod                  string  `json:"cdPriceMethod,omitempty"`
	CdPriceRate                    float64 `json:"cdPriceRate,omitempty"`
	CdPriceMaximum                 float64 `json:"cdPriceMaximum,omitempty"`
	CustomsDeclarationPrice        float64 `json:"customsDeclarationPrice,omitempty"`
	CustomsDeclarationPriceUnit    string  `json:"customsDeclarationPriceUnit,omitempty"`
	CustomsCode                    string  `json:"customsCode,omitempty"`
	ExportTax                      float64 `json:"exportTax,omitempty"`
	Declaration                    string  `json:"declaration,omitempty"`
	DangerousTransportGoodsType    string  `json:"dangerousTransportGoodsType,omitempty"`
	BatteryType                    int     `json:"batteryType,omitempty"`
	CcPriceMethod                  string  `json:"ccPriceMethod,omitempty"`
	CcPriceRate                    float64 `json:"ccPriceRate,omitempty"`
	CcPriceMaximum                 float64 `json:"ccPriceMaximum,omitempty"`
	CcPrice                        float64 `json:"ccPrice,omitempty"`
	CcPriceUnit                    string  `json:"ccPriceUnit,omitempty"`
	NeedQualityInspection          bool    `json:"needQualityInspection,omitempty"`
	Asin                           string  `json:"asin,omitempty"`
	// Carton/box fields
	CartonLength             float64 `json:"cartonLength,omitempty"`
	CartonWidth              float64 `json:"cartonWidth,omitempty"`
	CartonHeight             float64 `json:"cartonHeight,omitempty"`
	CartonWeight             float64 `json:"cartonWeight,omitempty"`
	CartonNetWeight          float64 `json:"cartonNetWeight,omitempty"`
	CartonDimensionUnit      string  `json:"cartonDimensionUnit,omitempty"`
	CartonWeightUnit         string  `json:"cartonWeightUnit,omitempty"`
}

// SubSkuDTO represents a sub-SKU within a combination/assembly product.
type SubSkuDTO struct {
	Sku      string `json:"sku"`
	Quantity int    `json:"quantity"`
}

// SkuSupplierDTO represents supplier information for a SKU.
type SkuSupplierDTO struct {
	Name                  string   `json:"name"`
	IsDefault             bool     `json:"isDefault"`
	PurchaserList         []string `json:"purchaserList"`
	DeliveryCycle         int      `json:"deliveryCycle"`
	PurchasePrice         float64  `json:"purchasePrice,omitempty"`
	PurchasePriceUnit     string   `json:"purchasePriceUnit,omitempty"`
	MinimumPurchaseQuantity int    `json:"minimumPurchaseQuantity,omitempty"`
	PurchaseTaxRate       float64  `json:"purchaseTaxRate,omitempty"`
	DrawbackRate          float64  `json:"drawbackRate,omitempty"`
	PurchaseURL           string   `json:"purchaseUrl,omitempty"`
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
	Sku               string       `json:"sku"`
	PayAmount         float64      `json:"payAmount,omitempty"`
	PaymentPrice      float64      `json:"paymentPrice,omitempty"`
	Quantity          int          `json:"quantity"`
	ShippingPrice     float64      `json:"shippingPrice,omitempty"`
	PromotionDiscount float64      `json:"promotionDiscount,omitempty"`
	BatchNo           string       `json:"batchNo,omitempty"`
	MfgDate           string       `json:"mfgDate,omitempty"`
	ExpDate           string       `json:"expDate,omitempty"`
	OriginCountry     string       `json:"originCountry,omitempty"`
	// Response-only fields
	OrderSkuID        int64        `json:"orderSkuId,omitempty"`
	OnlineItemID      string       `json:"onlineItemId,omitempty"`
	OnlineProductCode string       `json:"onlineProductCode,omitempty"`
	OnlineProductPicURL string     `json:"onlineProductPicUrl,omitempty"`
	OnlineProductTitle string      `json:"onlineProductTitle,omitempty"`
	OnlineTransactionID string     `json:"onlineTransactionId,omitempty"`
	OriginalPrice     float64      `json:"originalPrice,omitempty"`
	PlatformDiscount  float64      `json:"platformDiscount,omitempty"`
	DiscountPrice     float64      `json:"discountPrice,omitempty"`
	SubSkuList        []SubSkuDTO  `json:"subSkuList,omitempty"`
	Tag               *OrderSkuTag `json:"tag,omitempty"`
}

// OrderSkuTag represents flags on an order SKU line.
type OrderSkuTag struct {
	AllReturned  int `json:"allReturned,omitempty"`
	HasRefund    int `json:"hasRefund,omitempty"`
	IsGift       int `json:"isGift,omitempty"`
	OnlineShipped int `json:"onlineShipped,omitempty"`
	PreSale      int `json:"preSale,omitempty"`
}

// OrderTag represents flags/tags on a sales order.
type OrderTag struct {
	HasRefund                int `json:"hasRefund,omitempty"`
	ItemReturned             int `json:"itemReturned,omitempty"`
	Consolidated             int `json:"consolidated,omitempty"`
	Split                    int `json:"split,omitempty"`
	Locked                   int `json:"locked,omitempty"`
	SendWms                  int `json:"sendWms,omitempty"`
	SendFailed               int `json:"sendFailed,omitempty"`
	OnlineShipFeedbackAlready int `json:"onlineShipFeedbackAlready,omitempty"`
	OutOfStock               int `json:"outOfStock,omitempty"`
	PreSale                  int `json:"preSale,omitempty"`
	OnlineShipped            int `json:"onlineShipped,omitempty"`
	PlatformFulfillment      int `json:"platformFulfillment,omitempty"`
	PartRefund               int `json:"partRefund,omitempty"`
	AllRefund                int `json:"allRefund,omitempty"`
	PartReturned             int `json:"partReturned,omitempty"`
	AllReturned              int `json:"allReturned,omitempty"`
	ReShip                   int `json:"reShip,omitempty"`
}

// Order represents a sales order in QERP.
type Order struct {
	OrderNumber              string            `json:"orderNumber"`
	OnlineOrderNumber        string            `json:"onlineOrderNumber,omitempty"`
	ParentOrderNumber        string            `json:"parentOrderNumber,omitempty"`
	SubOrderNumberList       []string          `json:"subOrderNumberList,omitempty"`
	IsOriginalOrder          bool              `json:"isOriginalOrder,omitempty"`
	Shop                     string            `json:"shop"`
	ShopID                   int64             `json:"shopId,omitempty"`
	Warehouse                string            `json:"warehouse,omitempty"`
	Status                   string            `json:"status"`
	WMSStatus                string            `json:"wmsStatus,omitempty"`
	Currency                 string            `json:"currency"`
	TotalAmount              float64           `json:"totalAmount"`
	Freight                  float64           `json:"freight,omitempty"`
	Platform                 string            `json:"platform"`
	Carrier                  string            `json:"carrier,omitempty"`
	TrackingNumber           string            `json:"trackingNumber,omitempty"`
	PayTime                  int64             `json:"payTime,omitempty"`
	ShippingTime             int64             `json:"shippingTime,omitempty"`
	CreateTime               int64             `json:"createTime"`
	UpdateTime               int64             `json:"updateTime"`
	AuditTime                int64             `json:"auditTime,omitempty"`
	LatestShipDate           int64             `json:"latestShipDate,omitempty"`
	PlatformShippingTime     int64             `json:"platformShippingTime,omitempty"`
	Buyer                    *Buyer            `json:"buyer,omitempty"`
	SkuList                  []OrderSku        `json:"skuList,omitempty"`
	Tag                      *OrderTag         `json:"tag,omitempty"`
	BuyerMessage             string            `json:"buyerMessage,omitempty"`
	SellerRemarks            string            `json:"sellerRemarks,omitempty"`
	PaymentMethod            string            `json:"paymentMethod,omitempty"`
	LogisticsSelected        string            `json:"logisticsSelected,omitempty"`
	OnlineStatus             string            `json:"onlineStatus,omitempty"`
	SiteCode                 string            `json:"siteCode,omitempty"`
	IsDeleted                int               `json:"isDeleted,omitempty"`
	SalesRecordNumber        string            `json:"salesRecordNumber,omitempty"`
	IsAFN                    int               `json:"isAfn,omitempty"`
	IsBusinessOrder          bool              `json:"isBusinessOrder,omitempty"`
	EstimateFulfillmentFee   float64           `json:"estimateFulfillmentFee,omitempty"`
	TotalDiscount            float64           `json:"totalDiscount,omitempty"`
	SellerDiscount           float64           `json:"sellerDiscount,omitempty"`
	PlatformRebate           float64           `json:"platformRebate,omitempty"`
	BuyerPaidShippingFee     float64           `json:"buyerPaidShippingFee,omitempty"`
	FinalProductProtection   float64           `json:"finalProductProtection,omitempty"`
	SellerDiscountForWook    float64           `json:"sellerDiscountForWook,omitempty"`
	PlatformRebateForWook    float64           `json:"platformRebateForWook,omitempty"`
	PlatformReturnToSeller   float64           `json:"platformReturnToSeller,omitempty"`
	OrderCustomFieldValueVOList []CustomField  `json:"orderCustomFieldValueVOList,omitempty"`
}

// ReturnOrder represents a refund/return order in QERP.
type ReturnOrder struct {
	ReturnNumber   string      `json:"returnNumber"`
	OrderNumber    string      `json:"orderNumber,omitempty"`
	OnlineOrderNumber string   `json:"onlineOrderNumber,omitempty"`
	Warehouse      string      `json:"warehouse"`
	Status         string      `json:"status"`
	Shop           string      `json:"shop"`
	CreateTime     int64       `json:"createTime"`
	UpdateTime     int64       `json:"updateTime"`
	FinishTime     int64       `json:"finishTime,omitempty"`
	CommitTime     int64       `json:"commitTime,omitempty"`
	ReceiveTime    int64       `json:"receiveTime,omitempty"`
	Reason         string      `json:"reason,omitempty"`
	Remark         string      `json:"remark,omitempty"`
	Carrier        string      `json:"carrier,omitempty"`
	CustomNumber   string      `json:"customNumber,omitempty"`
	Type           string      `json:"type"`
	BuyerID        string      `json:"buyerId,omitempty"`
	Currency       string      `json:"currency,omitempty"`
	OrderRefundAmount float64  `json:"orderRefundAmount,omitempty"`
	TotalAmount    float64     `json:"totalAmount,omitempty"`
	AsnNumber      string      `json:"asnNumber,omitempty"`
	AsnID          string      `json:"asnId,omitempty"`
	OrderStatus    string      `json:"orderStatus,omitempty"`
	ReturnSkuList  []ReturnSku `json:"returnSkuList,omitempty"`
	RefundCustomFieldValueList []CustomField `json:"refundCustomFieldValueVOList,omitempty"`
}

// ReturnSku represents a product line item within a return order.
type ReturnSku struct {
	Sku              string           `json:"sku"`
	OrderSkuID       int64            `json:"orderSkuId,omitempty"`
	Quantity         int              `json:"quantity"`
	Remark           string           `json:"remark,omitempty"`
	Selected         int              `json:"selected,omitempty"`
	StorageLocationCode string        `json:"storageLocationCode,omitempty"`
	GoodQuantity     int              `json:"goodQuantity,omitempty"`
	BadQuantity      int              `json:"badQuantity,omitempty"`
	ItemRefundAmount float64          `json:"itemRefundAmount,omitempty"`
	RefundAmount     float64          `json:"refundAmount,omitempty"`
	LocAndQuantityList []LocAndQuantity `json:"locAndQuantityList,omitempty"`
}

// LocAndQuantity represents storage location code and receive quantity.
type LocAndQuantity struct {
	StorageLocationCode string `json:"storageLocationCode"`
	ReceiveQuantity     int    `json:"receiveQuantity"`
}

// SkuInventory represents the current inventory state for a SKU in a warehouse.
type SkuInventory struct {
	Sku                     string  `json:"sku"`
	SkuName                 string  `json:"skuName,omitempty"`
	Warehouse               string  `json:"warehouse"`
	WarehouseCode           string  `json:"warehouseCode,omitempty"`
	Total                   int     `json:"total"`
	Available               int     `json:"available"`
	Allocated               int     `json:"allocated"`
	Unavailable             int     `json:"unavailable,omitempty"`
	ShippingQuantity        int     `json:"shippingQuantity,omitempty"`
	PurchaseShippingQuantity int    `json:"purchaseShippingQuantity,omitempty"`
	FirstLegShippingQuantity int    `json:"firstLegShippingQuantity,omitempty"`
	TransferShippingQuantity int    `json:"transferShippingQuantity,omitempty"`
	AssemblyShippingQuantity int    `json:"assemblyShippingQuantity,omitempty"`
	ReturnShippingQuantity  int     `json:"returnShippingQuantity,omitempty"`
	ManualShippingQuantity  int     `json:"manualShippingQuantity,omitempty"`
	OrderAllocated          int     `json:"orderAllocated,omitempty"`
	FirstLegAllocated       int     `json:"firstLegAllocated,omitempty"`
	TransferAllocated       int     `json:"transferAllocated,omitempty"`
	AssemblyAllocated       int     `json:"assemblyAllocated,omitempty"`
	TotalCost               float64 `json:"totalCost,omitempty"`
	AvailableCost           float64 `json:"availableCost,omitempty"`
	UnavailableCost         float64 `json:"unavailableCost,omitempty"`
	AllocatedCost           float64 `json:"allocatedCost,omitempty"`
	ShippingCost            float64 `json:"shippingCost,omitempty"`
	TotalGoods              float64 `json:"totalGoods,omitempty"`
	UnavailableGoods        float64 `json:"unavailableGoods,omitempty"`
	AvailableGoods          float64 `json:"availableGoods,omitempty"`
	AllocatedGoods          float64 `json:"allocatedGoods,omitempty"`
	ShippingGoods           float64 `json:"shippingGoods,omitempty"`
}

// AsnOrder represents an inbound order (ASN) in QERP.
type AsnOrder struct {
	AsnNumber              string        `json:"asnNumber"`
	BusinessNumber         string        `json:"businessNumber,omitempty"`
	CustomNumber           string        `json:"customNumber,omitempty"`
	TrackNumber            string        `json:"trackNumber,omitempty"`
	WarehouseName          string        `json:"warehouseName"`
	Type                   string        `json:"type"`
	Status                 string        `json:"status"`
	Remark                 string        `json:"remark,omitempty"`
	CreateTime             string        `json:"createTime"`
	UpdateTime             string        `json:"updateTime,omitempty"`
	StockInTime            string        `json:"stockInTime,omitempty"`
	FinishTime             string        `json:"finishTime,omitempty"`
	LogisticsName          string        `json:"logisticsName,omitempty"`
	ContainerNumber        string        `json:"containerNumber,omitempty"`
	PurchasePriceCurrency  string        `json:"purchasePriceCurrency,omitempty"`
	FirstLegPriceCurrency  string        `json:"firstLegPriceCurrency,omitempty"`
	TransferPriceCurrency  string        `json:"transferPriceCurrency,omitempty"`
	AsnSkuVOList           []AsnSkuItem  `json:"asnSkuVOList,omitempty"`
	AsnCustomFieldValueList []CustomField `json:"asnCustomFieldValueVOList,omitempty"`
	Tag                    *AsnTag       `json:"tag,omitempty"`
}

// AsnTag represents send status flags on an inbound order.
type AsnTag struct {
	SendWms         int `json:"sendWms,omitempty"`
	SendFailed      int `json:"sendFailed,omitempty"`
	ReceiveException int `json:"receiveException,omitempty"`
	BoxedNeedSend   int `json:"boxedNeedSend,omitempty"`
	Overcharge      int `json:"overcharge,omitempty"`
}

// AsnSkuItem represents a SKU line within an inbound order.
type AsnSkuItem struct {
	Sku            string  `json:"sku"`
	Title          string  `json:"title,omitempty"`
	PurchasePrice  float64 `json:"purchasePrice"`
	FirstLegPrice  float64 `json:"firstLegPrice,omitempty"`
	TransferPrice  float64 `json:"transferPrice,omitempty"`
	ExpectQuantity int64   `json:"expectQuantity"`
	ReceiveQuantity int64  `json:"receiveQuantity,omitempty"`
	GoodNum        int64   `json:"goodNum,omitempty"`
	DamageNum      int64   `json:"damageNum,omitempty"`
	SkuNotes       string  `json:"skuNotes,omitempty"`
	PerBoxQuantity int     `json:"perBoxQuantity,omitempty"`
}

// OdoOrder represents an outbound delivery order in QERP.
type OdoOrder struct {
	Number          string        `json:"number"`
	CustomNumber    string        `json:"customNumber,omitempty"`
	TrackNumber     string        `json:"trackNumber,omitempty"`
	WarehouseName   string        `json:"warehouseName"`
	Type            string        `json:"type"`
	Status          string        `json:"status"`
	Remark          string        `json:"remark,omitempty"`
	CreateTime      string        `json:"createTime"`
	FinishTime      string        `json:"finishTime,omitempty"`
	ShippingMethod  string        `json:"shippingMethod,omitempty"`
	OdoSkuVOList    []OdoSkuItem  `json:"odoSkuVOList,omitempty"`
	OdoCustomFieldValueList []CustomField `json:"odoCustomFieldValueVOList,omitempty"`
}

// OdoSkuItem represents a SKU line within an outbound order.
type OdoSkuItem struct {
	Sku                string `json:"sku"`
	Title              string `json:"title,omitempty"`
	StorageLocationCode string `json:"storageLocationCode"`
	ReceiveNumber      string `json:"receiveNumber,omitempty"`
	Quantity           int64  `json:"quantity"`
	UnavailableQuantity int64 `json:"unavailableQuantity,omitempty"`
}

// AdjustmentOrder represents an inventory adjustment order.
type AdjustmentOrder struct {
	AdjustmentNumber   string           `json:"adjustmentNumber"`
	Source             string           `json:"source"`
	AutoSource         string           `json:"autoSource,omitempty"`
	WarehouseName      string           `json:"warehouseName"`
	Remark             string           `json:"remark,omitempty"`
	CreateTime         string           `json:"createTime"`
	AdjustmentSkuList  []AdjustmentSku  `json:"adjustmentSkuVOList,omitempty"`
}

// AdjustmentSku represents a SKU line within an adjustment order.
type AdjustmentSku struct {
	Sku               string `json:"sku"`
	Title             string `json:"title,omitempty"`
	StorageLocationCode string `json:"storageLocationCode,omitempty"`
	Color             string `json:"color,omitempty"`
	ItemPackage       string `json:"itemPackage,omitempty"`
	Size              string `json:"size,omitempty"`
	AvailableQuantity int64  `json:"availableQuantity"`
	AllocatedQuantity int64  `json:"allocatedQuantity"`
	TotalQuantity     int64  `json:"totalQuantity"`
	TotalAfter        int64  `json:"totalAfter"`
	Difference        int64  `json:"difference"`
	AvailableAfter    int64  `json:"availableAfter"`
}

// PurchaseOrder represents a purchase order in QERP.
type PurchaseOrder struct {
	PurchaseNumber        string             `json:"purchaseNumber"`
	AsnNumber             string             `json:"asnNumber,omitempty"`
	CustomNumber          string             `json:"customNumber,omitempty"`
	WarehouseName         string             `json:"warehouseName"`
	TransferWarehouseName string             `json:"transferWarehouseName,omitempty"`
	PurchaseType          string             `json:"purchaseType"`
	PurchaseMode          string             `json:"purchaseMode,omitempty"`
	SupplierName          string             `json:"supplierName"`
	SettlementType        string             `json:"settlementType"`
	PurchasePriceUnit     string             `json:"purchasePriceUnit"`
	PaymentType           string             `json:"paymentType"`
	TransportParty        string             `json:"transportParty"`
	TransportMode         string             `json:"transportMode"`
	Status                string             `json:"status"`
	PrepayRate            float64            `json:"prepayRate,omitempty"`
	ShippingCost          float64            `json:"shippingCost,omitempty"`
	CompanyName           string             `json:"companyName,omitempty"`
	BuyerTitle            string             `json:"buyerTitle,omitempty"`
	OrderTime             string             `json:"orderTime,omitempty"`
	PreReceiveTime        string             `json:"preReceiveTime,omitempty"`
	CreateTime            string             `json:"createTime"`
	UpdateTime            string             `json:"updateTime"`
	Remark                string             `json:"remark,omitempty"`
	SkuList               []PurchaseSkuItem  `json:"skuList,omitempty"`
}

// PurchaseSkuItem represents a SKU line within a purchase order.
type PurchaseSkuItem struct {
	Sku              string  `json:"sku"`
	Title            string  `json:"title,omitempty"`
	PurchasePrice    float64 `json:"purchasePrice"`
	PurchaseQuantity int64   `json:"purchaseQuantity"`
	PackSpecification int64  `json:"packSpecification,omitempty"`
	TaxRate          float64 `json:"taxRate"`
}

// CustomFieldValue represents a custom field value.
type CustomFieldValue struct {
	CustomFieldID int64  `json:"customFieldId"`
	Value         string `json:"value"`
}

// CustomField represents a custom field definition.
type CustomField struct {
	ID             int64    `json:"id,omitempty"`
	TableName      string   `json:"tableName"`
	ColumType      string   `json:"columType"`
	ColumName      string   `json:"columName"`
	DefaultValue   string   `json:"defaultValue,omitempty"`
	CandidateValue []string `json:"candidateValue,omitempty"`
	Remark         string   `json:"remark,omitempty"`
	Required       int      `json:"required,omitempty"`
	IsQuery        int      `json:"isQuery,omitempty"`
	IsShow         int      `json:"isShow,omitempty"`
}
