package qianyi

import "encoding/json"

type ShopService struct {
	client *Client
}

func NewShopService(client *Client) *ShopService {
	return &ShopService{client: client}
}

func (s *ShopService) QueryList(page, pageSize int, platform, status, siteCode, name string) ([]Shop, int, error) {
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
	biz, _ := json.Marshal(params)

	var shops []Shop
	w := &ResponseWrapper{Result: &shops}
	if err := s.client.Do("QUERY_SHOP_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return shops, w.BizContent.Total, nil
}
