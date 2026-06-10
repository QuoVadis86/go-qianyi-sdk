package qianyi

import "encoding/json"

type SkuService struct {
	client *Client
}

func NewSkuService(client *Client) *SkuService {
	return &SkuService{client: client}
}

func (s *SkuService) QueryList(page, pageSize int, opts ...SkuQueryOption) ([]Sku, int, error) {
	params := map[string]any{
		"page":     page,
		"pageSize": pageSize,
	}
	for _, opt := range opts {
		opt(params)
	}
	biz, _ := json.Marshal(params)

	var skus []Sku
	w := &ResponseWrapper{Result: &skus}
	if err := s.client.Do("QUERY_SIMPLE_LIST_SKU", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return skus, w.BizContent.Total, nil
}

type SkuQueryOption func(map[string]any)

func SkuFilterBySKU(skus []string) SkuQueryOption {
	return func(m map[string]any) {
		m["skus"] = skus
	}
}

func SkuFilterByUpdateTime(from, to string) SkuQueryOption {
	return func(m map[string]any) {
		if from != "" {
			m["updateTimeFrom"] = from
		}
		if to != "" {
			m["updateTimeTo"] = to
		}
	}
}

func SkuFilterByCategory(c1, c2, c3 string) SkuQueryOption {
	return func(m map[string]any) {
		if c1 != "" {
			m["categoryName1"] = c1
		}
		if c2 != "" {
			m["categoryName2"] = c2
		}
		if c3 != "" {
			m["categoryName3"] = c3
		}
	}
}

func SkuFilterByTitle(titles []string) SkuQueryOption {
	return func(m map[string]any) {
		m["titles"] = titles
	}
}

func (s *SkuService) Create(sku *Sku) error {
	biz, _ := json.Marshal(sku)
	w := &ResponseWrapper{}
	if err := s.client.Do("INSERT_SKU_INFO", string(biz), w); err != nil {
		return err
	}
	if w.HasError() {
		return &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return nil
}

func (s *SkuService) Update(sku *Sku) error {
	biz, _ := json.Marshal(sku)
	w := &ResponseWrapper{}
	if err := s.client.Do("UPDATE_SKU_INFO", string(biz), w); err != nil {
		return err
	}
	if w.HasError() {
		return &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return nil
}

func (s *SkuService) Enable(sku string, enable int) error {
	params := map[string]any{"sku": sku, "enable": enable}
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	if err := s.client.Do("ENABLE_SKU", string(biz), w); err != nil {
		return err
	}
	if w.HasError() {
		return &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return nil
}

func (s *SkuService) QuerySysSKU(page, pageSize int, skus []string) ([]Sku, int, error) {
	params := map[string]any{
		"page":     page,
		"pageSize": pageSize,
	}
	if len(skus) > 0 {
		params["skus"] = skus
	}
	biz, _ := json.Marshal(params)

	var result []Sku
	w := &ResponseWrapper{Result: &result}
	if err := s.client.Do("QUERY_SYS_SKU", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return result, w.BizContent.Total, nil
}
