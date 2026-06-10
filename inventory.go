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

// InventoryQueryV1Params holds parameters for inventory V1 query (deprecated).
type InventoryQueryV1Params struct {
	Page      int      `json:"page"`
	PageSize  int      `json:"pageSize"`
	Warehouse string   `json:"warehouse"`
	SkuList   []string `json:"skuList,omitempty"`
}

// QueryListV1 retrieves inventory list using V1 API (deprecated, use V2).
func (s *InventoryService) QueryListV1(params *InventoryQueryV1Params) ([]SkuInventory, int, error) {
	biz, _ := json.Marshal(params)
	var list []SkuInventory
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_SIMPLE_LIST_INVENTORY", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// InventoryQueryV2Params holds parameters for inventory V2 query.
type InventoryQueryV2Params struct {
	Page            int      `json:"page"`
	PageSize        int      `json:"pageSize"`
	Warehouse       string   `json:"warehouse"`
	SkuList         []string `json:"skuList,omitempty"`
	Warehouses      []string `json:"warehouses,omitempty"`
	ShowCombine     *bool    `json:"showCombine,omitempty"`
	ShowEmpty       *bool    `json:"showEmpty,omitempty"`
	OriginCurrency  *bool    `json:"originCurrency,omitempty"`
	FillCostAndGoods *bool   `json:"fillCostAndGoods,omitempty"`
}

// QueryListV2 retrieves inventory list using V2 API with enhanced options.
func (s *InventoryService) QueryListV2(params *InventoryQueryV2Params) ([]SkuInventory, int, error) {
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
	StorageLocation string   `json:"storageLocation,omitempty"`
	InventoryType   string   `json:"inventoryType,omitempty"`
	TypeList        []string `json:"typeList,omitempty"`
	RecordTypeList  []string `json:"recordTypeList,omitempty"`
}

// InventoryLog represents a single inventory change log entry.
type InventoryLog struct {
	Sku              string `json:"sku"`
	WarehouseName    string `json:"warehouseName"`
	StorageLocation  string `json:"storageLocation"`
	OperateType      string `json:"operateType"`
	BillNumber       string `json:"billNumber"`
	InTransitBefore  int64  `json:"inTransitBefore"`
	InTransitAfter   int64  `json:"inTransitAfter"`
	InTransitDiff    int64  `json:"inTransitDiff"`
	InventoryBefore  int64  `json:"inventoryBefore"`
	InventoryAfter   int64  `json:"inventoryAfter"`
	InventoryDiff    int64  `json:"inventoryDiff"`
	BatchNumber      string `json:"batchNumber,omitempty"`
	Operator         string `json:"operator"`
	OperateTime      string `json:"operateTime"`
	BizType          string `json:"bizType"`
	Cost             float64 `json:"cost,omitempty"`
	UniqueID         int64  `json:"uniqueId"`
}

// QueryLogList retrieves a paginated list of inventory change logs.
func (s *InventoryService) QueryLogList(params *InventoryLogQueryParams) ([]InventoryLog, int, error) {
	biz, _ := json.Marshal(params)
	var list []InventoryLog
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_INVENTORY_LOG_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// AssemblyQueryParams holds parameters for querying assembly orders.
type AssemblyQueryParams struct {
	AssemblyNumber  string `json:"assemblyNumber,omitempty"`
	CreateTimeFrom  string `json:"createTimeFrom,omitempty"`
	CreateTimeTo    string `json:"createTimeTo,omitempty"`
	WarehouseName   string `json:"warehouseName,omitempty"`
	Page            int    `json:"page"`
	PageSize        int    `json:"pageSize"`
}

// QueryAssemblyList retrieves assembly orders.
func (s *InventoryService) QueryAssemblyList(params *AssemblyQueryParams) ([]any, int, error) {
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_INVENTORY_ASSEMBLY_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// CreateTransferOrderParams holds parameters for creating a transfer order.
type CreateTransferOrderParams struct {
	WarehouseFrom string `json:"warehouseFrom"`
	WarehouseTo   string `json:"warehouseTo"`
	Sku           string `json:"sku"`
	Quantity      int    `json:"quantity"`
}

// CreateTransferOrder creates a transfer order between warehouses.
func (s *InventoryService) CreateTransferOrder(params *CreateTransferOrderParams) error {
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("CREATE_TRANSFER_ORDER", string(biz), w)
}

// QueryTransferOrderList queries transfer orders.
func (s *InventoryService) QueryTransferOrderList(page, pageSize int) ([]any, int, error) {
	params := map[string]any{"page": page, "pageSize": pageSize}
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_TRANSFER_ORDER_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// QuerySplitOrderList queries split orders.
func (s *InventoryService) QuerySplitOrderList(page, pageSize int) ([]any, int, error) {
	params := map[string]any{"page": page, "pageSize": pageSize}
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_SPLIT_ORDER_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// QueryStorageLocInventory queries inventory by storage location.
func (s *InventoryService) QueryStorageLocInventory(warehouse, storageLocation string, skuList []string) ([]any, error) {
	params := map[string]any{
		"warehouse":       warehouse,
		"storageLocation": storageLocation,
	}
	if len(skuList) > 0 {
		params["skuList"] = skuList
	}
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_STORAGE_LOC_INVENTORY", string(biz), w); err != nil {
		return nil, err
	}
	if w.HasError() {
		return nil, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, nil
}

// QueryBatchInventoryList queries batch-level inventory.
func (s *InventoryService) QueryBatchInventoryList(receiveTimeFrom, receiveTimeTo string, page, pageSize int) ([]any, int, error) {
	params := map[string]any{
		"receiveTimeFrom": receiveTimeFrom,
		"receiveTimeTo":   receiveTimeTo,
		"page":           page,
		"pageSize":       pageSize,
	}
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_BATCH_INVENTORY_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// TransferStorageLocationParams holds parameters for transferring goods between storage locations.
type TransferStorageLocationParams struct {
	FromLocation string `json:"fromLocation"`
	ToLocation   string `json:"toLocation"`
	SkuList      []struct {
		Sku      string `json:"sku"`
		Quantity int    `json:"quantity"`
	} `json:"skuList"`
}

// TransferStorageLocation transfers goods between storage locations.
func (s *InventoryService) TransferStorageLocation(params *TransferStorageLocationParams) error {
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("TRANSFER_STORAGE_LOCATION", string(biz), w)
}

// QuerySBSInventoryList queries Shopee SBS inventory.
func (s *InventoryService) QuerySBSInventoryList(warehouseID int64, page, pageSize int) ([]any, int, error) {
	params := map[string]any{
		"warehouseId": warehouseID,
		"page":       page,
		"pageSize":   pageSize,
	}
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_SBS_INVENTORY_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// QuerySBSWarehouseList queries Shopee SBS warehouse list.
func (s *InventoryService) QuerySBSWarehouseList() ([]any, error) {
	params := map[string]any{}
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_SBS_WAREHOUSE_LIST", string(biz), w); err != nil {
		return nil, err
	}
	if w.HasError() {
		return nil, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, nil
}
