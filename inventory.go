package qianyi

import "context"

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
func (s *InventoryService) QueryListV1(ctx context.Context, params *InventoryQueryV1Params) ([]SkuInventory, int, error) {
	return doList[SkuInventory](ctx, s.client, "QUERY_SIMPLE_LIST_INVENTORY", params)
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
func (s *InventoryService) QueryListV2(ctx context.Context, params *InventoryQueryV2Params) ([]SkuInventory, int, error) {
	return doList[SkuInventory](ctx, s.client, "QUERY_SIMPLE_LIST_INVENTORY_V2", params)
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
func (s *InventoryService) QueryLogList(ctx context.Context, params *InventoryLogQueryParams) ([]InventoryLog, int, error) {
	return doList[InventoryLog](ctx, s.client, "QUERY_INVENTORY_LOG_LIST", params)
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
	AssemblyNumber string        `json:"assemblyNumber,omitempty"`
	WarehouseName  string        `json:"warehouseName,omitempty"`
	Status         string        `json:"status,omitempty"`
	AsnNumber      string        `json:"asnNumber,omitempty"`
	FinishTime     string        `json:"finishTime,omitempty"`
	CreateTime     string        `json:"createTime,omitempty"`
	AssemblyList   []AssemblySku `json:"assemblyList,omitempty"`
}

// AssemblySku represents a SKU line in an assembly order.
type AssemblySku struct {
	Sku      string `json:"sku,omitempty"`
	Title    string `json:"title,omitempty"`
	Quantity int64  `json:"quantity,omitempty"`
}

// QueryAssemblyList retrieves assembly orders.
func (s *InventoryService) QueryAssemblyList(ctx context.Context, params *AssemblyQueryParams) ([]AssemblyOrder, int, error) {
	return doList[AssemblyOrder](ctx, s.client, "QUERY_INVENTORY_ASSEMBLY_LIST", params)
}

// CreateTransferOrderParams holds params for creating a transfer order.
type CreateTransferOrderParams struct {
	WarehouseFrom string `json:"warehouseFrom"`
	WarehouseTo   string `json:"warehouseTo"`
	Sku           string `json:"sku"`
	Quantity      int    `json:"quantity"`
}

// CreateTransferOrder creates a transfer order between warehouses.
func (s *InventoryService) CreateTransferOrder(ctx context.Context, params *CreateTransferOrderParams) error {
	return doAction(ctx, s.client, "CREATE_TRANSFER_ORDER", params)
}

// TransferOrder represents a warehouse transfer order.
type TransferOrder struct {
	Number          string `json:"number,omitempty"`
	Status          string `json:"status,omitempty"`
	SourceWarehouse string `json:"sourceWarehouse,omitempty"`
	TargetWarehouse string `json:"targetWarehouse,omitempty"`
	Sku             string `json:"sku,omitempty"`
	Quantity        int64  `json:"quantity,omitempty"`
	CreateTime      string `json:"createTime,omitempty"`
}

// QueryTransferOrderList queries transfer orders.
func (s *InventoryService) QueryTransferOrderList(ctx context.Context, page, pageSize int) ([]TransferOrder, int, error) {
	params := map[string]any{"page": page, "pageSize": pageSize}
	return doList[TransferOrder](ctx, s.client, "QUERY_TRANSFER_ORDER_LIST", params)
}

// SplitOrder represents an order split record.
type SplitOrder struct {
	OrderNumber      string `json:"orderNumber,omitempty"`
	SplitOrderNumber string `json:"splitOrderNumber,omitempty"`
	Sku              string `json:"sku,omitempty"`
	Quantity         int64  `json:"quantity,omitempty"`
	Status           string `json:"status,omitempty"`
	CreateTime       string `json:"createTime,omitempty"`
}

// QuerySplitOrderList queries split orders.
func (s *InventoryService) QuerySplitOrderList(ctx context.Context, page, pageSize int) ([]SplitOrder, int, error) {
	params := map[string]any{"page": page, "pageSize": pageSize}
	return doList[SplitOrder](ctx, s.client, "QUERY_SPLIT_ORDER_LIST", params)
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
func (s *InventoryService) QueryStorageLocInventory(ctx context.Context, warehouse, storageLocation string, skuList []string) ([]StorageLocInventory, error) {
	params := map[string]any{"warehouse": warehouse, "storageLocation": storageLocation}
	if len(skuList) > 0 {
		params["skuList"] = skuList
	}
	return doListNoTotal[StorageLocInventory](ctx, s.client, "QUERY_STORAGE_LOC_INVENTORY", params)
}

// BatchInventoryQueryParams holds parameters for batch inventory queries.
type BatchInventoryQueryParams struct {
	ReceiveTimeFrom string `json:"receiveTimeFrom"`
	ReceiveTimeTo   string `json:"receiveTimeTo"`
	Page            int    `json:"page"`
	PageSize        int    `json:"pageSize"`
	NumberParam     string `json:"numberParam,omitempty"`
	SkuParam        string `json:"skuParam,omitempty"`
	WarehouseName   string `json:"warehouseName,omitempty"`
	AsnType         string `json:"asnType,omitempty"`
}

// BatchInventory represents a batch inventory record.
type BatchInventory struct {
	Sku           string `json:"sku,omitempty"`
	SkuName       string `json:"skuName,omitempty"`
	WarehouseName string `json:"warehouseName,omitempty"`
	BatchNumber   string `json:"batchNumber,omitempty"`
	Quantity      int64  `json:"quantity,omitempty"`
	Available     int64  `json:"available,omitempty"`
	MfgDate       string `json:"mfgDate,omitempty"`
	ExpDate       string `json:"expDate,omitempty"`
}

// QueryBatchInventoryList queries batch-level inventory records.
func (s *InventoryService) QueryBatchInventoryList(ctx context.Context, params *BatchInventoryQueryParams) ([]BatchInventory, int, error) {
	return doList[BatchInventory](ctx, s.client, "QUERY_BATCH_INVENTORY_LIST", params)
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
func (s *InventoryService) TransferStorageLocation(ctx context.Context, params *TransferStorageLocationParams) error {
	return doAction(ctx, s.client, "TRANSFER_STORAGE_LOCATION", params)
}

// SBSInventory represents Shopee SBS inventory.
type SBSInventory struct {
	Sku         string `json:"sku,omitempty"`
	SkuName     string `json:"skuName,omitempty"`
	WarehouseID int64  `json:"warehouseId,omitempty"`
	TotalStock  int64  `json:"totalStock,omitempty"`
	Available   int64  `json:"available,omitempty"`
	Reserved    int64  `json:"reserved,omitempty"`
}

// QuerySBSInventoryList queries Shopee SBS inventory.
func (s *InventoryService) QuerySBSInventoryList(ctx context.Context, warehouseID int64, page, pageSize int) ([]SBSInventory, int, error) {
	params := map[string]any{"warehouseId": warehouseID, "page": page, "pageSize": pageSize}
	return doList[SBSInventory](ctx, s.client, "QUERY_SBS_INVENTORY_LIST", params)
}

// SBSWarehouse represents a Shopee SBS warehouse.
type SBSWarehouse struct {
	WarehouseID   int64  `json:"warehouseId,omitempty"`
	WarehouseName string `json:"warehouseName,omitempty"`
	Region        string `json:"region,omitempty"`
}

// QuerySBSWarehouseList queries Shopee SBS warehouse list.
func (s *InventoryService) QuerySBSWarehouseList(ctx context.Context) ([]SBSWarehouse, error) {
	return doListNoTotal[SBSWarehouse](ctx, s.client, "QUERY_SBS_WAREHOUSE_LIST", map[string]any{})
}
