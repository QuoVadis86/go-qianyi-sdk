package qianyi

import "encoding/json"

// AsnService provides access to inbound order (ASN) API operations.
type AsnService struct {
	client *Client
}

// NewAsnService creates a new AsnService.
func NewAsnService(client *Client) *AsnService {
	return &AsnService{client: client}
}

// AsnSku represents a SKU within an inbound order creation request.
type AsnSku struct {
	Sku            string  `json:"sku"`
	PurchasePrice  float64 `json:"purchasePrice"`
	FirstLegPrice  float64 `json:"firstLegPrice,omitempty"`
	TransferPrice  float64 `json:"transferPrice,omitempty"`
	ExpectQuantity int64   `json:"expectQuantity"`
	SkuStatus      string  `json:"skuStatus,omitempty"`
	PerBoxQuantity int     `json:"perBoxQuantity,omitempty"`
	BatchNo        string  `json:"batchNo,omitempty"`
	MfgDate        string  `json:"mfgDate,omitempty"`
	ExpDate        string  `json:"expDate,omitempty"`
	OriginCountry  string  `json:"originCountry,omitempty"`
	SkuNotes       string  `json:"skuNotes,omitempty"`
}

// CreateAsnParams holds the parameters for creating an inbound order.
type CreateAsnParams struct {
	WarehouseName          string           `json:"warehouseName"`
	AsnSkuVOList           []AsnSku         `json:"asnSkuVOList"`
	CustomNumber           string           `json:"customNumber,omitempty"`
	TrackNumber            string           `json:"trackNumber,omitempty"`
	Remark                 string           `json:"remark,omitempty"`
	PurchasePriceCurrency  string           `json:"purchasePriceCurrency"`
	FirstLegPriceCurrency  string           `json:"firstLegPriceCurrency,omitempty"`
	TransferPriceCurrency  string           `json:"transferPriceCurrency,omitempty"`
	SendWarehouseFlag      string           `json:"sendWarehouseFlag,omitempty"`
	PreArriveTime          string           `json:"preArriveTime,omitempty"`
	CustomerType           string           `json:"customerType,omitempty"`
	IsSpecifyBatch         bool             `json:"isSpecifyBatch,omitempty"`
}

// Create creates a new inbound order (ASN) in QERP.
func (s *AsnService) Create(params *CreateAsnParams) error {
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("CREATE_ASN_ORDER", string(biz), w)
}

// AsnQueryParams holds parameters for querying inbound orders.
type AsnQueryParams struct {
	Page          int    `json:"page"`
	PageSize      int    `json:"pageSize"`
	WarehouseName string `json:"warehouseName,omitempty"`
	Type          string `json:"type,omitempty"`
	Status        string `json:"status,omitempty"`
	SkuKeyWord    string `json:"skuKeyWord,omitempty"`
	Number        string `json:"number,omitempty"`
	TrackNumber   string `json:"trackNumber,omitempty"`
	TimeType      string `json:"timeType,omitempty"`
	TimeFrom      string `json:"timeFrom,omitempty"`
	TimeEnd       string `json:"timeEnd,omitempty"`
	ReturnBatchInfo bool `json:"returnBatchInfo,omitempty"`
}

// QueryList retrieves inbound orders with optional filters.
func (s *AsnService) QueryList(params *AsnQueryParams) ([]AsnOrder, int, error) {
	biz, _ := json.Marshal(params)
	var list []AsnOrder
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_ASN_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// Cancel cancels an inbound order by ASN number.
func (s *AsnService) Cancel(asnNumber string) error {
	params := map[string]any{"asnNumber": asnNumber}
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("CANCEL_ASN_ORDER", string(biz), w)
}

// Delete deletes an inbound order by ASN number.
func (s *AsnService) Delete(asnNumber string) error {
	params := map[string]any{"asnNumber": asnNumber}
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("DELETE_ASN_ORDER", string(biz), w)
}

// QueryBatchList queries inventory batch records for inbound orders.
func (s *AsnService) QueryBatchList(receiveTimeFrom, receiveTimeTo string, page, pageSize int) ([]any, int, error) {
	params := map[string]any{
		"receiveTimeFrom": receiveTimeFrom,
		"receiveTimeTo":   receiveTimeTo,
		"page":           page,
		"pageSize":       pageSize,
	}
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_ASN_BATCH_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}
