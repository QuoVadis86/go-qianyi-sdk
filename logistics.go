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
	FirstLegNumber    string            `json:"firstLegNumber"`
	AsnNumber         string            `json:"asnNumber"`
	CustomNumber      string            `json:"customNumber"`
	Status            string            `json:"status"`
	CreateTime        string            `json:"createTime"`
	UpdateTime        string            `json:"updateTime"`
	WarehouseName     string            `json:"warehouseName"`
	DestWarehouseName string            `json:"destWarehouseName"`
	PortFrom          string            `json:"portFrom,omitempty"`
	PortTo            string            `json:"portTo,omitempty"`
	LogisticsName     string            `json:"logisticsName,omitempty"`
	CarrierName       string            `json:"carrierName,omitempty"`
	TrackNumber       string            `json:"trackNumber,omitempty"`
	PreReceiveTime    string            `json:"preReceiveTime,omitempty"`
	ShippingTime      string            `json:"shippingTime,omitempty"`
	Remark            string            `json:"remark,omitempty"`
}

// QueryFirstLegList retrieves first-leg logistics orders.
func (s *LogisticsService) QueryFirstLegList(firstLegNumber, status string, page, pageSize int) ([]FirstLegOrder, int, error) {
	params := map[string]any{"page": page, "pageSize": pageSize}
	if firstLegNumber != "" {
		params["firstLegNumber"] = firstLegNumber
	}
	if status != "" {
		params["status"] = status
	}
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

// CreateFirstLegParams holds parameters for creating a first-leg logistics order.
type CreateFirstLegParams struct {
	DestWarehouseType string              `json:"destWarehouseType"`
	SkuDetailList     []FirstLegSkuDetail `json:"skuDetailList"`
}

// FirstLegSkuDetail represents a SKU detail within a first-leg order.
type FirstLegSkuDetail struct {
	LineID              int    `json:"lineId"`
	WarehouseName       string `json:"warehouseName"`
	DestWarehouseName   string `json:"destWarehouseName"`
	LogisticsName       string `json:"logisticsName"`
	Sku                 string `json:"sku"`
	PreExpectedQuantity int    `json:"preExpectedQuantity"`
	CustomNumber        string `json:"customNumber,omitempty"`
	TrackNumber         string `json:"trackNumber,omitempty"`
	Remark              string `json:"remark,omitempty"`
}

// CreateFirstLeg creates a first-leg logistics order.
func (s *LogisticsService) CreateFirstLeg(params *CreateFirstLegParams) error {
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("CREATE_FIRST_LEG_ORDER", string(biz), w)
}

// QueryFirstLegLogistics queries available first-leg logistics for a warehouse.
func (s *LogisticsService) QueryFirstLegLogistics(warehouseName string) ([]any, error) {
	params := map[string]any{"warehouseName": warehouseName}
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_FIRST_LRG_LOGISTICS", string(biz), w); err != nil {
		return nil, err
	}
	if w.HasError() {
		return nil, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, nil
}

// QueryFirstLegTracking queries first-leg tracking packages.
func (s *LogisticsService) QueryFirstLegTracking(updateTimeFrom, updateTimeTo string, page, pageSize int) ([]any, int, error) {
	params := map[string]any{
		"updateTimeFrom": updateTimeFrom,
		"updateTimeTo":   updateTimeTo,
		"page":          page,
		"pageSize":      pageSize,
	}
	biz, _ := json.Marshal(params)
	var list []any
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
