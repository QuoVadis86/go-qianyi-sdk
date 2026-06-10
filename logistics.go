package qianyi

import "encoding/json"

type LogisticsService struct {
	client *Client
}

func NewLogisticsService(client *Client) *LogisticsService {
	return &LogisticsService{client: client}
}

type LogisticsChannel struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Carrier  string `json:"carrier,omitempty"`
}

func (s *LogisticsService) QueryChannelList(page, pageSize int) ([]LogisticsChannel, int, error) {
	params := map[string]any{
		"page":     page,
		"pageSize": pageSize,
	}
	biz, _ := json.Marshal(params)
	var channels []LogisticsChannel
	w := &ResponseWrapper{Result: &channels}
	if err := s.client.Do("QUERY_LOGISTICS_CHANNEL_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return channels, w.BizContent.Total, nil
}

func (s *LogisticsService) QueryTracking(trackingNumber string) (any, error) {
	params := map[string]any{"trackingNumber": trackingNumber}
	biz, _ := json.Marshal(params)
	var result any
	w := &ResponseWrapper{Result: &result}
	if err := s.client.Do("QUERY_LOGISTICS_TRACKING", string(biz), w); err != nil {
		return nil, err
	}
	if w.HasError() {
		return nil, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return result, nil
}
