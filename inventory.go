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

// QueryListV1 retrieves inventory list using V1 API (deprecated).
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

// QueryListV2 retrieves inventory list using V2 API.
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
	Sku              string  `json:"sku"`
	WarehouseName    string  `json:"warehouseName"`
	StorageLocation  string  `json:"storageLocation"`
	OperateType      string  `json:"operateType"`
	BillNumber       string  `json:"billNumber"`
	InTransitBefore  int64   `json:"inTransitBefore"`
	InTransitAfter   int64   `json:"inTransitAfter"`
	InTransitDiff    int64   `json:"inTransitDiff"`
	InventoryBefore  int64   `json:"inventoryBefore"`
	InventoryAfter   int64   `json:"inventoryAfter"`
	InventoryDiff    int64   `json:"inventoryDiff"`
	BatchNumber      string  `json:"batchNumber,omitempty"`
	Operator         string  `json:"operator"`
	OperateTime      string  `json:"operateTime"`
	BizType          string  `json:"bizType"`
	Cost             float64 `json:"cost,omitempty"`
	UniqueID         int64   `json:"uniqueId"`
}

// QueryLogList retrieves inventory change logs.
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

// AssemblyQueryParams holds params for querying assembly orders.
type AssemblyQueryParams struct {
	AssemblyNumber string `json:"assemblyNumber,omitempty"`
	CreateTimeFrom string `json:"createTimeFrom,omitempty"`
	CreateTimeTo   string `json:"createTimeTo,omitempty"`
	WarehouseName  string `json:"warehouseName,omitempty"`
	Page           int    `json:"page"`
	PageSize       int    `json:"pageSize"`
}

// AssemblyOrder represents an inventory assembly order.
type AssemblyOrder struct {
	AssemblyNumber string           `json:"assemblyNumber,omitempty"`
	WarehouseName  string           `json:"warehouseName,omitempty"`
	Status         string           `json:"status,omitempty"`
	AsnNumber      string           `json:"asnNumber,omitempty"`
	FinishTime     string           `json:"finishTime,omitempty"`
	CreateTime     string           `json:"createTime,omitempty"`
	AssemblyList   []AssemblySku    `json:"assemblyList,omitempty"`
}

// AssemblySku represents a SKU line in an assembly order.
type AssemblySku struct {
	Sku      string `json:"sku,omitempty"`
	Title    string `json:"title,omitempty"`
	Quantity int64  `json:"quantity,omitempty"`
}

// QueryAssemblyList retrieves assembly orders.
func (s *InventoryService) QueryAssemblyList(params *AssemblyQueryParams) ([]AssemblyOrder, int, error) {
	biz, _ := json.Marshal(params)
	var list []AssemblyOrder
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_INVENTORY_ASSEMBLY_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// CreateTransferOrderParams holds params for creating a transfer order.
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

// TransferOrder represents a warehouse transfer order.
type TransferOrder struct {
	Number       string `json:"number,omitempty"`
	Status       string `json:"status,omitempty"`
	SourceWarehouse string `json:"sourceWarehouse,omitempty"`
	TargetWarehouse string `json:"targetWarehouse,omitempty"`
	Sku          string `json:"sku,omitempty"`
	Quantity     int64  `json:"quantity,omitempty"`
	CreateTime   string `json:"createTime,omitempty"`
}

// QueryTransferOrderList queries transfer orders.
func (s *InventoryService) QueryTransferOrderList(page, pageSize int) ([]TransferOrder, int, error) {
	params := map[string]any{"page": page, "pageSize": pageSize}
	biz, _ := json.Marshal(params)
	var list []TransferOrder
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_TRANSFER_ORDER_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// SplitOrder represents an order split record.
type SplitOrder struct {
	OrderNumber     string `json:"orderNumber,omitempty"`
	SplitOrderNumber string `json:"splitOrderNumber,omitempty"`
	Sku             string `json:"sku,omitempty"`
	Quantity        int64  `json:"quantity,omitempty"`
	Status          string `json:"status,omitempty"`
	CreateTime      string `json:"createTime,omitempty"`
}

// QuerySplitOrderList queries split orders.
func (s *InventoryService) QuerySplitOrderList(page, pageSize int) ([]SplitOrder, int, error) {
	params := map[string]any{"page": page, "pageSize": pageSize}
	biz, _ := json.Marshal(params)
	var list []SplitOrder
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_SPLIT_ORDER_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// StorageLocInventory represents inventory in a specific storage location.
type StorageLocInventory struct {
	Sku             string `json:"sku,omitempty"`
	SkuName         string `json:"skuName,omitempty"`
	Warehouse       string `json:"warehouse,omitempty"`
	StorageLocation string `json:"storageLocation,omitempty"`
	Quantity        int64  `json:"quantity,omitempty"`
	Available       int64  `json:"available,omitempty"`
}

// QueryStorageLocInventory queries inventory by storage location.
func (s *InventoryService) QueryStorageLocInventory(warehouse, storageLocation string, skuList []string) ([]StorageLocInventory, error) {
	params := map[string]any{"warehouse": warehouse, "storageLocation": storageLocation}
	if len(skuList) > 0 {
		params["skuList"] = skuList
	}
	biz, _ := json.Marshal(params)
	var list []StorageLocInventory
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_STORAGE_LOC_INVENTORY", string(biz), w); err != nil {
		return nil, err
	}
	if w.HasError() {
		return nil, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, nil
}

// BatchInventory represents a batch inventory record.
type BatchInventory struct {
	Sku             string `json:"sku,omitempty"`
	SkuName         string `json:"skuName,omitempty"`
	WarehouseName   string `json:"warehouseName,omitempty"`
	BatchNumber     string `json:"batchNumber,omitempty"`
	Quantity        int64  `json:"quantity,omitempty"`
	Available       int64  `json:"available,omitempty"`
	MfgDate         string `json:"mfgDate,omitempty"`
	ExpDate         string `json:"expDate,omitempty"`
}

// QueryBatchInventoryList queries batch-level inventory records.
func (s *InventoryService) QueryBatchInventoryList(receiveTimeFrom, receiveTimeTo string, page, pageSize int) ([]BatchInventory, int, error) {
	params := map[string]any{"receiveTimeFrom": receiveTimeFrom, "receiveTimeTo": receiveTimeTo, "page": page, "pageSize": pageSize}
	biz, _ := json.Marshal(params)
	var list []BatchInventory
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_BATCH_INVENTORY_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// TransferStorageLocationParams holds params for transferring goods between locations.
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

// SBSInventory represents Shopee SBS inventory.
type SBSInventory struct {
	Sku          string `json:"sku,omitempty"`
	SkuName      string `json:"skuName,omitempty"`
	WarehouseID  int64  `json:"warehouseId,omitempty"`
	TotalStock   int64  `json:"totalStock,omitempty"`
	Available    int64  `json:"available,omitempty"`
	Reserved     int64  `json:"reserved,omitempty"`
}

// QuerySBSInventoryList queries Shopee SBS inventory.
func (s *InventoryService) QuerySBSInventoryList(warehouseID int64, page, pageSize int) ([]SBSInventory, int, error) {
	params := map[string]any{"warehouseId": warehouseID, "page": page, "pageSize": pageSize}
	biz, _ := json.Marshal(params)
	var list []SBSInventory
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_SBS_INVENTORY_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// SBSWarehouse represents a Shopee SBS warehouse.
type SBSWarehouse struct {
	WarehouseID   int64  `json:"warehouseId,omitempty"`
	WarehouseName string `json:"warehouseName,omitempty"`
	Region        string `json:"region,omitempty"`
}

// QuerySBSWarehouseList queries Shopee SBS warehouse list.
func (s *InventoryService) QuerySBSWarehouseList() ([]SBSWarehouse, error) {
	params := map[string]any{}
	biz, _ := json.Marshal(params)
	var list []SBSWarehouse
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_SBS_WAREHOUSE_LIST", string(biz), w); err != nil {
		return nil, err
	}
	if w.HasError() {
		return nil, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, nil
}
