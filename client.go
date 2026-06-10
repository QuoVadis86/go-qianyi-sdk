package qianyi

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"sort"
	"strings"
	"time"
)

// HTTPClient is the interface for making HTTP requests.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var defaultClient HTTPClient = &http.Client{Timeout: 30 * time.Second}

// Client handles HTTP communication with the QERP Open Platform API.
type Client struct {
	AppID      string
	AppSecret  string
	BaseURL    string
	HTTPClient HTTPClient
}

// NewClient creates a new QERP API client with the given appId and appSecret.
// Default base URL is the domestic production environment (www.qianyierp.com).
func NewClient(appID, appSecret string, opts ...ClientOption) *Client {
	c := &Client{
		AppID:      appID,
		AppSecret:  appSecret,
		BaseURL:    "https://www.qianyierp.com",
		HTTPClient: defaultClient,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// ClientOption is a functional option for configuring a Client.
type ClientOption func(*Client)

// WithBaseURL overrides the default API base URL.
// Use https://gerp-test1.800best.com for testing or https://asia.qianyierp.com for overseas.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.BaseURL = strings.TrimRight(baseURL, "/")
	}
}

// WithHTTPClient sets a custom HTTP client implementation.
func WithHTTPClient(hc HTTPClient) ClientOption {
	return func(c *Client) {
		c.HTTPClient = hc
	}
}

// GenerateSign creates the MD5 signature for a QERP API request.
// Parameters are sorted by name (ASCII ascending), concatenated as
// key1=value1key2=value2..., then appSecret is appended and MD5 is computed.
func (c *Client) GenerateSign(serviceType, bizParam string, timestamp int64) string {
	params := map[string]string{
		"appId":       c.AppID,
		"bizParam":    bizParam,
		"serviceType": serviceType,
		"timestamp":   fmt.Sprintf("%d", timestamp),
	}
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var sb strings.Builder
	for _, k := range keys {
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(params[k])
	}
	sb.WriteString(c.AppSecret)
	h := md5.Sum([]byte(sb.String()))
	return hex.EncodeToString(h[:])
}

// Do sends an API request to the appropriate endpoint based on serviceType.
// It generates the signature, builds the multipart/form-data body, and
// parses the response into the provided result.
func (c *Client) Do(serviceType, bizParam string, result any) error {
	ts := time.Now().UnixMilli()
	sign := c.GenerateSign(serviceType, bizParam, ts)

	body := &strings.Builder{}
	w := multipart.NewWriter(body)
	_ = w.WriteField("appId", c.AppID)
	_ = w.WriteField("serviceType", serviceType)
	_ = w.WriteField("bizParam", bizParam)
	_ = w.WriteField("timestamp", fmt.Sprintf("%d", ts))
	_ = w.WriteField("sign", sign)
	if err := w.Close(); err != nil {
		return fmt.Errorf("close multipart: %w", err)
	}

	endpoint := endpointForService(serviceType)
	reqURL := fmt.Sprintf("%s/api/v1/%s", c.BaseURL, endpoint)
	req, err := http.NewRequest(http.MethodPost, reqURL, strings.NewReader(body.String()))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}
	return parseResponse(respBody, result)
}

// endpointForService maps each serviceType to its API endpoint path.
// Verified against https://open.qianyierp.com documentation.
func endpointForService(serviceType string) string {
	m := map[string]string{
		// Shop
		"QUERY_SHOP_LIST": "shop",
		// Warehouse
		"QUERY_WAREHOUSE_LIST": "warehouse",
		// SKU
		"QUERY_SIMPLE_LIST_SKU": "sku",
		"INSERT_SKU_INFO":       "sku",
		"UPDATE_SKU_INFO":       "sku",
		"ENABLE_SKU":            "sku",
		"QUERY_SYS_SKU":         "sku",
		// Sales Order
		"CREATE_SALES_ORDER":                "salesOrder",
		"CLOSE_SALES_ORDER":                 "salesOrder",
		"QUERY_SALES_ORDER_LIST":            "salesOrder",
		"QUERY_SALES_ORDER_NUMBER_LIST":     "salesOrder",
		"QUERY_SALES_ORDER_SHIPPING_INFO":   "salesOrder",
		"QUERY_SALES_ORDER_AUDIT":           "salesOrder",
		"CREATE_WAVE_ORDER":                 "salesOrder",
		"SEND_SALES_ORDER_TO_WMS":           "salesOrder",
		"QUERY_ORIGINAL_SALES_ORDER":        "salesOrder",
		"QUERY_SALES_ORDER_PICKUP_STATUS":   "salesOrder",
		"QUERY_SALES_ORDER_DOCUMENT":        "salesOrder",
		// Return/Refund
		"CREATE_RETURN_ORDER":   "returnOrder",
		"CLOSE_RETURN_ORDER":    "returnOrder",
		"QUERY_RETURN_ORDER_LIST": "returnOrder",
		"PUSH_RETURN_ORDER_INFO": "returnOrder",
		// Inventory
		"QUERY_SIMPLE_LIST_INVENTORY":    "inventory",
		"QUERY_SIMPLE_LIST_INVENTORY_V2": "inventory",
		"QUERY_INVENTORY_LOG_LIST":       "inventory",
		"QUERY_INVENTORY_ASSEMBLY_LIST":  "inventory",
		"CREATE_TRANSFER_ORDER":          "inventory",
		"QUERY_TRANSFER_ORDER_LIST":      "inventory",
		"QUERY_SPLIT_ORDER_LIST":         "inventory",
		"QUERY_STORAGE_LOC_INVENTORY":    "inventory",
		"QUERY_BATCH_INVENTORY_LIST":     "inventory",
		"TRANSFER_STORAGE_LOCATION":      "inventory",
		"QUERY_SBS_INVENTORY_LIST":       "inventory",
		"QUERY_SBS_WAREHOUSE_LIST":       "inventory",
		// ASN (Inbound)
		"CREATE_ASN_ORDER":    "asn",
		"QUERY_ASN_LIST":      "asn",
		"CANCEL_ASN_ORDER":    "asn",
		"DELETE_ASN_ORDER":    "asn",
		"PUSH_ASN_ORDER":      "asn",
		"QUERY_ASN_BATCH_LIST": "asn",
		// ODO (Outbound)
		"QUERY_ODO_LIST":        "odo",
		"QUERY_SALES_ODO_LIST":  "odo",
		"CREATE_ODO_ORDER":      "odo",
		"CANCEL_ODO_ORDER":      "odo",
		"PUSH_ODO_ORDER":        "odo",
		// Adjustment
		"QUERY_ADJUSTMENT_LIST":   "adjustment",
		"CREATE_ADJUSTMENT_ORDER": "adjustment",
		// Purchase
		"QUERY_PURCHASE_ORDER_LIST": "purchase",
		"CREATE_PURCHASE_ORDER":     "purchase",
		"UPDATE_PURCHASE_ORDER":     "purchase",
		// First Leg (Logistics)
		"QUERY_FIRST_LEG_ORDER_LIST":       "firstLeg",
		"CREATE_FIRST_LEG_ORDER":           "firstLeg",
		"QUERY_FIRST_LRG_LOGISTICS":       "firstLeg",
		"QUERY_FIRST_LRG_TRACKING_PACKAGE": "firstLeg",
		"WITHDRAW_AND_DEL_FIRST_LEG":       "firstLeg",
		// Report
		"QUERY_SHOPEE_TRANSACTION_DETAIL_LIST": "report",
		"QUERY_LAZADA_TRANSACTION_DETAIL_LIST": "report",
		"QUERY_TIKTOK_TRANSACTION_DETAIL_LIST": "report",
		"QUERY_SHOPEE_PAYOUT_DETAIL_LIST":      "report",
		"QUERY_LAZADA_ACCOUNT_TRANSACTION_LIST":  "report",
		"QUERY_TIKTOK_V2_TRANSACTION_DETAIL_LIST": "report",
		"QUERY_TIKTOK_PAYOUT_RECORD":           "report",
		"QUERY_INVENTORY_DAILY_REPORT":         "report",
		// Webhook / Subscribe
		"SUBSCRIBE_ORDER":               "salesOrder",
		// Customer Field
		"CUSTOMER_FIELD_QUERY": "property",
	}
	if ep, ok := m[serviceType]; ok {
		return ep
	}
	return ""
}
