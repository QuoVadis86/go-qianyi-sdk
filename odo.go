package qianyi

import "encoding/json"

// OdoService provides access to outbound delivery order (ODO) API operations.
type OdoService struct {
	client *Client
}

// NewOdoService creates a new OdoService.
func NewOdoService(client *Client) *OdoService {
	return &OdoService{client: client}
}

// OdoQueryParams holds parameters for querying outbound orders.
type OdoQueryParams struct {
	Page              int      `json:"page"`
	PageSize          int      `json:"pageSize"`
	WarehouseName     string   `json:"warehouseName,omitempty"`
	Type              string   `json:"type,omitempty"`
	Status            string   `json:"status,omitempty"`
	CreateTimeFrom    string   `json:"createTimeFrom,omitempty"`
	CreateTimeTo      string   `json:"createTimeTo,omitempty"`
	ShipTimeFrom      string   `json:"shipTimeFrom,omitempty"`
	ShipTimeTo        string   `json:"shipTimeTo,omitempty"`
	UpdateTimeFrom    string   `json:"updateTimeFrom,omitempty"`
	UpdateTimeTo      string   `json:"updateTimeTo,omitempty"`
	NumberParam       string   `json:"numberParam,omitempty"`
	SkuParam          string   `json:"skuParam,omitempty"`
	OrderNumberList   []string `json:"orderNumberList,omitempty"`
	CustomNumberList  []string `json:"customNumberList,omitempty"`
	CalculateCost     *bool    `json:"calculateCost,omitempty"`
}

// QueryList retrieves outbound delivery orders with optional filters.
func (s *OdoService) QueryList(params *OdoQueryParams) ([]OdoOrder, int, error) {
	biz, _ := json.Marshal(params)
	var list []OdoOrder
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_ODO_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// OdoReceiver holds recipient information for an outbound order.
type OdoReceiver struct {
	Name          string `json:"name"`
	StreetLine1   string `json:"streetLine1"`
	PostalCode    string `json:"postalCode"`
	CountryCode   string `json:"countryCode"`
	City          string `json:"city"`
	MobileNumber  string `json:"mobileNumber"`
	State         string `json:"state"`
	ShopCode      string `json:"shopCode,omitempty"`
}

// OdoSkuCreateItem represents a SKU line in outbound order creation.
type OdoSkuCreateItem struct {
	Sku                string `json:"sku"`
	Quantity           int64  `json:"quantity"`
	UnavailableQuantity int64 `json:"unavailableQuantity,omitempty"`
	BatchNo            string `json:"batchNo,omitempty"`
	MfgDate            string `json:"mfgDate,omitempty"`
	ExpDate            string `json:"expDate,omitempty"`
	OriginCountry      string `json:"originCountry,omitempty"`
	ApiCustom          string `json:"apiCustom,omitempty"`
	PartNo             string `json:"partNo,omitempty"`
}

// CreateOdoParams holds parameters for creating an outbound order.
type CreateOdoParams struct {
	WarehouseName      string              `json:"warehouseName"`
	CustomNumber       string              `json:"customNumber"`
	Remark             string              `json:"remark,omitempty"`
	Carrier            string              `json:"carrier,omitempty"`
	CarrierService     string              `json:"carrierService,omitempty"`
	SecondaryType      string              `json:"secondaryType,omitempty"`
	TrackingNumber     string              `json:"trackingNumber,omitempty"`
	ShippingMethod     string              `json:"shippingMethod,omitempty"`
	ShippingLabelSource string             `json:"shippingLabelSource,omitempty"`
	IsSpecifyBatch     bool                `json:"isSpecifyBatch,omitempty"`
	CustomerType       string              `json:"customerType,omitempty"`
	PurchaseOrderNumber string             `json:"purchaseOrderNumber,omitempty"`
	Receiver           *OdoReceiver        `json:"receiver,omitempty"`
	OdoSkuVOList       []OdoSkuCreateItem  `json:"odoSkuVOList"`
	OdoCustomFieldValues []CustomFieldValue `json:"odoCustomFieldValueVOList,omitempty"`
}

// Create creates a new outbound delivery order in QERP.
func (s *OdoService) Create(params *CreateOdoParams) error {
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("CREATE_ODO_ORDER", string(biz), w)
}

// Cancel cancels an outbound delivery order by custom number.
func (s *OdoService) Cancel(customNumber string) error {
	params := map[string]any{"customNumber": customNumber}
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("CANCEL_ODO_ORDER", string(biz), w)
}

// PushOdoParams holds parameters for pushing outbound order status.
type PushOdoParams struct {
	Number        string        `json:"number"`
	CustomNumber  string        `json:"customNumber,omitempty"`
	TrackNumber   string        `json:"trackNumber"`
	WarehouseName string        `json:"warehouseName"`
	Type          string        `json:"type"`
	Status        string        `json:"status"`
	ShippingTime  string        `json:"shippingTime,omitempty"`
	Carrier       string        `json:"carrier,omitempty"`
	SkuList       []PushOdoSku  `json:"skuList"`
}

// PushOdoSku represents a SKU in outbound push notification.
type PushOdoSku struct {
	Sku                string `json:"sku"`
	Title              string `json:"title"`
	Quantity           int64  `json:"quantity"`
	UnavailableQuantity int64 `json:"unavailableQuantity,omitempty"`
	ApiCustom          string `json:"apiCustom,omitempty"`
}

// PushOrder pushes outbound delivery order status notification.
func (s *OdoService) PushOrder(params *PushOdoParams) error {
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("PUSH_ODO_ORDER", string(biz), w)
}

// QuerySalesList queries sales-related outbound orders.
func (s *OdoService) QuerySalesList(createTimeFrom, createTimeTo string, page, pageSize int) ([]OdoOrder, int, error) {
	params := map[string]any{
		"createTimeFrom": createTimeFrom,
		"createTimeTo":   createTimeTo,
		"page":          page,
		"pageSize":      pageSize,
	}
	biz, _ := json.Marshal(params)
	var list []OdoOrder
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_SALES_ODO_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}
