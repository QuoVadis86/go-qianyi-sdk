package qianyi

import "encoding/json"

// InventoryService provides access to inventory API operations.
type InventoryService struct {
	client *Client
}

// NewInventoryService creates a new InventoryService.
func NewInventoryService(client *Client) *InventoryService {
	return &InventoryService{client: client}
}

// InventoryQueryParams holds parameters for querying inventory.
type InventoryQueryParams struct {
	Page      int      `json:"page"`
	PageSize  int      `json:"pageSize"`
	Warehouse string   `json:"warehouse"`
	SkuList   []string `json:"skuList,omitempty"`
}

// QueryListV2 retrieves a paginated list of inventory records (V2 API).
func (s *InventoryService) QueryListV2(params *InventoryQueryParams) ([]SkuInventory, int, error) {
	biz, _ := json.Marshal(params)
	var list []SkuInventory
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_SIMPLE_LIST_INVENTORY_V2", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// InventoryLogQueryParams holds parameters for querying inventory change logs.
type InventoryLogQueryParams struct {
	OperateTimeFrom string   `json:"operateTimeFrom"`
	OperateTimeTo   string   `json:"operateTimeTo"`
	Page            int      `json:"page"`
	PageSize        int      `json:"pageSize"`
	WarehouseName   string   `json:"warehouseName,omitempty"`
	OperateType     string   `json:"operateType,omitempty"`
	Sku             string   `json:"sku,omitempty"`
	BillNumber      string   `json:"billNumber,omitempty"`
}

// QueryLogList retrieves a paginated list of inventory change logs.
func (s *InventoryService) QueryLogList(params *InventoryLogQueryParams) ([]any, int, error) {
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_INVENTORY_LOG_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}
