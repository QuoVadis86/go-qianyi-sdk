package qianyi

import "context"

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
	Page           int    `json:"page"`
	PageSize       int    `json:"pageSize"`
	WarehouseName  string `json:"warehouseName,omitempty"`
	Source         string `json:"source"`
	AutoSource     string `json:"autoSource,omitempty"`
	CreateTimeFrom string `json:"createTimeFrom,omitempty"`
	CreateTimeTo   string `json:"createTimeTo,omitempty"`
	UpdateTimeFrom string `json:"updateTimeFrom,omitempty"`
	UpdateTimeTo   string `json:"updateTimeTo,omitempty"`
	Number         string `json:"number,omitempty"`
	SkuKeyWord     string `json:"skuKeyWord,omitempty"`
}

// QueryList retrieves adjustment orders with optional filters.
func (s *AdjustService) QueryList(ctx context.Context, params *AdjustQueryParams) ([]AdjustmentOrder, int, error) {
	return doList[AdjustmentOrder](ctx, s.client, "QUERY_ADJUSTMENT_LIST", params)
}

// CreateAdjustParams holds parameters for creating an adjustment order.
type CreateAdjustParams struct {
	WarehouseName  string            `json:"warehouseName"`
	ExternalNumber string            `json:"externalNumber,omitempty"`
	Remark         string            `json:"remark,omitempty"`
	AdjustmentType string            `json:"adjustmentType,omitempty"`
	AdjustSkuList  []AdjustSkuInput  `json:"adjustSkuList"`
}

// AdjustSkuInput represents a SKU adjustment line item.
type AdjustSkuInput struct {
	Sku                string `json:"sku"`
	Title              string `json:"title,omitempty"`
	StorageLocationCode string `json:"storageLocationCode,omitempty"`
	AdjustmentQtyStr   int64  `json:"adjustmentQtyStr"`
}

// Create creates a new inventory adjustment order.
func (s *AdjustService) Create(ctx context.Context, params *CreateAdjustParams) error {
	return doAction(ctx, s.client, "CREATE_ADJUSTMENT_ORDER", params)
}
