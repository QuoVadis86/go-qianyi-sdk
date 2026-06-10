package qianyi

import "encoding/json"

type AdjustService struct {
	client *Client
}

func NewAdjustService(client *Client) *AdjustService {
	return &AdjustService{client: client}
}

func (s *AdjustService) QueryList(page, pageSize int, warehouse, status string) ([]any, int, error) {
	params := map[string]any{
		"page":     page,
		"pageSize": pageSize,
	}
	if warehouse != "" {
		params["warehouseName"] = warehouse
	}
	if status != "" {
		params["status"] = status
	}
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_ADJUST_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

func (s *AdjustService) Create(params any) error {
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	if err := s.client.Do("CREATE_ADJUST_ORDER", string(biz), w); err != nil {
		return err
	}
	if w.HasError() {
		return &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return nil
}
