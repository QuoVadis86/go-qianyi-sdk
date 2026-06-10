package qianyi

import (
	"bytes"
	"context"
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

const (
	defaultBaseURL = "https://www.qianyierp.com"
	defaultTimeout = 30 * time.Second
	userAgent      = "go-qianyi-sdk/1.0"
)

// HTTPClient is the interface for making HTTP requests.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var defaultClient HTTPClient = &http.Client{Timeout: defaultTimeout}

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
		BaseURL:    defaultBaseURL,
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
func (c *Client) Do(ctx context.Context, serviceType, bizParam string, result any) error {
	ts := time.Now().UnixMilli()
	sign := c.GenerateSign(serviceType, bizParam, ts)

	var body bytes.Buffer
	w := multipart.NewWriter(&body)
	if err := w.WriteField("appId", c.AppID); err != nil {
		return fmt.Errorf("write appId field: %w", err)
	}
	if err := w.WriteField("serviceType", serviceType); err != nil {
		return fmt.Errorf("write serviceType field: %w", err)
	}
	if err := w.WriteField("bizParam", bizParam); err != nil {
		return fmt.Errorf("write bizParam field: %w", err)
	}
	if err := w.WriteField("timestamp", fmt.Sprintf("%d", ts)); err != nil {
		return fmt.Errorf("write timestamp field: %w", err)
	}
	if err := w.WriteField("sign", sign); err != nil {
		return fmt.Errorf("write sign field: %w", err)
	}
	if err := w.Close(); err != nil {
		return fmt.Errorf("close multipart writer: %w", err)
	}

	endpoint := endpointForService(serviceType)
	reqURL := fmt.Sprintf("%s/api/v1/%s", c.BaseURL, endpoint)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, &body)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("unexpected HTTP status %d: %s", resp.StatusCode, truncate(string(respBody), 200))
	}
	return parseResponse(respBody, result)
}

// endpointForService maps each serviceType to its API endpoint path.
// Verified against https://open.qianyierp.com documentation.
func endpointForService(serviceType string) string {
	m := map[string]string{
		ServiceTypeQueryShopList:                   "shop",
		ServiceTypeQueryWarehouseList:              "warehouse",
		ServiceTypeQuerySimpleListSku:              "sku",
		ServiceTypeInsertSkuInfo:                   "sku",
		ServiceTypeUpdateSkuInfo:                   "sku",
		ServiceTypeEnableSku:                       "sku",
		ServiceTypeQuerySysSku:                     "sku",
		ServiceTypeCreateSalesOrder:                "salesOrder",
		ServiceTypeCloseSalesOrder:                 "salesOrder",
		ServiceTypeQuerySalesOrderList:             "salesOrder",
		ServiceTypeQuerySalesOrderNumberList:       "salesOrder",
		ServiceTypeQuerySalesOrderShippingInfo:     "salesOrder",
		ServiceTypeQuerySalesOrderAudit:            "salesOrder",
		ServiceTypeCreateWaveOrder:                 "salesOrder",
		ServiceTypeSendSalesOrderToWms:             "salesOrder",
		ServiceTypeQueryOriginalSalesOrder:         "salesOrder",
		ServiceTypeQuerySalesOrderPickupStatus:     "salesOrder",
		ServiceTypeQuerySalesOrderDocument:         "salesOrder",
		ServiceTypeCreateReturnOrder:               "returnOrder",
		ServiceTypeCloseReturnOrder:                "returnOrder",
		ServiceTypeQueryReturnOrderList:            "returnOrder",
		ServiceTypePushReturnOrderInfo:             "returnOrder",
		ServiceTypeQuerySimpleListInventory:        "inventory",
		ServiceTypeQuerySimpleListInventoryV2:      "inventory",
		ServiceTypeQueryInventoryLogList:           "inventory",
		ServiceTypeQueryInventoryAssemblyList:      "inventory",
		ServiceTypeCreateTransferOrder:             "inventory",
		ServiceTypeQueryTransferOrderList:          "inventory",
		ServiceTypeQuerySplitOrderList:             "inventory",
		ServiceTypeQueryStorageLocInventory:        "inventory",
		ServiceTypeQueryBatchInventoryList:         "inventory",
		ServiceTypeTransferStorageLocation:         "inventory",
		ServiceTypeQuerySbsInventoryList:           "inventory",
		ServiceTypeQuerySbsWarehouseList:           "inventory",
		ServiceTypeCreateAsnOrder:                  "asn",
		ServiceTypeQueryAsnList:                    "asn",
		ServiceTypeCancelAsnOrder:                  "asn",
		ServiceTypeDeleteAsnOrder:                  "asn",
		ServiceTypePushAsnOrder:                    "asn",
		ServiceTypeQueryAsnBatchList:               "asn",
		ServiceTypeQueryOdoList:                    "odo",
		ServiceTypeQuerySalesOdoList:               "odo",
		ServiceTypeCreateOdoOrder:                  "odo",
		ServiceTypeCancelOdoOrder:                  "odo",
		ServiceTypePushOdoOrder:                    "odo",
		ServiceTypeQueryAdjustmentList:             "adjustment",
		ServiceTypeCreateAdjustmentOrder:           "adjustment",
		ServiceTypeQueryPurchaseOrderList:          "purchase",
		ServiceTypeCreatePurchaseOrder:             "purchase",
		ServiceTypeQueryFirstLegOrderList:          "firstLeg",
		ServiceTypeCreateFirstLegOrder:             "firstLeg",
		ServiceTypeQueryFirstLrgLogistics:          "firstLeg",
		ServiceTypeQueryFirstLrgTrackingPackage:    "firstLeg",
		ServiceTypeWithdrawAndDelFirstLeg:          "firstLeg",
		ServiceTypeQueryShopeeTransactionDetailList:    "report",
		ServiceTypeQueryLazadaTransactionDetailList:    "report",
		ServiceTypeQueryTiktokTransactionDetailList:    "report",
		ServiceTypeQueryShopeePayoutDetailList:         "report",
		ServiceTypeQueryLazadaAccountTransactionList:   "report",
		ServiceTypeQueryTiktokV2TransactionDetailList:  "report",
		ServiceTypeQueryTiktokPayoutRecord:             "report",
		ServiceTypeQueryInventoryDailyReport:           "report",
		ServiceTypeSubscribeOrder:                  "salesOrder",
		ServiceTypeCustomerFieldQuery:              "property",
		ServiceTypeQuerySupplierList:               "supplier",
		ServiceTypeCreateSupplier:                  "supplier",
		ServiceTypeQuerySupplierSkuList:            "supplier",
		ServiceTypeCreateSupplierSku:               "supplier",
	}
	if ep, ok := m[serviceType]; ok {
		return ep
	}
	return ""
}
