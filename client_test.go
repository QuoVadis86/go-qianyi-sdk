package qianyi

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

// mockHTTPClient implements HTTPClient for testing.
type mockHTTPClient struct {
	responseBody string
	statusCode   int
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: m.statusCode,
		Body:       io.NopCloser(strings.NewReader(m.responseBody)),
	}, nil
}

func TestGenerateSign(t *testing.T) {
	c := NewClient("test-app-id", "test-app-secret")
	sign := c.GenerateSign(ServiceTypeQueryShopList, `{"page":1,"pageSize":10}`, 1234567890)
	if sign == "" {
		t.Fatal("expected non-empty sign")
	}
	if len(sign) != 32 {
		t.Fatalf("expected 32-char hex, got %d: %s", len(sign), sign)
	}
}

func TestParseResponse_Success(t *testing.T) {
	body := `{"state":"success","errorCode":"","errorMsg":"","bizContent":"{\"state\":\"success\",\"total\":5,\"result\":[{\"shopId\":1}]}","requestId":"req-123"}`
	var base BaseResponse
	err := parseResponse([]byte(body), &base)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !base.IsSuccess() {
		t.Fatal("expected success")
	}
	if base.RequestID != "req-123" {
		t.Fatalf("expected req-123, got %s", base.RequestID)
	}
}

func TestParseResponse_APIError(t *testing.T) {
	body := `{"state":"failure","errorCode":"ERR_001","errorMsg":"invalid params","bizContent":"","requestId":"req-456"}`
	err := parseResponse([]byte(body), nil)
	if err == nil {
		t.Fatal("expected error")
	}
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.ErrorCode != "ERR_001" {
		t.Fatalf("expected ERR_001, got %s", apiErr.ErrorCode)
	}
	if apiErr.RequestID != "req-456" {
		t.Fatalf("expected req-456, got %s", apiErr.RequestID)
	}
}

func TestParseResponse_ResponseWrapper(t *testing.T) {
	body := `{"state":"success","errorCode":"","errorMsg":"","bizContent":"{\"state\":\"success\",\"total\":2,\"result\":[{\"shopId\":1,\"name\":\"Shop1\"},{\"shopId\":2,\"name\":\"Shop2\"}]}","requestId":"req-789"}`
	var shops []Shop
	w := &ResponseWrapper{Result: &shops}
	err := parseResponse([]byte(body), w)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if w.RequestID != "req-789" {
		t.Fatalf("expected req-789, got %s", w.RequestID)
	}
	if w.BizContent.Total != 2 {
		t.Fatalf("expected total=2, got %d", w.BizContent.Total)
	}
	if len(shops) != 2 {
		t.Fatalf("expected 2 shops, got %d", len(shops))
	}
	if shops[0].ShopID != 1 || shops[0].Name != "Shop1" {
		t.Fatalf("unexpected shop: %+v", shops[0])
	}
}

func TestParseResponse_WrapperWithError(t *testing.T) {
	body := `{"state":"success","errorCode":"ERR_002","errorMsg":"business error","bizContent":"","requestId":"req-err"}`
	w := &ResponseWrapper{}
	err := parseResponse([]byte(body), w)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !w.HasError() {
		t.Fatal("expected HasError")
	}
	if w.ErrorCode != "ERR_002" {
		t.Fatalf("expected ERR_002, got %s", w.ErrorCode)
	}
}

func TestClient_UserAgent(t *testing.T) {
	var capturedReq *http.Request
	c := NewClient("app", "secret")
	c.HTTPClient = httpClientFunc(func(req *http.Request) (*http.Response, error) {
		capturedReq = req
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"state":"success","errorCode":"","errorMsg":"","bizContent":"{}","requestId":"req-ua"}`)),
		}, nil
	})

	var base BaseResponse
	err := c.Do(context.Background(), "QUERY_SHOP_LIST", "{}", &base)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	ua := capturedReq.Header.Get("User-Agent")
	if ua != "go-qianyi-sdk/1.0" {
		t.Fatalf("expected User-Agent 'go-qianyi-sdk/1.0', got %q", ua)
	}
}

func TestClient_WithContextCancel(t *testing.T) {
	c := NewClient("app", "secret")
	c.HTTPClient = httpClientFunc(func(req *http.Request) (*http.Response, error) {
		// Simulate a check for context cancellation
		if err := req.Context().Err(); err != nil {
			return nil, err
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"state":"success","errorCode":"","errorMsg":"","bizContent":"{}","requestId":"req-cancel"}`)),
		}, nil
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	var base BaseResponse
	err := c.Do(ctx, ServiceTypeQueryShopList, "{}", &base)
	if err == nil {
		t.Fatal("expected error for cancelled context")
	}
}

func TestClientOption_WithBaseURL(t *testing.T) {
	c := NewClient("app", "secret", WithBaseURL("https://asia.qianyierp.com/"))
	if c.BaseURL != "https://asia.qianyierp.com" {
		t.Fatalf("expected https://asia.qianyierp.com, got %s", c.BaseURL)
	}
}

func TestAPIError_Error(t *testing.T) {
	err := &APIError{ErrorCode: "ERR_001", Message: "test error", RequestID: "req-1"}
	msg := err.Error()
	if !strings.Contains(msg, "ERR_001") || !strings.Contains(msg, "test error") {
		t.Fatalf("unexpected error message: %s", msg)
	}
}

func TestBaseResponse_IsSuccess(t *testing.T) {
	r := &BaseResponse{State: "success", ErrorCode: ""}
	if !r.IsSuccess() {
		t.Fatal("expected IsSuccess")
	}
	if r.HasError() {
		t.Fatal("expected !HasError")
	}

	r2 := &BaseResponse{State: "failure", ErrorCode: "ERR"}
	if r2.HasError() != true {
		t.Fatal("expected HasError")
	}
}

func TestTruncate(t *testing.T) {
	if truncate("hello", 10) != "hello" {
		t.Fatal("should not truncate short string")
	}
	if truncate("hello world", 5) != "hello..." {
		t.Fatal("should truncate with ellipsis")
	}
}

func TestAPIError_ImplementsError(t *testing.T) {
	var _ error = (*APIError)(nil)
}

func TestClientOptions(t *testing.T) {
	c := NewClient("app", "secret")
	if c.BaseURL != defaultBaseURL {
		t.Fatalf("expected default base URL %s, got %s", defaultBaseURL, c.BaseURL)
	}
	if c.AppID != "app" || c.AppSecret != "secret" {
		t.Fatal("app credentials not set")
	}
}

// httpClientFunc wraps a function as an HTTPClient.
type httpClientFunc func(*http.Request) (*http.Response, error)

func (f httpClientFunc) Do(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestEndpointForService(t *testing.T) {
	tests := []struct {
		serviceType string
		expected    string
	}{
		{ServiceTypeQueryShopList, "shop"},
		{ServiceTypeQueryWarehouseList, "warehouse"},
		{ServiceTypeQuerySimpleListSku, "sku"},
		{ServiceTypeInsertSkuInfo, "sku"},
		{ServiceTypeCreateSalesOrder, "salesOrder"},
		{ServiceTypeQueryReturnOrderList, "returnOrder"},
		{ServiceTypeCreateAsnOrder, "asn"},
		{ServiceTypeQueryInventoryLogList, "inventory"},
		{ServiceTypeQueryFirstLegOrderList, "firstLeg"},
		{ServiceTypeCustomerFieldQuery, "property"},
		{"UNKNOWN_TYPE", ""},
	}
	for _, tt := range tests {
		t.Run(tt.serviceType, func(t *testing.T) {
			got := endpointForService(tt.serviceType)
			if got != tt.expected {
				t.Errorf("endpointForService(%q) = %q, want %q", tt.serviceType, got, tt.expected)
			}
		})
	}
}

func TestClient_Do_HTTPError(t *testing.T) {
	c := NewClient("app", "secret")
	c.HTTPClient = httpClientFunc(func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       io.NopCloser(strings.NewReader("")),
		}, nil
	})

	var base BaseResponse
	err := c.Do(context.Background(), ServiceTypeQueryShopList, "{}", &base)
	if err == nil {
		t.Fatal("expected error for 500 response")
	}
}

func TestDoList(t *testing.T) {
	mock := &mockHTTPClient{
		responseBody: `{"state":"success","errorCode":"","errorMsg":"","bizContent":"{\"state\":\"success\",\"total\":1,\"result\":[{\"shopId\":1,\"name\":\"TestShop\"}]}","requestId":"req-list"}`,
		statusCode:   http.StatusOK,
	}
	c := NewClient("app", "secret", WithHTTPClient(mock))

	params := map[string]any{"page": 1, "pageSize": 10}
	result, total, err := doList[Shop](context.Background(), c, ServiceTypeQueryShopList, params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total != 1 {
		t.Fatalf("expected total=1, got %d", total)
	}
	if len(result) != 1 || result[0].Name != "TestShop" {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestDoAction(t *testing.T) {
	mock := &mockHTTPClient{
		responseBody: `{"state":"success","errorCode":"","errorMsg":"","bizContent":"{}","requestId":"req-action"}`,
		statusCode:   http.StatusOK,
	}
	c := NewClient("app", "secret", WithHTTPClient(mock))

	err := doAction(context.Background(), c, ServiceTypeCloseSalesOrder, map[string]any{"orderNumber": "ORD-001"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDoAction_BizError(t *testing.T) {
	mock := &mockHTTPClient{
		responseBody: `{"state":"success","errorCode":"BIZ_ERR","errorMsg":"business error","bizContent":"{}","requestId":"req-biz-err"}`,
		statusCode:   http.StatusOK,
	}
	c := NewClient("app", "secret", WithHTTPClient(mock))

	err := doAction(context.Background(), c, "TEST", map[string]any{})
	if err == nil {
		t.Fatal("expected business error")
	}
	apiErr, ok := err.(*APIError)
	if !ok || apiErr.ErrorCode != "BIZ_ERR" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestParseBizContent(t *testing.T) {
	r := &BaseResponse{
		BizContent: `{"state":"success","total":3,"result":[1,2,3]}`,
	}
	bc, err := r.ParseBizContent()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if bc.Total != 3 {
		t.Fatalf("expected total=3, got %d", bc.Total)
	}
	var nums []int
	if err := json.Unmarshal(bc.Result, &nums); err != nil {
		t.Fatalf("unmarshal result: %v", err)
	}
	if len(nums) != 3 || nums[0] != 1 {
		t.Fatalf("unexpected nums: %v", nums)
	}
}

func TestParseBizContent_Empty(t *testing.T) {
	r := &BaseResponse{}
	bc, err := r.ParseBizContent()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if bc == nil {
		t.Fatal("expected non-nil BizContent")
	}
}

func TestSDK_TestEnv(t *testing.T) {
	sdk := NewSDK("app", "secret")
	if sdk.Client.BaseURL != defaultBaseURL {
		t.Fatalf("expected default URL, got %s", sdk.Client.BaseURL)
	}
	sdk.TestEnv()
	if sdk.Client.BaseURL != "https://gerp-test1.800best.com" {
		t.Fatalf("expected test URL, got %s", sdk.Client.BaseURL)
	}
}

func TestNewSDK_Services(t *testing.T) {
	sdk := NewSDK("app", "secret")
	if sdk.Shop == nil {
		t.Fatal("Shop service not initialized")
	}
	if sdk.Order == nil {
		t.Fatal("Order service not initialized")
	}
	if sdk.Inventory == nil {
		t.Fatal("Inventory service not initialized")
	}
	if sdk.Report == nil {
		t.Fatal("Report service not initialized")
	}
	if sdk.CustomerField == nil {
		t.Fatal("CustomerField service not initialized")
	}
}
