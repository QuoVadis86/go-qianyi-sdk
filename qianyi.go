package qianyi

type SDK struct {
	Client        *Client
	Shop          *ShopService
	Sku           *SkuService
	Order         *OrderService
	Refund        *RefundService
	Warehouse     *WarehouseService
	Inventory     *InventoryService
	Asn           *AsnService
	Odo           *OdoService
	Adjust        *AdjustService
	Purchase      *PurchaseService
	Logistics     *LogisticsService
	Report        *ReportService
	CustomerField *CustomerFieldService
}

func NewSDK(appID, appSecret string, opts ...ClientOption) *SDK {
	c := NewClient(appID, appSecret, opts...)
	return &SDK{
		Client:        c,
		Shop:          NewShopService(c),
		Sku:           NewSkuService(c),
		Order:         NewOrderService(c),
		Refund:        NewRefundService(c),
		Warehouse:     NewWarehouseService(c),
		Inventory:     NewInventoryService(c),
		Asn:           NewAsnService(c),
		Odo:           NewOdoService(c),
		Adjust:        NewAdjustService(c),
		Purchase:      NewPurchaseService(c),
		Logistics:     NewLogisticsService(c),
		Report:        NewReportService(c),
		CustomerField: NewCustomerFieldService(c),
	}
}

func (s *SDK) TestEnv() {
	s.Client.BaseURL = "https://gerp-test1.800best.com"
}
