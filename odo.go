package qianyi

import "encoding/json"

type OdoService struct {
	client *Client
}

func NewOdoService(client *Client) *OdoService {
	return &OdoService{client: client}
}

func (s *OdoService) QueryList(page, pageSize int, status, warehouse string) ([]any, int, error) {
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
	if err := s.client.Do("QUERY_ODO_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

func (s *OdoService) Close(odoNumber string) error {
	params := map[string]any{"odoNumber": odoNumber}
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	if err := s.client.Do("CLOSE_ODO", string(biz), w); err != nil {
		return err
	}
	if w.HasError() {
		return &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return nil
}
