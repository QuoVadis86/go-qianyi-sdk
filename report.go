package qianyi

import "encoding/json"

// ReportService provides access to financial report API operations.
type ReportService struct {
	client *Client
}

// NewReportService creates a new ReportService.
func NewReportService(client *Client) *ReportService {
	return &ReportService{client: client}
}

// ReportQueryParams holds base parameters for report queries.
type ReportQueryParams struct {
	Page       int      `json:"page"`
	PageSize   int      `json:"pageSize"`
	ShopIDList []int64  `json:"shopIdList,omitempty"`
	ShopNameList []string `json:"shopNameList,omitempty"`
}

// QueryShopeeTransaction queries Shopee transaction details.
func (s *ReportService) QueryShopeeTransaction(params *ReportQueryParams) ([]any, int, error) {
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_SHOPEE_TRANSACTION_DETAIL_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// QueryLazadaTransaction queries Lazada transaction details.
func (s *ReportService) QueryLazadaTransaction(params *ReportQueryParams) ([]any, int, error) {
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_LAZADA_TRANSACTION_DETAIL_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// QueryTiktokTransaction queries TikTok transaction details.
func (s *ReportService) QueryTiktokTransaction(params *ReportQueryParams) ([]any, int, error) {
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_TIKTOK_TRANSACTION_DETAIL_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// QueryShopeePayout queries Shopee payout records.
func (s *ReportService) QueryShopeePayout(params *ReportQueryParams) ([]any, int, error) {
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_SHOPEE_PAYOUT_DETAIL_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// QueryInventoryDailyReport queries inventory daily statements.
func (s *ReportService) QueryInventoryDailyReport(params any) ([]any, int, error) {
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_INVENTORY_DAILY_REPORT", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}
