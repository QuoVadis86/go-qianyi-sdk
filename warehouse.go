package qianyi

import "encoding/json"

type WarehouseService struct {
	client *Client
}

func NewWarehouseService(client *Client) *WarehouseService {
	return &WarehouseService{client: client}
}

func (s *WarehouseService) QueryList(page, pageSize int, status, name string) ([]Warehouse, int, error) {
	params := map[string]any{
		"page":     page,
		"pageSize": pageSize,
	}
	if status != "" {
		params["status"] = status
	}
	if name != "" {
		params["name"] = name
	}
	biz, _ := json.Marshal(params)
	var warehouses []Warehouse
	w := &ResponseWrapper{Result: &warehouses}
	if err := s.client.Do("QUERY_WAREHOUSE_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return warehouses, w.BizContent.Total, nil
}
