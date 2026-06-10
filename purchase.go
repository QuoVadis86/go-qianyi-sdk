package qianyi

import "encoding/json"

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
	Page            int    `json:"page"`
	PageSize        int    `json:"pageSize"`
	PurchaseNumber  string `json:"purchaseNumber,omitempty"`
	CustomNumber    string `json:"customNumber,omitempty"`
	Status          string `json:"status,omitempty"`
	UpdateTimeFrom  string `json:"updateTimeFrom,omitempty"`
	UpdateTimeTo    string `json:"updateTimeTo,omitempty"`
}

// QueryList retrieves purchase orders with optional filters.
func (s *PurchaseService) QueryList(params *PurchaseQueryParams) ([]PurchaseOrder, int, error) {
	biz, _ := json.Marshal(params)
	var list []PurchaseOrder
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_PURCHASE_ORDER_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// PurchaseSkuInput represents a SKU line within a purchase order creation.
type PurchaseSkuInput struct {
	Sku              string  `json:"sku"`
	PurchasePrice    float64 `json:"purchasePrice"`
	PurchaseQuantity int64   `json:"purchaseQuantity"`
	TaxRate          float64 `json:"taxRate"`
	Remark           string  `json:"remark,omitempty"`
}

// CreatePurchaseParams holds parameters for creating a purchase order.
type CreatePurchaseParams struct {
	PurchaseType       string            `json:"purchaseType"`
	WarehouseName      string            `json:"warehouseName"`
	PurchaserName      string            `json:"purchaserName"`
	PurchaseDate       string            `json:"purchaseDate"`
	PurchasePriceUnit  string            `json:"purchasePriceUnit"`
	TransportParty     string            `json:"transportParty"`
	TransportMode      string            `json:"transportMode"`
	SupplierName       string            `json:"supplierName"`
	PaymentType        string            `json:"paymentType"`
	SettlementType     string            `json:"settlementType"`
	PrepayRate         float64           `json:"prepayRate,omitempty"`
	SkuList            []PurchaseSkuInput `json:"skuList"`
	Remark             string            `json:"remark,omitempty"`
	TrackingNumber     string            `json:"trackingNumber,omitempty"`
	CustomNumber       string            `json:"customNumber,omitempty"`
	IsUpdate           bool              `json:"isUpdate,omitempty"`
	PurchaseNumber     string            `json:"purchaseNumber,omitempty"`
}

// Create creates a new purchase order in QERP.
func (s *PurchaseService) Create(params *CreatePurchaseParams) error {
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("CREATE_PURCHASE_ORDER", string(biz), w)
}

// Update updates an existing purchase order. Requires purchaseNumber and isUpdate=true.
func (s *PurchaseService) Update(params *CreatePurchaseParams) error {
	params.IsUpdate = true
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("UPDATE_PURCHASE_ORDER", string(biz), w)
}
