package qianyi

import "encoding/json"

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
	TableName  string `json:"tableName"`
	ColumType  string `json:"columType,omitempty"`
	ColumName  string `json:"columName,omitempty"`
	Required   int    `json:"required,omitempty"`
	IsQuery    int    `json:"isQuery,omitempty"`
	IsShow     int    `json:"isShow,omitempty"`
}

// Query retrieves custom field definitions for a given table.
func (s *CustomerFieldService) Query(params *CustomerFieldQueryParams) ([]CustomField, error) {
	biz, _ := json.Marshal(params)
	var list []CustomField
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("CUSTOMER_FIELD_QUERY", string(biz), w); err != nil {
		return nil, err
	}
	if w.HasError() {
		return nil, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, nil
}
