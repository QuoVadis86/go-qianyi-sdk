package qianyi

import "encoding/json"

type PurchaseService struct {
	client *Client
}

func NewPurchaseService(client *Client) *PurchaseService {
	return &PurchaseService{client: client}
}

func (s *PurchaseService) QueryList(page, pageSize int, status, warehouse string) ([]any, int, error) {
	params := map[string]any{
		"page":     page,
		"pageSize": pageSize,
	}
	if status != "" {
		params["status"] = status
	}
	if warehouse != "" {
		params["warehouseName"] = warehouse
	}
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_PURCHASE_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

func (s *PurchaseService) Create(params any) error {
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	if err := s.client.Do("CREATE_PURCHASE_ORDER", string(biz), w); err != nil {
		return err
	}
	if w.HasError() {
		return &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return nil
}

func (s *PurchaseService) Update(params any) error {
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	if err := s.client.Do("UPDATE_PURCHASE_ORDER", string(biz), w); err != nil {
		return err
	}
	if w.HasError() {
		return &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return nil
}
