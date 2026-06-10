package qianyi

import "encoding/json"

// ReportService provides access to financial report API operations.
type ReportService struct {
	client *Client
}

// NewReportService creates a new ReportService.
func NewReportService(client *Client) *ReportService {
	return &ReportService{client: client}
}

// ReportPageParams holds common pagination + shop filter params.
type ReportPageParams struct {
	Page         int      `json:"page"`
	PageSize     int      `json:"pageSize"`
	ShopIDList   []int64  `json:"shopIdList,omitempty"`
	ShopNameList []string `json:"shopNameList,omitempty"`
}

// ShopeeReportQuery extends ReportPageParams with Shopee-specific filters.
type ShopeeReportQuery struct {
	Page               int      `json:"page"`
	PageSize           int      `json:"pageSize"`
	ShopIDList         []int64  `json:"shopIdList,omitempty"`
	ShopNameList       []string `json:"shopNameList,omitempty"`
	Type               int      `json:"type"`
	AdjustmentLevel    string   `json:"adjustmentLevel,omitempty"`
	PayoutTimeFrom     string   `json:"payoutTimeFrom,omitempty"`
	PayoutTimeTo       string   `json:"payoutTimeTo,omitempty"`
	UpdateTimeFrom     string   `json:"updateTimeFrom,omitempty"`
	UpdateTimeTo       string   `json:"updateTimeTo,omitempty"`
	OnlineOrderID      string   `json:"onlineOrderId,omitempty"`
	OnlineOrderIDList  []string `json:"onlineOrderIdList,omitempty"`
}

// ShopeeReportDTO represents Shopee transaction details.
type ShopeeReportDTO struct {
	ShopID            int64   `json:"shopId"`
	OnlineShopID      string  `json:"onlineShopId"`
	ShopName          string  `json:"shopName"`
	OrderSN            string  `json:"ordersn"`
	PayoutTime        int64   `json:"payoutTime"`
	PayoutTimeFormatted string `json:"payoutTimeFormatted"`
	Currency          string  `json:"currency"`
	AdjustmentLevel   string  `json:"adjustmentLevel,omitempty"`
	OrderStatus       string  `json:"orderStatus,omitempty"`
	ReturnSN          string  `json:"returnsn,omitempty"`
	RefundSN          string  `json:"refundsn,omitempty"`
	EscrowAmount      float64 `json:"escrowAmount,omitempty"`
	OriginalPrice     float64 `json:"originalPrice,omitempty"`
	OrderSellingPrice float64 `json:"orderSellingPrice,omitempty"`
	SellerDiscount    float64 `json:"sellerDiscount,omitempty"`
	ProductSaleAmount float64 `json:"productSaleAmount,omitempty"`
	OriginalPriceBeforeRefund float64 `json:"originalPriceBeforeRefund,omitempty"`
	SellingPriceBeforeRefund  float64 `json:"sellingPriceBeforeRefund,omitempty"`
	SellerDiscountBeforeRefund float64 `json:"sellerDiscountBeforeRefund,omitempty"`
	RefundAmount      float64 `json:"refundAmount,omitempty"`
	ShopeeDiscount    float64 `json:"shopeeDiscount,omitempty"`
	VoucherFromSeller float64 `json:"voucherFromSeller,omitempty"`
	SellerCoinCashBack float64 `json:"sellerCoinCashBack,omitempty"`
	BuyerPaidShippingFee float64 `json:"buyerPaidShippingFee,omitempty"`
	ShopeeShippingRebate float64 `json:"shopeeShippingRebate,omitempty"`
	ActualShippingFee float64 `json:"actualShippingFee,omitempty"`
	ReverseShippingFee float64 `json:"reverseShippingFee,omitempty"`
	CommissionFee     float64 `json:"commissionFee,omitempty"`
	ServiceFee        float64 `json:"serviceFee,omitempty"`
	SellerTransactionFee float64 `json:"sellerTransactionFee,omitempty"`
	EscrowTax         float64 `json:"escrowTax,omitempty"`
	SellerLostCompensation float64 `json:"sellerLostCompensation,omitempty"`
	BuyerTotalAmount  float64 `json:"buyerTotalAmount,omitempty"`
	Coins             float64 `json:"coins,omitempty"`
	VoucherFromShopee float64 `json:"voucherFromShopee,omitempty"`
	UpdateTime        int64   `json:"updateTime,omitempty"`
	ShopTimeZone      string  `json:"shopTimeZone,omitempty"`
	PublicID          string  `json:"publicId,omitempty"`
}

// QueryShopeeTransaction queries Shopee transaction details.
func (s *ReportService) QueryShopeeTransaction(params *ShopeeReportQuery) ([]ShopeeReportDTO, int, error) {
	biz, _ := json.Marshal(params)
	var list []ShopeeReportDTO
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_SHOPEE_TRANSACTION_DETAIL_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// LazadaReportQuery extends ReportPageParams with Lazada-specific filters.
type LazadaReportQuery struct {
	Page             int      `json:"page"`
	PageSize         int      `json:"pageSize"`
	ShopIDList       []int64  `json:"shopIdList,omitempty"`
	ShopNameList     []string `json:"shopNameList,omitempty"`
	PayoutTimeFrom   string   `json:"payoutTimeFrom,omitempty"`
	PayoutTimeTo     string   `json:"payoutTimeTo,omitempty"`
	UpdateTimeFrom   string   `json:"updateTimeFrom,omitempty"`
	UpdateTimeTo     string   `json:"updateTimeTo,omitempty"`
	OnlineOrderID    string   `json:"onlineOrderId,omitempty"`
	OnlineOrderIDList []string `json:"onlineOrderIdList,omitempty"`
	FeeName          string   `json:"feeName,omitempty"`
}

// LazadaReportDTO represents Lazada transaction details.
type LazadaReportDTO struct {
	ShopID             int64  `json:"shopId"`
	OnlineShopID       string `json:"onlineShopId"`
	ShopName           string `json:"shopName"`
	TransactionDate    string `json:"transactionDate"`
	TransactionTimestamp int64 `json:"transactionTimestamp"`
	TransactionNumber  string `json:"transactionNumber"`
	TransactionType    string `json:"transactionType"`
	OrderNo            string `json:"orderNo"`
	FeeName            string `json:"feeName,omitempty"`
	Currency           string `json:"currency,omitempty"`
	Amount             string `json:"amount,omitempty"`
	Details            string `json:"details,omitempty"`
	SellerSku          string `json:"sellerSku,omitempty"`
	VatInAmount        string `json:"vatInAmount,omitempty"`
	WhtAmount          string `json:"whtAmount,omitempty"`
	Comment            string `json:"comment,omitempty"`
	UpdateTime         int64  `json:"updateTime,omitempty"`
	ShopTimeZone       string `json:"shopTimeZone,omitempty"`
	PublicID           string `json:"publicId,omitempty"`
}

// QueryLazadaTransaction queries Lazada transaction details.
func (s *ReportService) QueryLazadaTransaction(params *LazadaReportQuery) ([]LazadaReportDTO, int, error) {
	biz, _ := json.Marshal(params)
	var list []LazadaReportDTO
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_LAZADA_TRANSACTION_DETAIL_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// TiktokReportQuery extends ReportPageParams with TikTok-specific filters.
type TiktokReportQuery struct {
	Page             int      `json:"page"`
	PageSize         int      `json:"pageSize"`
	ShopIDList       []int64  `json:"shopIdList,omitempty"`
	ShopNameList     []string `json:"shopNameList,omitempty"`
	PayoutTimeFrom   string   `json:"payoutTimeFrom"`
	PayoutTimeTo     string   `json:"payoutTimeTo"`
	OnlineOrderID    string   `json:"onlineOrderId,omitempty"`
	OnlineOrderIDList []string `json:"onlineOrderIdList,omitempty"`
}

// TiktokReportDTO represents TikTok transaction details.
type TiktokReportDTO struct {
	ShopID          int64   `json:"shopId"`
	ShopName        string  `json:"shopName"`
	Currency        string  `json:"currency"`
	OrderID         string  `json:"orderId"`
	AdjustmentID    string  `json:"adjustmentId,omitempty"`
	RelatedOrderID  string  `json:"relatedOrderId,omitempty"`
	SettlementTime  int64   `json:"settlementTime,omitempty"`
	SkuID           string  `json:"skuId,omitempty"`
	SkuName         string  `json:"skuName,omitempty"`
	ProductName     string  `json:"productName,omitempty"`
	UserPay         float64 `json:"userPay,omitempty"`
	PlatformPromotion float64 `json:"platformPromotion,omitempty"`
	ShippingFeeSubsidy float64 `json:"shippingFeeSubsidy,omitempty"`
	Refund          float64 `json:"refund,omitempty"`
	PaymentFee      float64 `json:"paymentFee,omitempty"`
	PlatformCommission float64 `json:"platformCommission,omitempty"`
	FlatFee         float64 `json:"flatFee,omitempty"`
	SalesFee        float64 `json:"salesFee,omitempty"`
	AffiliateCommission float64 `json:"affiliateCommission,omitempty"`
	Vat             float64 `json:"vat,omitempty"`
	ShippingFee     float64 `json:"shippingFee,omitempty"`
	SettlementAmount float64 `json:"settlementAmount,omitempty"`
	UpdateTime      int64   `json:"updateTime,omitempty"`
	ShopTimeZone    string  `json:"shopTimeZone,omitempty"`
	PublicID        string  `json:"publicId,omitempty"`
}

// QueryTiktokTransaction queries TikTok transaction details.
func (s *ReportService) QueryTiktokTransaction(params *TiktokReportQuery) ([]TiktokReportDTO, int, error) {
	biz, _ := json.Marshal(params)
	var list []TiktokReportDTO
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_TIKTOK_TRANSACTION_DETAIL_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// ShopeePayoutQuery extends ReportPageParams with Shopee payout filters.
type ShopeePayoutQuery struct {
	Page           int      `json:"page"`
	PageSize       int      `json:"pageSize"`
	ShopIDList     []int64  `json:"shopIdList,omitempty"`
	ShopNameList   []string `json:"shopNameList,omitempty"`
	PayoutTimeFrom string   `json:"payoutTimeFrom"`
	PayoutTimeTo   string   `json:"payoutTimeTo"`
}

// QueryShopeePayout queries Shopee payout records.
func (s *ReportService) QueryShopeePayout(params *ShopeePayoutQuery) ([]any, int, error) {
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_SHOPEE_PAYOUT_DETAIL_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// QueryLazadaBalance queries Lazada my balance.
func (s *ReportService) QueryLazadaBalance(params *ReportPageParams) ([]any, int, error) {
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_LAZADA_MY_BALANCE", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// QueryTiktokV2Transaction queries TikTok V2 transaction details.
func (s *ReportService) QueryTiktokV2Transaction(params *TiktokReportQuery) ([]any, int, error) {
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_TIKTOK_V2_TRANSACTION_DETAIL_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// QueryTiktokPayout queries TikTok payout records.
func (s *ReportService) QueryTiktokPayout(params *ReportPageParams) ([]any, int, error) {
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_TIKTOK_PAYOUT_RECORD", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// InventoryDailyReportQuery holds params for inventory daily report.
type InventoryDailyReportQuery struct {
	Page       int      `json:"page"`
	PageSize   int      `json:"pageSize"`
	DateFrom   string   `json:"dateFrom"`
	DateTo     string   `json:"dateTo"`
	ShopIDList []int64  `json:"shopIdList,omitempty"`
}

// QueryInventoryDailyReport queries inventory daily statement.
func (s *ReportService) QueryInventoryDailyReport(params *InventoryDailyReportQuery) ([]any, int, error) {
	biz, _ := json.Marshal(params)
	var list []any
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_INVENTORY_DAILY_REPORT", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}
