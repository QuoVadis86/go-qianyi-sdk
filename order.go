package qianyi

import (
	"context"
	"encoding/json"
)

// OrderService provides access to sales order API operations.
type OrderService struct {
	client *Client
}

// NewOrderService creates a new OrderService.
func NewOrderService(client *Client) *OrderService {
	return &OrderService{client: client}
}

// CreateOrderParams holds the parameters for creating a sales order.
type CreateOrderParams struct {
	Shop                   string            `json:"shop"`
	OnlineOrderNumber      string            `json:"onlineOrderNumber"`
	PaymentMethod          string            `json:"paymentMethod"`
	Currency               string            `json:"currency"`
	PayTime                string            `json:"payTime"`
	Buyer                  *Buyer            `json:"buyer"`
	SkuList                []OrderSku        `json:"skuList"`
	Freight                float64           `json:"freight,omitempty"`
	CodPayAmount           float64           `json:"codPayAmount,omitempty"`
	BuyerMessage           string            `json:"buyerMessage,omitempty"`
	SellerRemarks          string            `json:"sellerRemarks,omitempty"`
	LogisticsSelected      string            `json:"logisticsSelected,omitempty"`
	TrackingNumber         string            `json:"trackingNumber,omitempty"`
	IsSpecifyBatch         bool              `json:"isSpecifyBatch,omitempty"`
	ShippingLabel          string            `json:"shippingLabel,omitempty"`
	ImgType                string            `json:"imgType,omitempty"`
	DocumentFile           string            `json:"documentFile,omitempty"`
	DocumentType           string            `json:"documentType,omitempty"`
	DocumentName           string            `json:"documentName,omitempty"`
	CustomerType           string            `json:"customerType,omitempty"`
	OrderCustomFieldValues []CustomFieldValue `json:"orderCustomFieldValueVOList,omitempty"`
}

// Create creates a new sales order in QERP.
func (s *OrderService) Create(ctx context.Context, params *CreateOrderParams) (*Order, error) {
	return doSingle[Order](ctx, s.client, ServiceTypeCreateSalesOrder, params)
}

// Cancel cancels a sales order by online order number and shop name.
func (s *OrderService) Cancel(ctx context.Context, onlineOrderNumber, shop string) error {
	return doAction(ctx, s.client, ServiceTypeCloseSalesOrder, map[string]any{"onlineOrderNumber": onlineOrderNumber, "shop": shop})
}

// OrderQueryParams holds parameters for querying sales orders.
type OrderQueryParams struct {
	Page              int      `json:"page"`
	PageSize          int      `json:"pageSize"`
	Status            string   `json:"status,omitempty"`
	Shop              string   `json:"shop,omitempty"`
	OrderNumber       string   `json:"orderNumber,omitempty"`
	OnlineOrderNumber string   `json:"onlineOrderNumber,omitempty"`
	FromPayTime       string   `json:"fromPayTime,omitempty"`
	ToPayTime         string   `json:"toPayTime,omitempty"`
	UpdateTimeFrom    string   `json:"updateTimeFrom,omitempty"`
	UpdateTimeTo      string   `json:"updateTimeTo,omitempty"`
	ShippingTimeFrom  string   `json:"shippingTimeFrom,omitempty"`
	ShippingTimeTo    string   `json:"shippingTimeTo,omitempty"`
	ShopIDList        []int64  `json:"shopIdList,omitempty"`
	OrderByParam      string   `json:"orderByParam,omitempty"`
	OrderByOrder      string   `json:"orderByOrder,omitempty"`
	CombineSku        *bool    `json:"combineSku,omitempty"`
	ReturnSnList      *bool    `json:"returnSnList,omitempty"`
	ReturnGiftFlag    *bool    `json:"returnGiftFlag,omitempty"`
}

// QueryList retrieves a paginated list of sales orders with optional filters.
func (s *OrderService) QueryList(ctx context.Context, params *OrderQueryParams) ([]Order, int, error) {
	return doList[Order](ctx, s.client, ServiceTypeQuerySalesOrderList, params)
}

// QueryNumberList retrieves a paginated list of sales order numbers.
func (s *OrderService) QueryNumberList(ctx context.Context, status, shop, fromPayTime, toPayTime string, page, pageSize int) ([]string, int, error) {
	params := map[string]any{"page": page, "pageSize": pageSize}
	if status != "" {
		params["status"] = status
	}
	if shop != "" {
		params["shop"] = shop
	}
	if fromPayTime != "" {
		params["fromPayTime"] = fromPayTime
	}
	if toPayTime != "" {
		params["toPayTime"] = toPayTime
	}
	return doList[string](ctx, s.client, ServiceTypeQuerySalesOrderNumberList, params)
}

// QueryShippingInfo queries the shipping information for a sales order.
func (s *OrderService) QueryShippingInfo(ctx context.Context, orderNumber string) (json.RawMessage, error) {
	params := map[string]any{"orderNumber": orderNumber}
	biz, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	var raw json.RawMessage
	w := &ResponseWrapper{Result: &raw}
	if err := s.client.Do(ctx, ServiceTypeQuerySalesOrderShippingInfo, string(biz), w); err != nil {
		return nil, err
	}
	if w.HasError() {
		return nil, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return raw, nil
}

// AuditParams holds parameters for auditing a sales order.
type AuditParams struct {
	OrderNumber       string `json:"orderNumber"`
	Shop              string `json:"shop"`
	OnlineOrderNumber string `json:"onlineOrderNumber"`
}

// Audit audits a sales order.
func (s *OrderService) Audit(ctx context.Context, params *AuditParams) error {
	return doAction(ctx, s.client, ServiceTypeQuerySalesOrderAudit, params)
}

// CreateWaveOrderParams holds parameters for creating a wave order.
type CreateWaveOrderParams struct {
	OrderNumberList []string `json:"orderNumberList"`
	Shop            string   `json:"shop"`
}

// CreateWaveOrder creates a wave order for batch picking.
func (s *OrderService) CreateWaveOrder(ctx context.Context, params *CreateWaveOrderParams) error {
	return doAction(ctx, s.client, ServiceTypeCreateWaveOrder, params)
}

// SendToWmsParams holds parameters for sending an order to WMS.
type SendToWmsParams struct {
	OrderNumber       string `json:"orderNumber"`
	OnlineOrderNumber string `json:"onlineOrderNumber,omitempty"`
	Shop              string `json:"shop"`
}

// SendToWms sends a sales order to the warehouse management system.
func (s *OrderService) SendToWms(ctx context.Context, params *SendToWmsParams) error {
	return doAction(ctx, s.client, ServiceTypeSendSalesOrderToWms, params)
}

// QueryOriginalOrder queries the original sales order by shop and online order number.
func (s *OrderService) QueryOriginalOrder(ctx context.Context, shop, onlineOrderNumber string) (*Order, error) {
	params := map[string]any{"shop": shop, "onlineOrderNumber": onlineOrderNumber}
	return doSingle[Order](ctx, s.client, ServiceTypeQueryOriginalSalesOrder, params)
}

// QueryPickupStatus queries the pickup status of a sales order.
func (s *OrderService) QueryPickupStatus(ctx context.Context, shop, onlineOrderNumber string) (json.RawMessage, error) {
	params := map[string]any{"shop": shop, "onlineOrderNumber": onlineOrderNumber}
	biz, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	var raw json.RawMessage
	w := &ResponseWrapper{Result: &raw}
	if err := s.client.Do(ctx, ServiceTypeQuerySalesOrderPickupStatus, string(biz), w); err != nil {
		return nil, err
	}
	if w.HasError() {
		return nil, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return raw, nil
}

// QueryOrderDocument queries the document attached to a sales order.
func (s *OrderService) QueryOrderDocument(ctx context.Context, orderNumber string) (json.RawMessage, error) {
	params := map[string]any{"orderNumber": orderNumber}
	biz, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	var raw json.RawMessage
	w := &ResponseWrapper{Result: &raw}
	if err := s.client.Do(ctx, ServiceTypeQuerySalesOrderDocument, string(biz), w); err != nil {
		return nil, err
	}
	if w.HasError() {
		return nil, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return raw, nil
}

// SubscribeOrderItem represents an order to subscribe to.
type SubscribeOrderItem struct {
	OrderNumber string `json:"orderNumber"`
}

// SubscribeOrderParams holds parameters for subscribing to order status.
type SubscribeOrderParams struct {
	OrderType string              `json:"orderType"`
	OrderList []SubscribeOrderItem `json:"orderList"`
}

// SubscribeOrderResult represents a subscription result item.
type SubscribeOrderResult struct {
	OrderNumber  string `json:"orderNumber"`
	ErrorMessage string `json:"errorMessage"`
}

// SubscribeOrder subscribes to order status push notifications.
func (s *OrderService) SubscribeOrder(ctx context.Context, orderType string, orderNumbers []string) ([]SubscribeOrderResult, error) {
	list := make([]SubscribeOrderItem, len(orderNumbers))
	for i, n := range orderNumbers {
		list[i] = SubscribeOrderItem{OrderNumber: n}
	}
	params := SubscribeOrderParams{
		OrderType: orderType,
		OrderList: list,
	}
	result, err := doSingle[[]SubscribeOrderResult](ctx, s.client, ServiceTypeSubscribeOrder, params)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}
	return *result, nil
}
