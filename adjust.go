package qianyi

import "encoding/json"

// AdjustService provides access to inventory adjustment API operations.
type AdjustService struct {
	client *Client
}

// NewAdjustService creates a new AdjustService.
func NewAdjustService(client *Client) *AdjustService {
	return &AdjustService{client: client}
}

// AdjustQueryParams holds parameters for querying adjustment orders.
type AdjustQueryParams struct {
	Page            int    `json:"page"`
	PageSize        int    `json:"pageSize"`
	WarehouseName   string `json:"warehouseName,omitempty"`
	Source          string `json:"source"`
	AutoSource      string `json:"autoSource,omitempty"`
	CreateTimeFrom  string `json:"createTimeFrom,omitempty"`
	CreateTimeTo    string `json:"createTimeTo,omitempty"`
	UpdateTimeFrom  string `json:"updateTimeFrom,omitempty"`
	UpdateTimeTo    string `json:"updateTimeTo,omitempty"`
	Number          string `json:"number,omitempty"`
	SkuKeyWord      string `json:"skuKeyWord,omitempty"`
}

// QueryList retrieves adjustment orders with optional filters.
func (s *AdjustService) QueryList(params *AdjustQueryParams) ([]AdjustmentOrder, int, error) {
	biz, _ := json.Marshal(params)
	var list []AdjustmentOrder
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_ADJUSTMENT_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// CreateAdjustParams holds parameters for creating an adjustment order.
type CreateAdjustParams struct {
	WarehouseName   string `json:"warehouseName"`
	ExternalNumber  string `json:"externalNumber,omitempty"`
	Remark          string `json:"remark,omitempty"`
	AdjustmentType  string `json:"adjustmentType,omitempty"`
	AdjustSkuList   []AdjustSkuInput `json:"adjustSkuList"`
}

// AdjustSkuInput represents a SKU adjustment line item.
type AdjustSkuInput struct {
	Sku                string `json:"sku"`
	Title              string `json:"title,omitempty"`
	StorageLocationCode string `json:"storageLocationCode,omitempty"`
	AdjustmentQtyStr   int64  `json:"adjustmentQtyStr"`
}

// Create creates a new inventory adjustment order.
func (s *AdjustService) Create(params *CreateAdjustParams) error {
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("CREATE_ADJUSTMENT_ORDER", string(biz), w)
}
