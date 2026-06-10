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
	Page            int    `json:"page"`
	PageSize        int    `json:"pageSize"`
	WarehouseName   string `json:"warehouseName,omitempty"`
	Type            string `json:"type,omitempty"`
	Status          string `json:"status,omitempty"`
	CreateTimeFrom  string `json:"createTimeFrom,omitempty"`
	CreateTimeTo    string `json:"createTimeTo,omitempty"`
	ShipTimeFrom    string `json:"shipTimeFrom,omitempty"`
	ShipTimeTo      string `json:"shipTimeTo,omitempty"`
	UpdateTimeFrom  string `json:"updateTimeFrom,omitempty"`
	UpdateTimeTo    string `json:"updateTimeTo,omitempty"`
	NumberParam     string `json:"numberParam,omitempty"`
	SkuParam        string `json:"skuParam,omitempty"`
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

// CreateOdoParams holds parameters for creating an outbound order.
type CreateOdoParams struct {
	WarehouseName    string         `json:"warehouseName"`
	CustomNumber     string         `json:"customNumber"`
	Remark           string         `json:"remark,omitempty"`
	Carrier          string         `json:"carrier,omitempty"`
	CarrierService   string         `json:"carrierService,omitempty"`
	TrackingNumber   string         `json:"trackingNumber,omitempty"`
	IsSpecifyBatch   bool           `json:"isSpecifyBatch,omitempty"`
	CustomerType     string         `json:"customerType,omitempty"`
	OdoSkuVOList     []OdoSkuItem   `json:"odoSkuVOList"`
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
