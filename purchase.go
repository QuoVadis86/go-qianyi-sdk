package qianyi

import "context"

// PurchaseService provides access to purchase order API operations.
type PurchaseService struct {
	client *Client
}

// NewPurchaseService creates a new PurchaseService.
func NewPurchaseService(client *Client) *PurchaseService {
	return &PurchaseService{client: client}
}

// PurchaseQueryParams holds parameters for querying purchase orders.
type PurchaseQueryParams struct {
	Page           int    `json:"page"`
	PageSize       int    `json:"pageSize"`
	PurchaseNumber string `json:"purchaseNumber,omitempty"`
	CustomNumber   string `json:"customNumber,omitempty"`
	Status         string `json:"status,omitempty"`
	UpdateTimeFrom string `json:"updateTimeFrom,omitempty"`
	UpdateTimeTo   string `json:"updateTimeTo,omitempty"`
}

// QueryList retrieves purchase orders with optional filters.
func (s *PurchaseService) QueryList(ctx context.Context, params *PurchaseQueryParams) ([]PurchaseOrder, int, error) {
	return doList[PurchaseOrder](ctx, s.client, "QUERY_PURCHASE_ORDER_LIST", params)
}

// PurchaseSkuInput represents a SKU line within a purchase order creation.
type PurchaseSkuInput struct {
	Sku               string  `json:"sku"`
	PurchasePrice     float64 `json:"purchasePrice"`
	PurchasePriceUnit string  `json:"purchasePriceUnit,omitempty"`
	PurchaseQuantity  int64   `json:"purchaseQuantity"`
	TaxRate           float64 `json:"taxRate"`
	PackSpecification int64   `json:"packSpecification,omitempty"`
	Remark            string  `json:"remark,omitempty"`
}

// PurchaseExtVO represents 1688 extended fields for a purchase order.
type PurchaseExtVO struct {
	Open1688AccountName string `json:"open1688AccountName"`
	Open1688Address     string `json:"open1688Address"`
	Open1688BuyerMsg    string `json:"open1688BuyerMsg,omitempty"`
	Open1688OrderType   string `json:"open1688OrderType"`
	Open1688TradeType   string `json:"open1688TradeType,omitempty"`
}

// CreatePurchaseParams holds parameters for creating a purchase order.
type CreatePurchaseParams struct {
	PurchaseType          string              `json:"purchaseType"`
	WarehouseName         string              `json:"warehouseName"`
	PurchaserName         string              `json:"purchaserName"`
	PurchaseDate          string              `json:"purchaseDate"`
	PurchasePriceUnit     string              `json:"purchasePriceUnit"`
	TransportParty        string              `json:"transportParty"`
	TransportMode         string              `json:"transportMode"`
	SupplierName          string              `json:"supplierName"`
	PaymentType           string              `json:"paymentType"`
	SettlementType        string              `json:"settlementType"`
	SkuList               []PurchaseSkuInput  `json:"skuList"`
	IsUpdate              bool                `json:"isUpdate,omitempty"`
	PurchaseNumber        string              `json:"purchaseNumber,omitempty"`
	CustomNumber          string              `json:"customNumber,omitempty"`
	TransferWarehouseName string              `json:"transferWarehouseName,omitempty"`
	PrepayRate            float64             `json:"prepayRate,omitempty"`
	ShippingCost          float64             `json:"shippingCost,omitempty"`
	CompanyName           string              `json:"companyName,omitempty"`
	BuyerTitle            string              `json:"buyerTitle,omitempty"`
	PreReceiveTime        string              `json:"preReceiveTime,omitempty"`
	Remark                string              `json:"remark,omitempty"`
	TrackingNumber        string              `json:"trackingNumber,omitempty"`
	EffectiveNode         string              `json:"effectiveNode,omitempty"`
	AccountPeriodOpt      string              `json:"accountPeriodOpt,omitempty"`
	BillingDate           int64               `json:"billingDate,omitempty"`
	PurchaseExtVO         *PurchaseExtVO      `json:"purchaseExtVO,omitempty"`
}

// Create creates a new purchase order in QERP.
func (s *PurchaseService) Create(ctx context.Context, params *CreatePurchaseParams) error {
	return doAction(ctx, s.client, "CREATE_PURCHASE_ORDER", params)
}

// Update updates an existing purchase order. Requires purchaseNumber and isUpdate=true.
func (s *PurchaseService) Update(ctx context.Context, params *CreatePurchaseParams) error {
	params.IsUpdate = true
	return doAction(ctx, s.client, "UPDATE_PURCHASE_ORDER", params)
}
