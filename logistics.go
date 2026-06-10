package qianyi

import "encoding/json"

// LogisticsService provides access to first-leg logistics API operations.
type LogisticsService struct {
	client *Client
}

// NewLogisticsService creates a new LogisticsService.
func NewLogisticsService(client *Client) *LogisticsService {
	return &LogisticsService{client: client}
}

// FirstLegOrder represents a first-leg logistics order.
type FirstLegOrder struct {
	FirstLegNumber    string             `json:"firstLegNumber"`
	AsnNumber         string             `json:"asnNumber"`
	CustomNumber      string             `json:"customNumber"`
	Status            string             `json:"status"`
	CreateTime        string             `json:"createTime"`
	UpdateTime        string             `json:"updateTime"`
	WarehouseName     string             `json:"warehouseName"`
	DestWarehouseName string             `json:"destWarehouseName"`
	PortFrom          string             `json:"portFrom,omitempty"`
	PortTo            string             `json:"portTo,omitempty"`
	LogisticsName     string             `json:"logisticsName,omitempty"`
	CarrierName       string             `json:"carrierName,omitempty"`
	BuyerTitle        string             `json:"buyerTitle,omitempty"`
	Forwarder         string             `json:"forwarder,omitempty"`
	TrackNumber       string             `json:"trackNumber,omitempty"`
	PreReceiveTime    string             `json:"preReceiveTime,omitempty"`
	PreShipTime       string             `json:"preShipTime,omitempty"`
	ShippingTime      string             `json:"shippingTime,omitempty"`
	Remark            string             `json:"remark,omitempty"`
	FeeList           []FirstLegFee      `json:"feeList,omitempty"`
	SkuList           []FirstLegOrderSku `json:"skuList,omitempty"`
}

// FirstLegFee represents a fee entry in a first-leg order.
type FirstLegFee struct {
	Name          string  `json:"name"`
	Amount        float64 `json:"amount"`
	AppliedAmount float64 `json:"appliedAmount"`
	PaidAmount    float64 `json:"paidAmount"`
	InvalidAmount float64 `json:"invalidAmount,omitempty"`
	Status        string  `json:"status"`
}

// FirstLegOrderSku represents a SKU line in a first-leg order.
type FirstLegOrderSku struct {
	Sku               string  `json:"sku"`
	Title             string  `json:"title"`
	ExpectedQuantity  int64   `json:"expectedQuantity"`
	ReceiveQuantity   int64   `json:"receiveQuantity"`
	PackSpecification int64   `json:"packSpecification"`
	NetWeight         float64 `json:"netWeight"`
	Weight            float64 `json:"weight"`
	WeightUnit        string  `json:"weightUnit"`
	Length            float64 `json:"length"`
	Width             float64 `json:"width"`
	Height            float64 `json:"height"`
	DimensionUnit     string  `json:"dimensionUnit"`
	PurchasePrice     float64 `json:"purchasePrice"`
	OriginAsnNumber   string  `json:"originAsnNumber"`
}

// FirstLegQueryParams holds params for querying first-leg orders.
type FirstLegQueryParams struct {
	FirstLegNumber  string `json:"firstLegNumber,omitempty"`
	Status          string `json:"status,omitempty"`
	UpdateTimeFrom  string `json:"updateTimeFrom,omitempty"`
	UpdateTimeTo    string `json:"updateTimeTo,omitempty"`
	Page            int    `json:"page"`
	PageSize        int    `json:"pageSize"`
}

// QueryFirstLegList retrieves first-leg logistics orders.
func (s *LogisticsService) QueryFirstLegList(params *FirstLegQueryParams) ([]FirstLegOrder, int, error) {
	biz, _ := json.Marshal(params)
	var list []FirstLegOrder
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_FIRST_LEG_ORDER_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// FirstLegSkuDetail represents a SKU detail within a first-leg order creation.
type FirstLegSkuDetail struct {
	LineID              int    `json:"lineId"`
	WarehouseName       string `json:"warehouseName"`
	DestWarehouseName   string `json:"destWarehouseName"`
	LogisticsName       string `json:"logisticsName"`
	Sku                 string `json:"sku"`
	PreExpectedQuantity int    `json:"preExpectedQuantity"`
	FbaNo               string `json:"fbaNo,omitempty"`
	CustomNumber        string `json:"customNumber,omitempty"`
	PortFrom            string `json:"portFrom,omitempty"`
	PortTo              string `json:"portTo,omitempty"`
	BuyerTitle          string `json:"buyerTitle,omitempty"`
	TrackNumber         string `json:"trackNumber,omitempty"`
	PreReceiveTime      string `json:"preReceiveTime,omitempty"`
	PreShipTime         string `json:"preShipTime,omitempty"`
	Remark              string `json:"remark,omitempty"`
	ContainerNumber     string `json:"containerNumber,omitempty"`
	OriginalAsnNumber   string `json:"originalAsnNumber,omitempty"`
	PackingRate         int    `json:"packingRate,omitempty"`
	NetWeight           float64 `json:"netWeight,omitempty"`
	Length              float64 `json:"length,omitempty"`
	Width               float64 `json:"width,omitempty"`
	Height              float64 `json:"height,omitempty"`
	SkuRemark           string `json:"skuRemark,omitempty"`
	RefID               string `json:"refId,omitempty"`
}

// CreateFirstLegParams holds parameters for creating a first-leg logistics order.
type CreateFirstLegParams struct {
	DestWarehouseType string              `json:"destWarehouseType"`
	SkuDetailList     []FirstLegSkuDetail `json:"skuDetailList"`
}

// CreateFirstLeg creates a first-leg logistics order.
func (s *LogisticsService) CreateFirstLeg(params *CreateFirstLegParams) error {
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("CREATE_FIRST_LEG_ORDER", string(biz), w)
}

// FirstLegLogisticsInfo represents a logistics option for first-leg.
type FirstLegLogisticsInfo struct {
	ID            int64  `json:"id"`
	LogisticsName string `json:"logisticsName"`
}

// QueryFirstLegLogistics queries available first-leg logistics for a warehouse.
func (s *LogisticsService) QueryFirstLegLogistics(warehouseName string) ([]FirstLegLogisticsInfo, error) {
	params := map[string]any{"warehouseName": warehouseName}
	biz, _ := json.Marshal(params)
	var list []FirstLegLogisticsInfo
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_FIRST_LRG_LOGISTICS", string(biz), w); err != nil {
		return nil, err
	}
	if w.HasError() {
		return nil, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, nil
}

// TrackingPackage represents a logistics tracking package.
type TrackingPackage struct {
	OrderNumber    string        `json:"orderNumber"`
	OnlineOrderID  string        `json:"onlineOrderId"`
	TrackingNumber string        `json:"trackingNumber"`
	Carrier        string        `json:"carrier"`
	Status         string        `json:"status"`
	EventList      []TrackingEvent `json:"eventList,omitempty"`
}

// TrackingEvent represents a single tracking event.
type TrackingEvent struct {
	EventDate  int64  `json:"eventDate"`
	Event      string `json:"event"`
	TimeZoneID string `json:"timeZoneId"`
}

// TrackingQueryParams holds params for querying tracking packages.
type TrackingQueryParams struct {
	UpdateTimeFrom  string   `json:"updateTimeFrom"`
	UpdateTimeTo    string   `json:"updateTimeTo"`
	Page            int      `json:"page"`
	PageSize        int      `json:"pageSize"`
	OrderNumbers    []string `json:"orderNumbers,omitempty"`
	OnlineOrderIDs  []string `json:"onlineOrderIds,omitempty"`
	TrackingNumbers []string `json:"trackingNumbers,omitempty"`
	ReturnDetails   *bool    `json:"returnDetails,omitempty"`
}

// QueryFirstLegTracking queries first-leg tracking packages.
func (s *LogisticsService) QueryFirstLegTracking(params *TrackingQueryParams) ([]TrackingPackage, int, error) {
	biz, _ := json.Marshal(params)
	var list []TrackingPackage
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_FIRST_LRG_TRACKING_PACKAGE", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// WithdrawFirstLegParams holds parameters for withdrawing a first-leg order.
type WithdrawFirstLegParams struct {
	FirstLegNumber string `json:"firstLegNumber,omitempty"`
	CustomNumber   string `json:"customNumber,omitempty"`
	DelFlag        *bool  `json:"delFlag,omitempty"`
}

// WithdrawFirstLeg withdraws or deletes a first-leg logistics order.
func (s *LogisticsService) WithdrawFirstLeg(params *WithdrawFirstLegParams) error {
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("WITHDRAW_AND_DEL_FIRST_LEG", string(biz), w)
}

// PushTrackingParams holds parameters for pushing tracking status.
type PushTrackingParams struct {
	OrderNumber     string `json:"orderNumber"`
	OnlineOrderID   string `json:"onlineOrderId"`
	TrackingNumber  string `json:"trackingNumber"`
	Carrier         string `json:"carrier"`
	Status          string `json:"status"`
}

// PushTrackingPackage pushes logistics tracking status notification.
func (s *LogisticsService) PushTrackingPackage(params *PushTrackingParams) error {
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("PUSH_TRACKING_PACKAGE", string(biz), w)
}
