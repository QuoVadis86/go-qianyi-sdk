package qianyi

import "context"

// ShopService provides access to shop-related API operations.
type ShopService struct {
	client *Client
}

// NewShopService creates a new ShopService.
func NewShopService(client *Client) *ShopService {
	return &ShopService{client: client}
}

// QueryList retrieves a paginated list of shops with optional filters.
func (s *ShopService) QueryList(ctx context.Context, page, pageSize int, platform, status, siteCode, name string) ([]Shop, int, error) {
	params := map[string]any{
		"page":     page,
		"pageSize": pageSize,
	}
	if platform != "" {
		params["platform"] = platform
	}
	if status != "" {
		params["status"] = status
	}
	if siteCode != "" {
		params["siteCode"] = siteCode
	}
	if name != "" {
		params["name"] = name
	}
	return doList[Shop](ctx, s.client, "QUERY_SHOP_LIST", params)
}
