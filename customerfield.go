package qianyi

import "encoding/json"

type CustomerFieldService struct {
	client *Client
}

func NewCustomerFieldService(client *Client) *CustomerFieldService {
	return &CustomerFieldService{client: client}
}

func (s *CustomerFieldService) Query(tableName string) ([]any, error) {
	params := map[string]any{"tableName": tableName}
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("CUSTOMER_FIELD_QUERY", string(biz), w); err != nil {
		return nil, err
	}
	if w.HasError() {
		return nil, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, nil
}
