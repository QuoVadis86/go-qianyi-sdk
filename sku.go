package qianyi

import "encoding/json"

// SkuService provides access to SKU/product-related API operations.
type SkuService struct {
	client *Client
}

// NewSkuService creates a new SkuService.
func NewSkuService(client *Client) *SkuService {
	return &SkuService{client: client}
}

// QueryList retrieves a paginated list of SKUs with optional filters.
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

// SkuQueryOption is a functional option for filtering SKU list queries.
type SkuQueryOption func(map[string]any)

// SkuFilterBySKU filters by specific SKU codes.
func SkuFilterBySKU(skus []string) SkuQueryOption {
	return func(m map[string]any) { m["skus"] = skus }
}

// SkuFilterByUpdateTime filters by update time range (format: YYYY-MM-DD).
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

// SkuFilterByCategory filters by product category levels.
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

// SkuFilterByTitle filters by product titles.
func SkuFilterByTitle(titles []string) SkuQueryOption {
	return func(m map[string]any) { m["titles"] = titles }
}

// Create creates a new SKU in QERP.
func (s *SkuService) Create(sku *Sku) error {
	biz, _ := json.Marshal(sku)
	w := &ResponseWrapper{}
	return s.client.Do("INSERT_SKU_INFO", string(biz), w)
}

// Update updates an existing SKU in QERP.
func (s *SkuService) Update(sku *Sku) error {
	biz, _ := json.Marshal(sku)
	w := &ResponseWrapper{}
	return s.client.Do("UPDATE_SKU_INFO", string(biz), w)
}

// Enable enables or disables a SKU. Set enable to 1 for active, 0 for inactive.
func (s *SkuService) Enable(sku string, enable int) error {
	params := map[string]any{"sku": sku, "enable": enable}
	biz, _ := json.Marshal(params)
	w := &ResponseWrapper{}
	return s.client.Do("ENABLE_SKU", string(biz), w)
}

// QuerySysSKU queries system SKUs with optional SKU code filter.
func (s *SkuService) QuerySysSKU(page, pageSize int, skus []string) ([]Sku, int, error) {
	params := map[string]any{"page": page, "pageSize": pageSize}
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
