package qianyi

import (
	"context"
	"encoding/json"
)

// SkuService provides access to SKU/product-related API operations.
type SkuService struct {
	client *Client
}

// NewSkuService creates a new SkuService.
func NewSkuService(client *Client) *SkuService {
	return &SkuService{client: client}
}

// QueryList retrieves a paginated list of SKUs with optional filters.
func (s *SkuService) QueryList(ctx context.Context, page, pageSize int, opts ...SkuQueryOption) ([]Sku, int, error) {
	params := map[string]any{
		"page":     page,
		"pageSize": pageSize,
	}
	for _, opt := range opts {
		opt(params)
	}
	return doList[Sku](ctx, s.client, "QUERY_SIMPLE_LIST_SKU", params)
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
func (s *SkuService) Create(ctx context.Context, sku *Sku) error {
	biz, err := json.Marshal(sku)
	if err != nil {
		return err
	}
	w := &ResponseWrapper{}
	return s.client.Do(ctx, "INSERT_SKU_INFO", string(biz), w)
}

// Update updates an existing SKU in QERP.
func (s *SkuService) Update(ctx context.Context, sku *Sku) error {
	biz, err := json.Marshal(sku)
	if err != nil {
		return err
	}
	w := &ResponseWrapper{}
	return s.client.Do(ctx, "UPDATE_SKU_INFO", string(biz), w)
}

// Enable enables or disables a SKU. Set enable to 1 for active, 0 for inactive.
func (s *SkuService) Enable(ctx context.Context, sku string, enable int) error {
	return doAction(ctx, s.client, "ENABLE_SKU", map[string]any{"sku": sku, "enable": enable})
}

// QuerySysSKU queries system SKUs with optional SKU code filter.
func (s *SkuService) QuerySysSKU(ctx context.Context, page, pageSize int, skus []string) ([]Sku, int, error) {
	params := map[string]any{"page": page, "pageSize": pageSize}
	if len(skus) > 0 {
		params["skus"] = skus
	}
	return doList[Sku](ctx, s.client, "QUERY_SYS_SKU", params)
}
