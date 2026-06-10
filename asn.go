package qianyi

import "encoding/json"

type AsnService struct {
	client *Client
}

func NewAsnService(client *Client) *AsnService {
	return &AsnService{client: client}
}

type AsnSku struct {
	Sku            string  `json:"sku"`
	ExpectQuantity float64 `json:"expectQuantity"`
}

type CreateAsnParams struct {
	WarehouseName string  `json:"warehouseName"`
	AsnSkuVOList  []AsnSku `json:"asnSkuVOList"`
	Remark        string  `json:"remark,omitempty"`
}

func (s *AsnService) Create(params *CreateAsnParams) error {
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	if err := s.client.Do("CREATE_ASN_ORDER", string(biz), w); err != nil {
		return err
	}
	if w.HasError() {
		return &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return nil
}

func (s *AsnService) QueryList(page, pageSize int, warehouse string, status string) ([]any, int, error) {
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
	if err := s.client.Do("QUERY_ASN_ORDER_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

func (s *AsnService) Close(asnNumber string) error {
	params := map[string]any{"asnNumber": asnNumber}
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	if err := s.client.Do("CLOSE_ASN_ORDER", string(biz), w); err != nil {
		return err
	}
	if w.HasError() {
		return &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return nil
}
