package qianyi

import "context"

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
	Warehouse                   string            `json:"warehouse"`
	Shop                        string            `json:"shop"`
	OrderNumber                 string            `json:"orderNumber,omitempty"`
	Reason                      string            `json:"reason,omitempty"`
	Remark                      string            `json:"remark,omitempty"`
	Carrier                     string            `json:"carrier,omitempty"`
	CustomNumber                string            `json:"customNumber,omitempty"`
	AutoCommit                  *bool             `json:"autoCommit,omitempty"`
	ExpectArriveTime            string            `json:"expectArriveTime,omitempty"`
	ReturnSkuList               []ReturnSku       `json:"returnSkuList"`
	RefundCustomFieldValueList  []CustomFieldValue `json:"refundCustomFieldValueVOList,omitempty"`
}

// Create creates a new refund/return order in QERP.
func (s *RefundService) Create(ctx context.Context, params *CreateRefundParams) (*ReturnOrder, error) {
	return doSingle[ReturnOrder](ctx, s.client, ServiceTypeCreateReturnOrder, params)
}

// Cancel cancels a refund/return order by return number.
func (s *RefundService) Cancel(ctx context.Context, returnNumber string) error {
	return doAction(ctx, s.client, ServiceTypeCloseReturnOrder, map[string]any{"returnNumber": returnNumber})
}

// RefundQueryParams holds parameters for querying refund orders.
type RefundQueryParams struct {
	Page            int     `json:"page"`
	PageSize        int     `json:"pageSize"`
	ReturnNumber    string  `json:"returnNumber,omitempty"`
	Warehouse       string  `json:"warehouse,omitempty"`
	Status          string  `json:"status,omitempty"`
	FromCreateTime  string  `json:"fromCreateTime,omitempty"`
	ToCreateTime    string  `json:"toCreateTime,omitempty"`
	UpdateTimeFrom  string  `json:"updateTimeFrom,omitempty"`
	UpdateTimeTo    string  `json:"updateTimeTo,omitempty"`
	FinishTimeFrom  string  `json:"finishTimeFrom,omitempty"`
	FinishTimeTo    string  `json:"finishTimeTo,omitempty"`
	ReceiveTimeFrom string  `json:"receiveTimeFrom,omitempty"`
	ReceiveTimeTo   string  `json:"receiveTimeTo,omitempty"`
	ShopIDList      []int64 `json:"shopIdList,omitempty"`
	ShopGroupIDList []int64 `json:"shopGroupIdList,omitempty"`
	OrderStatus     string  `json:"orderStatus,omitempty"`
	WithoutType     *bool   `json:"withoutType,omitempty"`
	WithCommitTime  *bool   `json:"withCommitTime,omitempty"`
}

// QueryList retrieves a paginated list of refund orders with optional filters.
func (s *RefundService) QueryList(ctx context.Context, params *RefundQueryParams) ([]ReturnOrder, int, error) {
	return doList[ReturnOrder](ctx, s.client, ServiceTypeQueryReturnOrderList, params)
}

// PushReturnOrderInfoParams holds parameters for pushing refund return info.
type PushReturnOrderInfoParams struct {
	ReturnNumber      string      `json:"returnNumber"`
	OrderNumber       string      `json:"orderNumber"`
	OnlineOrderNumber string      `json:"onlineOrderNumber"`
	Status            string      `json:"status"`
	ReturnSkuList     []ReturnSku `json:"returnSkuList"`
}

// PushReturnInfo pushes refund/return order process information.
func (s *RefundService) PushReturnInfo(ctx context.Context, params *PushReturnOrderInfoParams) error {
	return doAction(ctx, s.client, ServiceTypePushReturnOrderInfo, params)
}
