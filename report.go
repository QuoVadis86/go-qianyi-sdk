package qianyi

import "encoding/json"

type ReportService struct {
	client *Client
}

func NewReportService(client *Client) *ReportService {
	return &ReportService{client: client}
}

func (s *ReportService) QueryList(page, pageSize int, reportType string) ([]any, int, error) {
	params := map[string]any{
		"page":     page,
		"pageSize": pageSize,
	}
	if reportType != "" {
		params["type"] = reportType
	}
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_REPORT_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

func (s *ReportService) QuerySales(params any) (any, error) {
	biz, _ := json.Marshal(params)
	var result any
	w := &ResponseWrapper{Result: &result}
	if err := s.client.Do("QUERY_SALES_REPORT", string(biz), w); err != nil {
		return nil, err
	}
	if w.HasError() {
		return nil, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return result, nil
}
