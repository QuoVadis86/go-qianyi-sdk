package qianyi

import "encoding/json"

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
	Shop                string            `json:"shop"`
	OnlineOrderNumber   string            `json:"onlineOrderNumber"`
	PaymentMethod       string            `json:"paymentMethod"`
	Currency            string            `json:"currency"`
	PayTime             string            `json:"payTime"`
	Buyer               *Buyer            `json:"buyer"`
	SkuList             []OrderSku        `json:"skuList"`
	Freight             float64           `json:"freight,omitempty"`
	CodPayAmount        float64           `json:"codPayAmount,omitempty"`
	BuyerMessage        string            `json:"buyerMessage,omitempty"`
	SellerRemarks       string            `json:"sellerRemarks,omitempty"`
	LogisticsSelected   string            `json:"logisticsSelected,omitempty"`
	TrackingNumber      string            `json:"trackingNumber,omitempty"`
	IsSpecifyBatch      bool              `json:"isSpecifyBatch,omitempty"`
	ShippingLabel       string            `json:"shippingLabel,omitempty"`
	ImgType             string            `json:"imgType,omitempty"`
	DocumentFile        string            `json:"documentFile,omitempty"`
	DocumentType        string            `json:"documentType,omitempty"`
	DocumentName        string            `json:"documentName,omitempty"`
	CustomerType        string            `json:"customerType,omitempty"`
	OrderCustomFieldValues []CustomFieldValue `json:"orderCustomFieldValueVOList,omitempty"`
}

// Create creates a new sales order in QERP.
func (s *OrderService) Create(params *CreateOrderParams) (*Order, error) {
	biz, _ := json.Marshal(params)
	var order Order
	w := &ResponseWrapper{Result: &order}
	if err := s.client.Do("CREATE_SALES_ORDER", string(biz), w); err != nil {
		return nil, err
	}
	if w.HasError() {
		return nil, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return &order, nil
}

// Cancel cancels a sales order by online order number and shop name.
func (s *OrderService) Cancel(onlineOrderNumber, shop string) error {
	params := map[string]any{"onlineOrderNumber": onlineOrderNumber, "shop": shop}
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("CLOSE_SALES_ORDER", string(biz), w)
}

// OrderQueryParams holds parameters for querying sales orders.
type OrderQueryParams struct {
	Page                int    `json:"page"`
	PageSize            int    `json:"pageSize"`
	Status              string `json:"status,omitempty"`
	Shop                string `json:"shop,omitempty"`
	OrderNumber         string `json:"orderNumber,omitempty"`
	OnlineOrderNumber   string `json:"onlineOrderNumber,omitempty"`
	FromPayTime         string `json:"fromPayTime,omitempty"`
	ToPayTime           string `json:"toPayTime,omitempty"`
	UpdateTimeFrom      string `json:"updateTimeFrom,omitempty"`
	UpdateTimeTo        string `json:"updateTimeTo,omitempty"`
	ShippingTimeFrom    string `json:"shippingTimeFrom,omitempty"`
	ShippingTimeTo      string `json:"shippingTimeTo,omitempty"`
	ShopIDList          []int64 `json:"shopIdList,omitempty"`
	OrderByParam        string `json:"orderByParam,omitempty"`
	OrderByOrder        string `json:"orderByOrder,omitempty"`
	CombineSku          *bool   `json:"combineSku,omitempty"`
	ReturnSnList        *bool   `json:"returnSnList,omitempty"`
	ReturnGiftFlag      *bool   `json:"returnGiftFlag,omitempty"`
}

// QueryList retrieves a paginated list of sales orders with optional filters.
func (s *OrderService) QueryList(params *OrderQueryParams) ([]Order, int, error) {
	biz, _ := json.Marshal(params)
	var orders []Order
	w := &ResponseWrapper{Result: &orders}
	if err := s.client.Do("QUERY_SALES_ORDER_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return orders, w.BizContent.Total, nil
}

// QueryNumberList retrieves a paginated list of sales order numbers.
func (s *OrderService) QueryNumberList(status, shop, fromPayTime, toPayTime string, page, pageSize int) ([]string, int, error) {
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
	biz, _ := json.Marshal(params)
	var numbers []string
	w := &ResponseWrapper{Result: &numbers}
	if err := s.client.Do("QUERY_SALES_ORDER_NUMBER_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return numbers, w.BizContent.Total, nil
}

// SalesOrderShippingInfoParams holds parameters for querying order shipping info.
type SalesOrderShippingInfoParams struct {
	OrderNumber string `json:"orderNumber"`
}

// QueryShippingInfo queries the shipping information for a sales order.
func (s *OrderService) QueryShippingInfo(orderNumber string) error {
	params := map[string]any{"orderNumber": orderNumber}
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("QUERY_SALES_ORDER_SHIPPING_INFO", string(biz), w)
}

// AuditParams holds parameters for auditing a sales order.
type AuditParams struct {
	OrderNumber       string `json:"orderNumber"`
	Shop              string `json:"shop"`
	OnlineOrderNumber string `json:"onlineOrderNumber"`
}

// Audit audits a sales order.
func (s *OrderService) Audit(params *AuditParams) error {
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("QUERY_SALES_ORDER_AUDIT", string(biz), w)
}

// CreateWaveOrderParams holds parameters for creating a wave order.
type CreateWaveOrderParams struct {
	OrderNumberList []string `json:"orderNumberList"`
	Shop            string   `json:"shop"`
}

// CreateWaveOrder creates a wave order for batch picking.
func (s *OrderService) CreateWaveOrder(params *CreateWaveOrderParams) error {
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("CREATE_WAVE_ORDER", string(biz), w)
}

// SendToWmsParams holds parameters for sending an order to WMS.
type SendToWmsParams struct {
	OrderNumber       string `json:"orderNumber"`
	OnlineOrderNumber string `json:"onlineOrderNumber,omitempty"`
	Shop              string `json:"shop"`
}

// SendToWms sends a sales order to the warehouse management system.
func (s *OrderService) SendToWms(params *SendToWmsParams) error {
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("SEND_SALES_ORDER_TO_WMS", string(biz), w)
}

// QueryOriginalOrder queries the original sales order by shop and online order number.
func (s *OrderService) QueryOriginalOrder(shop, onlineOrderNumber string) (*Order, error) {
	params := map[string]any{"shop": shop, "onlineOrderNumber": onlineOrderNumber}
	biz, _ := json.Marshal(params)
	var order Order
	w := &ResponseWrapper{Result: &order}
	if err := s.client.Do("QUERY_ORIGINAL_SALES_ORDER", string(biz), w); err != nil {
		return nil, err
	}
	if w.HasError() {
		return nil, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return &order, nil
}

// QueryPickupStatus queries the pickup status of a sales order.
func (s *OrderService) QueryPickupStatus(shop, onlineOrderNumber string) error {
	params := map[string]any{"shop": shop, "onlineOrderNumber": onlineOrderNumber}
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("QUERY_SALES_ORDER_PICKUP_STATUS", string(biz), w)
}

// QueryOrderDocument queries the document attached to a sales order.
func (s *OrderService) QueryOrderDocument(orderNumber string) error {
	params := map[string]any{"orderNumber": orderNumber}
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("QUERY_SALES_ORDER_DOCUMENT", string(biz), w)
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

// SubscribeOrder subscribes to order status push notifications.
func (s *OrderService) SubscribeOrder(orderType string, orderNumbers []string) ([]any, error) {
	list := make([]SubscribeOrderItem, len(orderNumbers))
	for i, n := range orderNumbers {
		list[i] = SubscribeOrderItem{OrderNumber: n}
	}
	params := SubscribeOrderParams{
		OrderType: orderType,
		OrderList: list,
	}
	biz, _ := json.Marshal(params)
	var result []any
	w := &ResponseWrapper{Result: &result}
	if err := s.client.Do("SUBSCRIBE_ORDER", string(biz), w); err != nil {
		return nil, err
	}
	if w.HasError() {
		return nil, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return result, nil
}
