package qianyi

import "context"

// WarehouseService provides access to warehouse API operations.
type WarehouseService struct {
	client *Client
}

// NewWarehouseService creates a new WarehouseService.
func NewWarehouseService(client *Client) *WarehouseService {
	return &WarehouseService{client: client}
}

// QueryList retrieves a paginated list of warehouses with optional filters.
func (s *WarehouseService) QueryList(ctx context.Context, page, pageSize int, status, name string) ([]Warehouse, int, error) {
	params := map[string]any{"page": page, "pageSize": pageSize}
	if status != "" {
		params["status"] = status
	}
	if name != "" {
		params["name"] = name
	}
	return doList[Warehouse](ctx, s.client, ServiceTypeQueryWarehouseList, params)
}
