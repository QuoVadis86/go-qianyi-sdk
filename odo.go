package qianyi

import "encoding/json"

// OdoService provides access to outbound delivery order (ODO) API operations.
type OdoService struct {
	client *Client
}

// NewOdoService creates a new OdoService.
func NewOdoService(client *Client) *OdoService {
	return &OdoService{client: client}
}

// QueryList retrieves a paginated list of outbound delivery orders.
func (s *OdoService) QueryList(page, pageSize int, status, warehouse string) ([]any, int, error) {
	params := map[string]any{"page": page, "pageSize": pageSize}
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

// Close closes an outbound delivery order by ODO number.
func (s *OdoService) Close(odoNumber string) error {
	params := map[string]any{"odoNumber": odoNumber}
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("CLOSE_ODO", string(biz), w)
}
