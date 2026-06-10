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
	Shop              string     `json:"shop"`
	OnlineOrderNumber string     `json:"onlineOrderNumber"`
	PaymentMethod     string     `json:"paymentMethod"`
	Currency          string     `json:"currency"`
	PayTime           string     `json:"payTime"`
	Buyer             *Buyer     `json:"buyer"`
	SkuList           []OrderSku `json:"skuList"`
	Freight           float64    `json:"freight,omitempty"`
	CodPayAmount      float64    `json:"codPayAmount,omitempty"`
	BuyerMessage      string     `json:"buyerMessage,omitempty"`
	SellerRemarks     string     `json:"sellerRemarks,omitempty"`
	LogisticsSelected string     `json:"logisticsSelected,omitempty"`
	TrackingNumber    string     `json:"trackingNumber,omitempty"`
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
	Page              int    `json:"page"`
	PageSize          int    `json:"pageSize"`
	Status            string `json:"status,omitempty"`
	Shop              string `json:"shop,omitempty"`
	OrderNumber       string `json:"orderNumber,omitempty"`
	OnlineOrderNumber string `json:"onlineOrderNumber,omitempty"`
	FromPayTime       string `json:"fromPayTime,omitempty"`
	ToPayTime         string `json:"toPayTime,omitempty"`
	UpdateTimeFrom    string `json:"updateTimeFrom,omitempty"`
	UpdateTimeTo      string `json:"updateTimeTo,omitempty"`
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
