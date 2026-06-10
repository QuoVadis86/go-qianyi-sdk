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

// AsnSku represents a SKU within an inbound order.
type AsnSku struct {
	Sku            string  `json:"sku"`
	ExpectQuantity float64 `json:"expectQuantity"`
}

// CreateAsnParams holds the parameters for creating an inbound order.
type CreateAsnParams struct {
	WarehouseName string  `json:"warehouseName"`
	AsnSkuVOList  []AsnSku `json:"asnSkuVOList"`
	Remark        string  `json:"remark,omitempty"`
}

// Create creates a new inbound order (ASN) in QERP.
func (s *AsnService) Create(params *CreateAsnParams) error {
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("CREATE_ASN_ORDER", string(biz), w)
}

// QueryList retrieves a paginated list of inbound orders.
func (s *AsnService) QueryList(page, pageSize int, warehouse, status string) ([]any, int, error) {
	params := map[string]any{"page": page, "pageSize": pageSize}
	if warehouse != "" {
		params["warehouseName"] = warehouse
	}
	if status != "" {
		params["status"] = status
	}
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_ASN_ORDER_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// Close closes an inbound order by ASN number.
func (s *AsnService) Close(asnNumber string) error {
	params := map[string]any{"asnNumber": asnNumber}
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("CLOSE_ASN_ORDER", string(biz), w)
}
