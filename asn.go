package qianyi

import "context"

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
	ApiCustom      string  `json:"apiCustom,omitempty"`
}

// CreateAsnParams holds the parameters for creating an inbound order.
type CreateAsnParams struct {
	WarehouseName            string            `json:"warehouseName"`
	AsnSkuVOList             []AsnSku          `json:"asnSkuVOList"`
	CustomNumber             string            `json:"customNumber,omitempty"`
	TrackNumber              string            `json:"trackNumber,omitempty"`
	Remark                   string            `json:"remark,omitempty"`
	PurchasePriceCurrency    string            `json:"purchasePriceCurrency"`
	FirstLegPriceCurrency    string            `json:"firstLegPriceCurrency,omitempty"`
	TransferPriceCurrency    string            `json:"transferPriceCurrency,omitempty"`
	SendWarehouseFlag        string            `json:"sendWarehouseFlag,omitempty"`
	PreArriveTime            string            `json:"preArriveTime,omitempty"`
	ShippingType             string            `json:"shippingType,omitempty"`
	ContainerModel           string            `json:"containerModel,omitempty"`
	PackageType              string            `json:"packageType,omitempty"`
	BoxCount                 int               `json:"boxCount,omitempty"`
	CustomerType             string            `json:"customerType,omitempty"`
	IsSpecifyBatch           bool              `json:"isSpecifyBatch,omitempty"`
	MergeDuplicateSkuLines   bool              `json:"mergeDuplicateSkuLines,omitempty"`
	AsnCustomFieldValueList  []CustomFieldValue `json:"asnCustomFieldValueVOList,omitempty"`
}

// Create creates a new inbound order (ASN) in QERP.
func (s *AsnService) Create(ctx context.Context, params *CreateAsnParams) error {
	return doAction(ctx, s.client, ServiceTypeCreateAsnOrder, params)
}

// AsnQueryParams holds parameters for querying inbound orders.
type AsnQueryParams struct {
	Page            int     `json:"page"`
	PageSize        int     `json:"pageSize"`
	WarehouseName   string  `json:"warehouseName,omitempty"`
	Type            string  `json:"type,omitempty"`
	Status          string  `json:"status,omitempty"`
	SkuKeyWord      string  `json:"skuKeyWord,omitempty"`
	Number          string  `json:"number,omitempty"`
	TrackNumber     string  `json:"trackNumber,omitempty"`
	TimeType        string  `json:"timeType,omitempty"`
	TimeFrom        string  `json:"timeFrom,omitempty"`
	TimeEnd         string  `json:"timeEnd,omitempty"`
	ReturnBatchInfo bool    `json:"returnBatchInfo,omitempty"`
	Tag             *AsnTag `json:"tag,omitempty"`
}

// QueryList retrieves inbound orders with optional filters.
func (s *AsnService) QueryList(ctx context.Context, params *AsnQueryParams) ([]AsnOrder, int, error) {
	return doList[AsnOrder](ctx, s.client, ServiceTypeQueryAsnList, params)
}

// Cancel cancels an inbound order by ASN number.
func (s *AsnService) Cancel(ctx context.Context, asnNumber string) error {
	return doAction(ctx, s.client, ServiceTypeCancelAsnOrder, map[string]any{"asnNumber": asnNumber})
}

// Delete deletes an inbound order by ASN number.
func (s *AsnService) Delete(ctx context.Context, asnNumber string) error {
	return doAction(ctx, s.client, ServiceTypeDeleteAsnOrder, map[string]any{"asnNumber": asnNumber})
}

// PushAsnParams holds the parameters for pushing inbound order status.
type PushAsnParams struct {
	AsnNumber    string       `json:"asnNumber"`
	Status       string       `json:"status"`
	TrackNumber  string       `json:"trackNumber,omitempty"`
	FinishedTime int64        `json:"finishedTime,omitempty"`
	CustomNumber string       `json:"customNumber,omitempty"`
	SkuList      []PushAsnSku `json:"skuList"`
}

// PushAsnSku represents a SKU receipt data in push notification.
type PushAsnSku struct {
	Sku      string `json:"sku"`
	Quantity int64  `json:"quantity"`
	BatchNo  string `json:"batchNo,omitempty"`
	MfgDate  string `json:"mfgDate,omitempty"`
	ExpDate  string `json:"expDate,omitempty"`
}

// PushOrder pushes inbound order receipt status notification.
func (s *AsnService) PushOrder(ctx context.Context, params *PushAsnParams) error {
	return doAction(ctx, s.client, ServiceTypePushAsnOrder, params)
}

// AsnBatchRecord represents a batch record from ASN batch list query.
type AsnBatchRecord struct {
	ReceiveTimeFrom string `json:"receiveTimeFrom,omitempty"`
	ReceiveTimeTo   string `json:"receiveTimeTo,omitempty"`
	Sku             string `json:"sku,omitempty"`
	SkuName         string `json:"skuName,omitempty"`
	Title           string `json:"title,omitempty"`
	WarehouseName   string `json:"warehouseName,omitempty"`
	BatchNumber     string `json:"batchNumber,omitempty"`
	Quantity        int64  `json:"quantity,omitempty"`
	Available       int64  `json:"available,omitempty"`
	MfgDate         string `json:"mfgDate,omitempty"`
	ExpDate         string `json:"expDate,omitempty"`
	OriginCountry   string `json:"originCountry,omitempty"`
}

// QueryBatchList queries inventory batch records for inbound orders.
func (s *AsnService) QueryBatchList(ctx context.Context, receiveTimeFrom, receiveTimeTo string, page, pageSize int) ([]AsnBatchRecord, int, error) {
	params := map[string]any{
		"receiveTimeFrom": receiveTimeFrom,
		"receiveTimeTo":   receiveTimeTo,
		"page":            page,
		"pageSize":        pageSize,
	}
	return doList[AsnBatchRecord](ctx, s.client, ServiceTypeQueryAsnBatchList, params)
}
