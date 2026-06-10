package qianyi

import "encoding/json"

// RefundService provides access to refund/return order API operations.
type RefundService struct {
	client *Client
}

// NewRefundService creates a new RefundService.
func NewRefundService(client *Client) *RefundService {
	return &RefundService{client: client}
}

// CreateRefundParams holds the parameters for creating a refund/return order.
type CreateRefundParams struct {
	Warehouse     string      `json:"warehouse"`
	Shop          string      `json:"shop"`
	OrderNumber   string      `json:"orderNumber,omitempty"`
	Reason        string      `json:"reason,omitempty"`
	Remark        string      `json:"remark,omitempty"`
	Carrier       string      `json:"carrier,omitempty"`
	CustomNumber  string      `json:"customNumber,omitempty"`
	ReturnSkuList []ReturnSku `json:"returnSkuList"`
}

// Create creates a new refund/return order in QERP.
func (s *RefundService) Create(params *CreateRefundParams) (*ReturnOrder, error) {
	biz, _ := json.Marshal(params)
	var ret ReturnOrder
	w := &ResponseWrapper{Result: &ret}
	if err := s.client.Do("CREATE_RETURN_ORDER", string(biz), w); err != nil {
		return nil, err
	}
	if w.HasError() {
		return nil, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return &ret, nil
}

// Cancel cancels a refund/return order by return number.
func (s *RefundService) Cancel(returnNumber string) error {
	params := map[string]any{"returnNumber": returnNumber}
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("CLOSE_RETURN_ORDER", string(biz), w)
}

// RefundQueryParams holds parameters for querying refund orders.
type RefundQueryParams struct {
	Page           int    `json:"page"`
	PageSize       int    `json:"pageSize"`
	ReturnNumber   string `json:"returnNumber,omitempty"`
	Warehouse      string `json:"warehouse,omitempty"`
	Status         string `json:"status,omitempty"`
	FromCreateTime string `json:"fromCreateTime,omitempty"`
	ToCreateTime   string `json:"toCreateTime,omitempty"`
	UpdateTimeFrom string `json:"updateTimeFrom,omitempty"`
	UpdateTimeTo   string `json:"updateTimeTo,omitempty"`
}

// QueryList retrieves a paginated list of refund orders with optional filters.
func (s *RefundService) QueryList(params *RefundQueryParams) ([]ReturnOrder, int, error) {
	biz, _ := json.Marshal(params)
	var list []ReturnOrder
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_RETURN_ORDER_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}
