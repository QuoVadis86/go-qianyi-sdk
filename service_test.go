package qianyi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

// bizContentJSON builds a bizContent JSON string from the given result JSON.
// The QERP API returns bizContent as a JSON-encoded string inside the outer envelope.
func bizContentJSON(total int, resultJSON string) string {
	escaped, _ := json.Marshal(`{"state":"success","total":` + fmt.Sprintf("%d", total) + `,"result":` + resultJSON + `}`)
	return string(escaped)
}

// wrapResponse wraps a bizContent result JSON into a full API response envelope.
func wrapResponse(resultJSON string) string {
	return `{"state":"success","errorCode":"","errorMsg":"","bizContent":` +
		bizContentJSON(1, resultJSON) + `,"requestId":"req-1"}`
}

func wrapListResponse(resultJSON string) string {
	return `{"state":"success","errorCode":"","errorMsg":"","bizContent":` +
		bizContentJSON(2, resultJSON) + `,"requestId":"req-list"}`
}

func errorResponse() string {
	return `{"state":"success","errorCode":"BIZ_ERR","errorMsg":"business error","bizContent":"{}","requestId":"req-err"}`
}

func mockOK(body string) *mockHTTPClient {
	return &mockHTTPClient{responseBody: body, statusCode: http.StatusOK}
}

func svcClient(mock *mockHTTPClient) *Client {
	return NewClient("app", "secret", WithHTTPClient(mock))
}

// ---------------------------------------------------------------------------
// ShopService
// ---------------------------------------------------------------------------

func TestShopService_QueryList(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"shopId":1,"name":"S1","platform":"shopee"},{"shopId":2,"name":"S2","platform":"lazada"}]`)))
	svc := NewShopService(c)

	shops, total, err := svc.QueryList(context.Background(), 1, 10, "", "", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 {
		t.Fatalf("expected total=2, got %d", total)
	}
	if len(shops) != 2 || shops[0].ShopID != 1 || shops[0].Platform != "shopee" {
		t.Fatalf("unexpected shops: %+v", shops)
	}
}

func TestShopService_QueryList_Error(t *testing.T) {
	c := svcClient(mockOK(errorResponse()))
	svc := NewShopService(c)

	_, _, err := svc.QueryList(context.Background(), 1, 10, "", "", "", "")
	if err == nil {
		t.Fatal("expected error")
	}
	assertAPIError(t, err, "BIZ_ERR")
}

// ---------------------------------------------------------------------------
// WarehouseService
// ---------------------------------------------------------------------------

func TestWarehouseService_QueryList(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"id":1,"name":"WH1","kind":"standard","country":"CN"},{"id":2,"name":"WH2","kind":"overseas","country":"US"}]`)))
	svc := NewWarehouseService(c)

	whs, total, err := svc.QueryList(context.Background(), 1, 10, "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(whs) != 2 || whs[0].Name != "WH1" {
		t.Fatalf("unexpected warehouses: %+v", whs)
	}
}

// ---------------------------------------------------------------------------
// CustomerFieldService
// ---------------------------------------------------------------------------

func TestCustomerFieldService_Query(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`[{"tableName":"sales_order","columType":"text","columName":"cf"}]`)))
	svc := NewCustomerFieldService(c)

	fields, err := svc.Query(context.Background(), &CustomerFieldQueryParams{TableName: "sales_order"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(fields) != 1 || fields[0].TableName != "sales_order" {
		t.Fatalf("unexpected fields: %+v", fields)
	}
}

// ---------------------------------------------------------------------------
// SkuService
// ---------------------------------------------------------------------------

func TestSkuService_QueryList(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"sku":"SKU1","title":"T1","type":"simple"},{"sku":"SKU2","title":"T2","type":"simple"}]`)))
	svc := NewSkuService(c)

	skus, total, err := svc.QueryList(context.Background(), 1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(skus) != 2 || skus[0].Sku != "SKU1" {
		t.Fatalf("unexpected skus: %+v", skus)
	}
}

func TestSkuService_QueryList_WithFilter(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"sku":"SKU1","title":"T1","type":"simple"}]`)))
	svc := NewSkuService(c)

	skus, total, err := svc.QueryList(context.Background(), 1, 10, SkuFilterBySKU([]string{"SKU1"}), SkuFilterByTitle([]string{"T1"}))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(skus) != 1 {
		t.Fatalf("unexpected result: %+v", skus)
	}
}

func TestSkuService_Create(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewSkuService(c)

	err := svc.Create(context.Background(), &Sku{Sku: "NEW", Title: "New SKU"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSkuService_Update(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewSkuService(c)

	err := svc.Update(context.Background(), &Sku{Sku: "EXISTING", Title: "Updated"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSkuService_Enable(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewSkuService(c)

	err := svc.Enable(context.Background(), "SKU1", 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSkuService_QuerySysSKU(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"sku":"SYS1","title":"Sys SKU","type":"simple"}]`)))
	svc := NewSkuService(c)

	skus, total, err := svc.QuerySysSKU(context.Background(), 1, 10, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(skus) != 1 {
		t.Fatalf("unexpected result: %+v", skus)
	}
}

// ---------------------------------------------------------------------------
// OrderService
// ---------------------------------------------------------------------------

func TestOrderService_Create(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{"orderNumber":"ORD-001","status":"pending","shop":"Shop1","currency":"USD","platform":"shopee"}`)))
	svc := NewOrderService(c)

	order, err := svc.Create(context.Background(), &CreateOrderParams{
		Shop:              "Shop1",
		OnlineOrderNumber: "ON-001",
		Currency:          "USD",
		Buyer:             &Buyer{ReceiverName: "Test", Country: "US", Province: "CA", City: "LA", PostCode: "90001", Address1: "Addr"},
		SkuList:           []OrderSku{{Sku: "SKU1", Quantity: 1}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if order.OrderNumber != "ORD-001" || order.Status != "pending" {
		t.Fatalf("unexpected order: %+v", order)
	}
}

func TestOrderService_Cancel(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewOrderService(c)

	err := svc.Cancel(context.Background(), "ON-001", "Shop1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestOrderService_QueryList(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"orderNumber":"ORD-001","status":"pending","shop":"Shop1","currency":"USD","platform":"shopee"}]`)))
	svc := NewOrderService(c)

	orders, total, err := svc.QueryList(context.Background(), &OrderQueryParams{Page: 1, PageSize: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(orders) != 1 || orders[0].OrderNumber != "ORD-001" {
		t.Fatalf("unexpected orders: %+v", orders)
	}
}

func TestOrderService_QueryNumberList(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`["ORD-001","ORD-002"]`)))
	svc := NewOrderService(c)

	nums, total, err := svc.QueryNumberList(context.Background(), "", "", "", "", 1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(nums) != 2 || nums[0] != "ORD-001" {
		t.Fatalf("unexpected numbers: %+v", nums)
	}
}

func TestOrderService_QueryShippingInfo(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{"carrier":"DHL","trackingNumber":"TN-001"}`)))
	svc := NewOrderService(c)

	raw, err := svc.QueryShippingInfo(context.Background(), "ORD-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var info map[string]any
	if err := json.Unmarshal(raw, &info); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if info["carrier"] != "DHL" {
		t.Fatalf("unexpected shipping info: %+v", info)
	}
}

func TestOrderService_Audit(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewOrderService(c)

	err := svc.Audit(context.Background(), &AuditParams{OrderNumber: "ORD-001", Shop: "Shop1", OnlineOrderNumber: "ON-001"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestOrderService_CreateWaveOrder(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewOrderService(c)

	err := svc.CreateWaveOrder(context.Background(), &CreateWaveOrderParams{OrderNumberList: []string{"ORD-001"}, Shop: "Shop1"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestOrderService_SendToWms(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewOrderService(c)

	err := svc.SendToWms(context.Background(), &SendToWmsParams{OrderNumber: "ORD-001", Shop: "Shop1"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestOrderService_QueryOriginalOrder(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{"orderNumber":"ORD-ORIG","status":"pending","shop":"Shop1","currency":"USD","platform":"shopee"}`)))
	svc := NewOrderService(c)

	order, err := svc.QueryOriginalOrder(context.Background(), "Shop1", "ON-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if order.OrderNumber != "ORD-ORIG" {
		t.Fatalf("unexpected order: %+v", order)
	}
}

func TestOrderService_QueryPickupStatus(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{"status":"picked","carrier":"DHL"}`)))
	svc := NewOrderService(c)

	raw, err := svc.QueryPickupStatus(context.Background(), "Shop1", "ON-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var status map[string]any
	if err := json.Unmarshal(raw, &status); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if status["status"] != "picked" {
		t.Fatalf("unexpected pickup status: %+v", status)
	}
}

func TestOrderService_QueryOrderDocument(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{"documentUrl":"https://example.com/doc","type":"invoice"}`)))
	svc := NewOrderService(c)

	raw, err := svc.QueryOrderDocument(context.Background(), "ORD-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var doc map[string]any
	if err := json.Unmarshal(raw, &doc); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if doc["documentUrl"] != "https://example.com/doc" {
		t.Fatalf("unexpected document: %+v", doc)
	}
}

func TestOrderService_SubscribeOrder(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`[{"orderNumber":"ORD-001","errorMessage":""},{"orderNumber":"ORD-002","errorMessage":"already subscribed"}]`)))
	svc := NewOrderService(c)

	results, err := svc.SubscribeOrder(context.Background(), "sales", []string{"ORD-001", "ORD-002"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 2 || results[0].OrderNumber != "ORD-001" {
		t.Fatalf("unexpected results: %+v", results)
	}
}

// ---------------------------------------------------------------------------
// RefundService
// ---------------------------------------------------------------------------

func TestRefundService_Create(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{"returnNumber":"RET-001","status":"pending","shop":"Shop1","warehouse":"WH1","type":"refund"}`)))
	svc := NewRefundService(c)

	ret, err := svc.Create(context.Background(), &CreateRefundParams{
		Warehouse:     "WH1",
		Shop:          "Shop1",
		ReturnSkuList: []ReturnSku{{Sku: "SKU1", Quantity: 1}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ret.ReturnNumber != "RET-001" || ret.Status != "pending" {
		t.Fatalf("unexpected return: %+v", ret)
	}
}

func TestRefundService_Cancel(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewRefundService(c)

	err := svc.Cancel(context.Background(), "RET-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRefundService_QueryList(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"returnNumber":"RET-001","status":"pending"}]`)))
	svc := NewRefundService(c)

	list, total, err := svc.QueryList(context.Background(), &RefundQueryParams{Page: 1, PageSize: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 || list[0].ReturnNumber != "RET-001" {
		t.Fatalf("unexpected list: %+v", list)
	}
}

func TestRefundService_PushReturnInfo(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewRefundService(c)

	err := svc.PushReturnInfo(context.Background(), &PushReturnOrderInfoParams{
		ReturnNumber:      "RET-001",
		OrderNumber:       "ORD-001",
		OnlineOrderNumber: "ON-001",
		Status:            "received",
		ReturnSkuList:     []ReturnSku{{Sku: "SKU1", Quantity: 1}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// ---------------------------------------------------------------------------
// InventoryService
// ---------------------------------------------------------------------------

func TestInventoryService_QueryListV1(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"sku":"SKU1","warehouse":"WH1","total":100,"available":80,"allocated":20}]`)))
	svc := NewInventoryService(c)

	list, total, err := svc.QueryListV1(context.Background(), &InventoryQueryV1Params{Page: 1, PageSize: 10, Warehouse: "WH1"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 || list[0].Sku != "SKU1" {
		t.Fatalf("unexpected inventory: %+v", list)
	}
}

func TestInventoryService_QueryListV2(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"sku":"SKU1","warehouse":"WH1","total":100,"available":80,"allocated":20}]`)))
	svc := NewInventoryService(c)

	list, total, err := svc.QueryListV2(context.Background(), &InventoryQueryV2Params{Page: 1, PageSize: 10, Warehouse: "WH1"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 {
		t.Fatalf("unexpected: %+v", list)
	}
}

func TestInventoryService_QueryLogList(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"sku":"SKU1","warehouseName":"WH1","operateType":"in","operator":"admin","billNumber":"B-001","bizType":"purchase","uniqueId":1}]`)))
	svc := NewInventoryService(c)

	list, total, err := svc.QueryLogList(context.Background(), &InventoryLogQueryParams{OperateTimeFrom: "2025-01-01", OperateTimeTo: "2025-01-31", Page: 1, PageSize: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 || list[0].Sku != "SKU1" {
		t.Fatalf("unexpected: %+v", list)
	}
}

func TestInventoryService_QueryAssemblyList(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"assemblyNumber":"AS-001","warehouseName":"WH1","status":"done"}]`)))
	svc := NewInventoryService(c)

	list, total, err := svc.QueryAssemblyList(context.Background(), &AssemblyQueryParams{Page: 1, PageSize: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 {
		t.Fatalf("unexpected: %+v", list)
	}
}

func TestInventoryService_CreateTransferOrder(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewInventoryService(c)

	err := svc.CreateTransferOrder(context.Background(), &CreateTransferOrderParams{WarehouseFrom: "WH1", WarehouseTo: "WH2", Sku: "SKU1", Quantity: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestInventoryService_QueryTransferOrderList(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"number":"TO-001","status":"done"}]`)))
	svc := NewInventoryService(c)

	list, total, err := svc.QueryTransferOrderList(context.Background(), 1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 {
		t.Fatalf("unexpected: %+v", list)
	}
}

func TestInventoryService_QuerySplitOrderList(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"orderNumber":"ORD-001","splitOrderNumber":"SPLIT-001","status":"done"}]`)))
	svc := NewInventoryService(c)

	list, total, err := svc.QuerySplitOrderList(context.Background(), 1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 {
		t.Fatalf("unexpected: %+v", list)
	}
}

func TestInventoryService_QueryStorageLocInventory(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`[{"sku":"SKU1","warehouse":"WH1","storageLocation":"A-01","quantity":50,"available":50}]`)))
	svc := NewInventoryService(c)

	list, err := svc.QueryStorageLocInventory(context.Background(), "WH1", "A-01", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(list) != 1 || list[0].Sku != "SKU1" {
		t.Fatalf("unexpected: %+v", list)
	}
}

func TestInventoryService_QueryBatchInventoryList(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"sku":"SKU1","warehouseName":"WH1","batchNumber":"B-001","quantity":50,"available":50}]`)))
	svc := NewInventoryService(c)

	list, total, err := svc.QueryBatchInventoryList(context.Background(), &BatchInventoryQueryParams{
		ReceiveTimeFrom: "2025-01-01", ReceiveTimeTo: "2025-01-31", Page: 1, PageSize: 10,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 {
		t.Fatalf("unexpected: %+v", list)
	}
}

func TestInventoryService_TransferStorageLocation(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewInventoryService(c)

	err := svc.TransferStorageLocation(context.Background(), &TransferStorageLocationParams{
		FromLocation: "A-01", ToLocation: "A-02",
		SkuList: []struct {
			Sku      string `json:"sku"`
			Quantity int    `json:"quantity"`
		}{{Sku: "SKU1", Quantity: 10}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestInventoryService_QuerySBSInventoryList(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"sku":"SKU1","warehouseId":1,"totalStock":100,"available":80,"reserved":20}]`)))
	svc := NewInventoryService(c)

	list, total, err := svc.QuerySBSInventoryList(context.Background(), 1, 1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 || list[0].Sku != "SKU1" {
		t.Fatalf("unexpected: %+v", list)
	}
}

func TestInventoryService_QuerySBSWarehouseList(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`[{"warehouseId":1,"warehouseName":"SBS-WH","region":"SG"}]`)))
	svc := NewInventoryService(c)

	list, err := svc.QuerySBSWarehouseList(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(list) != 1 || list[0].WarehouseID != 1 {
		t.Fatalf("unexpected: %+v", list)
	}
}

// ---------------------------------------------------------------------------
// AsnService
// ---------------------------------------------------------------------------

func TestAsnService_Create(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewAsnService(c)

	err := svc.Create(context.Background(), &CreateAsnParams{
		WarehouseName:         "WH1",
		PurchasePriceCurrency: "USD",
		AsnSkuVOList:          []AsnSku{{Sku: "SKU1", PurchasePrice: 10, ExpectQuantity: 100}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAsnService_QueryList(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"asnNumber":"ASN-001","warehouseName":"WH1","type":"inbound","status":"pending","createTime":"2025-01-01"}]`)))
	svc := NewAsnService(c)

	list, total, err := svc.QueryList(context.Background(), &AsnQueryParams{Page: 1, PageSize: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 || list[0].AsnNumber != "ASN-001" {
		t.Fatalf("unexpected: %+v", list)
	}
}

func TestAsnService_Cancel(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewAsnService(c)

	err := svc.Cancel(context.Background(), "ASN-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAsnService_Delete(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewAsnService(c)

	err := svc.Delete(context.Background(), "ASN-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAsnService_PushOrder(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewAsnService(c)

	err := svc.PushOrder(context.Background(), &PushAsnParams{
		AsnNumber: "ASN-001", Status: "received",
		SkuList: []PushAsnSku{{Sku: "SKU1", Quantity: 100}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAsnService_QueryBatchList(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"sku":"SKU1","warehouseName":"WH1","batchNumber":"B-001","quantity":50,"available":50}]`)))
	svc := NewAsnService(c)

	list, total, err := svc.QueryBatchList(context.Background(), "2025-01-01", "2025-01-31", 1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 {
		t.Fatalf("unexpected: %+v", list)
	}
}

// ---------------------------------------------------------------------------
// OdoService
// ---------------------------------------------------------------------------

func TestOdoService_QueryList(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"number":"ODO-001","warehouseName":"WH1","type":"outbound","status":"pending","createTime":"2025-01-01"}]`)))
	svc := NewOdoService(c)

	list, total, err := svc.QueryList(context.Background(), &OdoQueryParams{Page: 1, PageSize: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 {
		t.Fatalf("unexpected: %+v", list)
	}
}

func TestOdoService_Create(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewOdoService(c)

	err := svc.Create(context.Background(), &CreateOdoParams{
		WarehouseName: "WH1",
		CustomNumber:  "EXT-001",
		OdoSkuVOList:  []OdoSkuCreateItem{{Sku: "SKU1", Quantity: 10}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestOdoService_Cancel(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewOdoService(c)

	err := svc.Cancel(context.Background(), "EXT-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestOdoService_PushOrder(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewOdoService(c)

	err := svc.PushOrder(context.Background(), &PushOdoParams{
		Number: "ODO-001", CustomNumber: "EXT-001", TrackNumber: "TN-001",
		WarehouseName: "WH1", Type: "outbound", Status: "shipped",
		SkuList: []PushOdoSku{{Sku: "SKU1", Title: "T1", Quantity: 10}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestOdoService_QuerySalesList(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"number":"ODO-001","warehouseName":"WH1","type":"outbound","status":"pending","createTime":"2025-01-01"}]`)))
	svc := NewOdoService(c)

	list, total, err := svc.QuerySalesList(context.Background(), "2025-01-01", "2025-01-31", 1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 {
		t.Fatalf("unexpected: %+v", list)
	}
}

// ---------------------------------------------------------------------------
// AdjustService
// ---------------------------------------------------------------------------

func TestAdjustService_QueryList(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"adjustmentNumber":"ADJ-001","source":"manual","warehouseName":"WH1","createTime":"2025-01-01"}]`)))
	svc := NewAdjustService(c)

	list, total, err := svc.QueryList(context.Background(), &AdjustQueryParams{Page: 1, PageSize: 10, Source: "manual"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 {
		t.Fatalf("unexpected: %+v", list)
	}
}

func TestAdjustService_Create(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewAdjustService(c)

	err := svc.Create(context.Background(), &CreateAdjustParams{
		WarehouseName: "WH1",
		AdjustSkuList: []AdjustSkuInput{{Sku: "SKU1", AdjustmentQtyStr: 10}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// ---------------------------------------------------------------------------
// PurchaseService
// ---------------------------------------------------------------------------

func TestPurchaseService_QueryList(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"purchaseNumber":"PO-001","warehouseName":"WH1","purchaseType":"local","status":"pending"}]`)))
	svc := NewPurchaseService(c)

	list, total, err := svc.QueryList(context.Background(), &PurchaseQueryParams{Page: 1, PageSize: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 {
		t.Fatalf("unexpected: %+v", list)
	}
}

func TestPurchaseService_Create(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewPurchaseService(c)

	err := svc.Create(context.Background(), &CreatePurchaseParams{
		PurchaseType:      "local",
		WarehouseName:     "WH1",
		PurchaserName:     "buyer",
		PurchaseDate:      "2025-01-01",
		PurchasePriceUnit: "USD",
		TransportParty:    "seller",
		TransportMode:     "sea",
		SupplierName:      "Supplier1",
		PaymentType:       "credit",
		SettlementType:    "monthly",
		SkuList:           []PurchaseSkuInput{{Sku: "SKU1", PurchasePrice: 10, PurchaseQuantity: 100, TaxRate: 0.05}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPurchaseService_Update(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewPurchaseService(c)

	err := svc.Update(context.Background(), &CreatePurchaseParams{
		PurchaseNumber:    "PO-001",
		PurchaseType:      "local",
		WarehouseName:     "WH1",
		PurchaserName:     "buyer",
		PurchaseDate:      "2025-01-01",
		PurchasePriceUnit: "USD",
		TransportParty:    "seller",
		TransportMode:     "sea",
		SupplierName:      "Supplier1",
		PaymentType:       "credit",
		SettlementType:    "monthly",
		SkuList:           []PurchaseSkuInput{{Sku: "SKU1", PurchasePrice: 12, PurchaseQuantity: 150, TaxRate: 0.05}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// ---------------------------------------------------------------------------
// LogisticsService
// ---------------------------------------------------------------------------

func TestLogisticsService_QueryFirstLegList(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"firstLegNumber":"FL-001","asnNumber":"ASN-001","customNumber":"EXT-001","status":"pending","createTime":"2025-01-01","updateTime":"2025-01-01","warehouseName":"WH1","destWarehouseName":"WH2"}]`)))
	svc := NewLogisticsService(c)

	list, total, err := svc.QueryFirstLegList(context.Background(), &FirstLegQueryParams{Page: 1, PageSize: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 {
		t.Fatalf("unexpected: %+v", list)
	}
}

func TestLogisticsService_CreateFirstLeg(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewLogisticsService(c)

	err := svc.CreateFirstLeg(context.Background(), &CreateFirstLegParams{
		DestWarehouseType: "FBA",
		SkuDetailList: []FirstLegSkuDetail{
			{
				LineID: 1, WarehouseName: "WH1", DestWarehouseName: "FBA-US",
				LogisticsName: "DHL", Sku: "SKU1", PreExpectedQuantity: 100,
			},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestLogisticsService_QueryFirstLegLogistics(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`[{"id":1,"logisticsName":"DHL"},{"id":2,"logisticsName":"FedEx"}]`)))
	svc := NewLogisticsService(c)

	list, err := svc.QueryFirstLegLogistics(context.Background(), "WH1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(list) != 2 || list[0].LogisticsName != "DHL" {
		t.Fatalf("unexpected: %+v", list)
	}
}

func TestLogisticsService_QueryFirstLegTracking(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"orderNumber":"ORD-001","onlineOrderId":"ON-001","trackingNumber":"TN-001","carrier":"DHL","status":"in_transit"}]`)))
	svc := NewLogisticsService(c)

	list, total, err := svc.QueryFirstLegTracking(context.Background(), &TrackingQueryParams{
		UpdateTimeFrom: "2025-01-01", UpdateTimeTo: "2025-01-31", Page: 1, PageSize: 10,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 {
		t.Fatalf("unexpected: %+v", list)
	}
}

func TestLogisticsService_WithdrawFirstLeg(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewLogisticsService(c)

	err := svc.WithdrawFirstLeg(context.Background(), &WithdrawFirstLegParams{FirstLegNumber: "FL-001"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestLogisticsService_PushTrackingPackage(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewLogisticsService(c)

	err := svc.PushTrackingPackage(context.Background(), &PushTrackingParams{
		OrderNumber: "ORD-001", OnlineOrderID: "ON-001",
		TrackingNumber: "TN-001", Carrier: "DHL", Status: "delivered",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// ---------------------------------------------------------------------------
// ReportService
// ---------------------------------------------------------------------------

func TestReportService_QueryShopeeTransaction(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"shopId":1,"onlineShopId":"shop1","shopName":"S1","ordersn":"ORD-001","payoutTime":1700000000,"payoutTimeFormatted":"2025-01-01","currency":"SGD"}]`)))
	svc := NewReportService(c)

	list, total, err := svc.QueryShopeeTransaction(context.Background(), &ShopeeReportQuery{Page: 1, PageSize: 10, Type: 1})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 || list[0].ShopName != "S1" {
		t.Fatalf("unexpected: %+v", list)
	}
}

func TestReportService_QueryLazadaTransaction(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"shopId":1,"onlineShopId":"shop1","shopName":"L1","transactionDate":"2025-01-01","transactionTimestamp":1700000000,"transactionNumber":"T-001","transactionType":"sale","orderNo":"ORD-001"}]`)))
	svc := NewReportService(c)

	list, total, err := svc.QueryLazadaTransaction(context.Background(), &LazadaReportQuery{Page: 1, PageSize: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 {
		t.Fatalf("unexpected: %+v", list)
	}
}

func TestReportService_QueryTiktokTransaction(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"shopId":1,"shopName":"T1","currency":"USD","orderId":"ORD-001"}]`)))
	svc := NewReportService(c)

	list, total, err := svc.QueryTiktokTransaction(context.Background(), &TiktokReportQuery{
		Page: 1, PageSize: 10, PayoutTimeFrom: "2025-01-01", PayoutTimeTo: "2025-01-31",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 {
		t.Fatalf("unexpected: %+v", list)
	}
}

func TestReportService_QueryShopeePayout(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"shopId":1,"onlineShopId":"shop1","shopName":"S1","currency":"SGD","type":"payout"}]`)))
	svc := NewReportService(c)

	list, total, err := svc.QueryShopeePayout(context.Background(), &ShopeePayoutQuery{
		Page: 1, PageSize: 10, PayoutTimeFrom: "2025-01-01", PayoutTimeTo: "2025-01-31",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 {
		t.Fatalf("unexpected: %+v", list)
	}
}

func TestReportService_QueryLazadaAccountTransaction(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"customerId":1,"shopId":1,"onlineShopId":"shop1","siteCode":"SG","transactionNumber":"T-001","transactionTime":1700000000,"type":"payout"}]`)))
	svc := NewReportService(c)

	list, total, err := svc.QueryLazadaAccountTransaction(context.Background(), &LazadaAccountQuery{
		Page: 1, PageSize: 10, PayoutTimeFrom: "2025-01-01", PayoutTimeTo: "2025-01-31",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 {
		t.Fatalf("unexpected: %+v", list)
	}
}

func TestReportService_QueryTiktokV2Transaction(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"shopId":1,"shopName":"T1","currency":"USD","orderId":"ORD-001","orderType":"normal","orderStatus":"completed"}]`)))
	svc := NewReportService(c)

	list, total, err := svc.QueryTiktokV2Transaction(context.Background(), &TiktokV2ReportQuery{
		Page: 1, PageSize: 10, PayoutTimeFrom: "2025-01-01", PayoutTimeTo: "2025-01-31",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 {
		t.Fatalf("unexpected: %+v", list)
	}
}

func TestReportService_QueryTiktokPayout(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"shopId":1,"onlineShopId":"shop1","shopName":"T1","currency":"USD","payoutTime":1700000000,"payoutTimeFormatted":"2025-01-01","transactionId":"TXN-001","transactionType":"payout","amount":100.5,"balance":500,"status":"completed"}]`)))
	svc := NewReportService(c)

	list, total, err := svc.QueryTiktokPayout(context.Background(), &TiktokPayoutQuery{
		Page: 1, PageSize: 10, PayoutTimeFrom: "2025-01-01", PayoutTimeTo: "2025-01-31",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 {
		t.Fatalf("unexpected: %+v", list)
	}
}

func TestReportService_QueryInventoryDailyReport(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"sku":"SKU1","skuName":"Product A","warehouseName":"WH1","date":"2025-01-01","beginQuantity":100,"inQuantity":20,"outQuantity":10,"endQuantity":110}]`)))
	svc := NewReportService(c)

	list, total, err := svc.QueryInventoryDailyReport(context.Background(), &InventoryDailyReportQuery{
		Page: 1, PageSize: 10, DateFrom: "2025-01-01", DateTo: "2025-01-31",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 || len(list) != 1 || list[0].Sku != "SKU1" {
		t.Fatalf("unexpected: %+v", list)
	}
}

// ---------------------------------------------------------------------------
// all-service error-path test
// ---------------------------------------------------------------------------

func TestAllServices_BizError(t *testing.T) {
	tests := []struct {
		name string
		run  func(t *testing.T, c *Client)
	}{
		{"Shop.QueryList", func(t *testing.T, c *Client) {
			_, _, err := NewShopService(c).QueryList(ctx, 1, 10, "", "", "", "")
			assertAPIError(t, err, "BIZ_ERR")
		}},
		{"Sku.QueryList", func(t *testing.T, c *Client) {
			_, _, err := NewSkuService(c).QueryList(ctx, 1, 10)
			assertAPIError(t, err, "BIZ_ERR")
		}},
		{"Order.QueryList", func(t *testing.T, c *Client) {
			_, _, err := NewOrderService(c).QueryList(ctx, &OrderQueryParams{Page: 1, PageSize: 10})
			assertAPIError(t, err, "BIZ_ERR")
		}},
		{"Refund.QueryList", func(t *testing.T, c *Client) {
			_, _, err := NewRefundService(c).QueryList(ctx, &RefundQueryParams{Page: 1, PageSize: 10})
			assertAPIError(t, err, "BIZ_ERR")
		}},
		{"Inventory.QueryListV2", func(t *testing.T, c *Client) {
			_, _, err := NewInventoryService(c).QueryListV2(ctx, &InventoryQueryV2Params{Page: 1, PageSize: 10, Warehouse: "WH1"})
			assertAPIError(t, err, "BIZ_ERR")
		}},
		{"Asn.QueryList", func(t *testing.T, c *Client) {
			_, _, err := NewAsnService(c).QueryList(ctx, &AsnQueryParams{Page: 1, PageSize: 10})
			assertAPIError(t, err, "BIZ_ERR")
		}},
		{"Odo.QueryList", func(t *testing.T, c *Client) {
			_, _, err := NewOdoService(c).QueryList(ctx, &OdoQueryParams{Page: 1, PageSize: 10})
			assertAPIError(t, err, "BIZ_ERR")
		}},
		{"Adjust.QueryList", func(t *testing.T, c *Client) {
			_, _, err := NewAdjustService(c).QueryList(ctx, &AdjustQueryParams{Page: 1, PageSize: 10, Source: "manual"})
			assertAPIError(t, err, "BIZ_ERR")
		}},
		{"Purchase.QueryList", func(t *testing.T, c *Client) {
			_, _, err := NewPurchaseService(c).QueryList(ctx, &PurchaseQueryParams{Page: 1, PageSize: 10})
			assertAPIError(t, err, "BIZ_ERR")
		}},
		{"Logistics.QueryFirstLegList", func(t *testing.T, c *Client) {
			_, _, err := NewLogisticsService(c).QueryFirstLegList(ctx, &FirstLegQueryParams{Page: 1, PageSize: 10})
			assertAPIError(t, err, "BIZ_ERR")
		}},
		{"Report.QueryShopeeTransaction", func(t *testing.T, c *Client) {
			_, _, err := NewReportService(c).QueryShopeeTransaction(ctx, &ShopeeReportQuery{Page: 1, PageSize: 10, Type: 1})
			assertAPIError(t, err, "BIZ_ERR")
		}},
		{"CustomerField.Query", func(t *testing.T, c *Client) {
			_, err := NewCustomerFieldService(c).Query(ctx, &CustomerFieldQueryParams{TableName: "sales_order"})
			assertAPIError(t, err, "BIZ_ERR")
		}},
		{"Supplier.QueryList", func(t *testing.T, c *Client) {
			_, _, err := NewSupplierService(c).QueryList(ctx, &QuerySupplierParams{Page: 1, PageSize: 10})
			assertAPIError(t, err, "BIZ_ERR")
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := svcClient(mockOK(errorResponse()))
			tt.run(t, c)
		})
	}
}

// ---------------------------------------------------------------------------
// SupplierService
// ---------------------------------------------------------------------------

func TestSupplierService_QueryList(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"name":"SupplierA","enable":true,"category":"COMMON"}]`)))
	svc := NewSupplierService(c)
	result, total, err := svc.QueryList(ctx, &QuerySupplierParams{Page: 1, PageSize: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 {
		t.Fatalf("expected total=2, got %d", total)
	}
	if len(result) != 1 || result[0].Name != "SupplierA" {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestSupplierService_Create(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewSupplierService(c)
	err := svc.Create(ctx, &CreateSupplierParams{
		Name:                      "SupplierA",
		Category:                  "COMMON",
		PurchaserUserName:         "zhangsan",
		SettlementWay:             "DELIVERY_ON_CASH",
		PaymentWay:                "CASH",
		TransportParty:            "SUPPLIER",
		DefectiveProductResolution: "PURCHASER_UNDERTAKING",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSupplierService_QuerySkuList(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"sku":"SKU001","title":"Product A","skuSupplierList":[{"supplierName":"SupplierA","category":"COMMON","deliveryCycle":7}]}]`)))
	svc := NewSupplierService(c)
	result, total, err := svc.QuerySkuList(ctx, &QuerySupplierSkuParams{Page: 1, PageSize: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 {
		t.Fatalf("expected total=2, got %d", total)
	}
	if len(result) != 1 || result[0].Sku != "SKU001" {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestSupplierService_CreateSku(t *testing.T) {
	c := svcClient(mockOK(wrapResponse(`{}`)))
	svc := NewSupplierService(c)
	err := svc.CreateSku(ctx, &CreateSupplierSkuParams{
		Sku:                      "SKU001",
		SupplierName:             "SupplierA",
		DefaultPurchaserUserName: "zhangsan",
		DeliveryCycle:            7,
		PurchaseMethod:           "BULK",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSupplierService_QueryList_WithFilter(t *testing.T) {
	c := svcClient(mockOK(wrapListResponse(`[{"name":"SupplierA","enable":true}]`)))
	svc := NewSupplierService(c)
	enable := true
	result, total, err := svc.QueryList(ctx, &QuerySupplierParams{
		Page:     1,
		PageSize: 10,
		Name:     "SupplierA",
		Enable:   &enable,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 2 {
		t.Fatalf("expected total=2, got %d", total)
	}
	if len(result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result))
	}
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

var ctx = context.Background()

func assertAPIError(t *testing.T, err error, code string) {
	t.Helper()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	ae, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T: %v", err, err)
	}
	if ae.ErrorCode != code {
		t.Fatalf("expected error code %q, got %q", code, ae.ErrorCode)
	}
}

// ---------------------------------------------------------------------------
// Transport-level: 500 error and network failure
// ---------------------------------------------------------------------------

func TestShopService_QueryList_HTTP500(t *testing.T) {
	c := NewClient("app", "secret")
	c.HTTPClient = &mockHTTPClient{
		responseBody: `{"error":"internal"}`,
		statusCode:   http.StatusInternalServerError,
	}
	svc := NewShopService(c)
	_, _, err := svc.QueryList(context.Background(), 1, 10, "", "", "", "")
	if err == nil {
		t.Fatal("expected error for 500")
	}
}

func TestShopService_QueryList_NetworkError(t *testing.T) {
	c := NewClient("app", "secret")
	c.HTTPClient = httpClientFunc(func(*http.Request) (*http.Response, error) {
		return nil, io.ErrUnexpectedEOF
	})
	svc := NewShopService(c)
	_, _, err := svc.QueryList(context.Background(), 1, 10, "", "", "", "")
	if err == nil {
		t.Fatal("expected network error")
	}
}
