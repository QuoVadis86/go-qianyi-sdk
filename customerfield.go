package qianyi

import "context"

// CustomerFieldService provides access to custom field query API operations.
type CustomerFieldService struct {
	client *Client
}

// NewCustomerFieldService creates a new CustomerFieldService.
func NewCustomerFieldService(client *Client) *CustomerFieldService {
	return &CustomerFieldService{client: client}
}

// CustomerFieldQueryParams holds parameters for querying custom fields.
type CustomerFieldQueryParams struct {
	TableName string `json:"tableName"`
	ColumType string `json:"columType,omitempty"`
	ColumName string `json:"columName,omitempty"`
	Required  int    `json:"required,omitempty"`
	IsQuery   int    `json:"isQuery,omitempty"`
	IsShow    int    `json:"isShow,omitempty"`
}

// Query retrieves custom field definitions for a given table.
func (s *CustomerFieldService) Query(ctx context.Context, params *CustomerFieldQueryParams) ([]CustomField, error) {
	return doListNoTotal[CustomField](ctx, s.client, ServiceTypeCustomerFieldQuery, params)
}
