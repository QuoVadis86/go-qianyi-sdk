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

// ReportPageParams holds common pagination and shop filter params.
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
	ShopID                     int64   `json:"shopId"`
	OnlineShopID               string  `json:"onlineShopId"`
	ShopName                   string  `json:"shopName"`
	OrderSN                    string  `json:"ordersn"`
	PayoutTime                 int64   `json:"payoutTime"`
	PayoutTimeFormatted        string  `json:"payoutTimeFormatted"`
	Currency                   string  `json:"currency"`
	AdjustmentLevel            string  `json:"adjustmentLevel,omitempty"`
	OrderStatus                string  `json:"orderStatus,omitempty"`
	ReturnSN                   string  `json:"returnsn,omitempty"`
	RefundSN                   string  `json:"refundsn,omitempty"`
	EscrowAmount               float64 `json:"escrowAmount,omitempty"`
	OriginalPrice              float64 `json:"originalPrice,omitempty"`
	OrderSellingPrice          float64 `json:"orderSellingPrice,omitempty"`
	SellerDiscount             float64 `json:"sellerDiscount,omitempty"`
	ProductSaleAmount          float64 `json:"productSaleAmount,omitempty"`
	OriginalPriceBeforeRefund  float64 `json:"originalPriceBeforeRefund,omitempty"`
	SellingPriceBeforeRefund   float64 `json:"sellingPriceBeforeRefund,omitempty"`
	SellerDiscountBeforeRefund float64 `json:"sellerDiscountBeforeRefund,omitempty"`
	RefundAmount               float64 `json:"refundAmount,omitempty"`
	ShopeeDiscount             float64 `json:"shopeeDiscount,omitempty"`
	VoucherFromSeller          float64 `json:"voucherFromSeller,omitempty"`
	SellerCoinCashBack         float64 `json:"sellerCoinCashBack,omitempty"`
	BuyerPaidShippingFee       float64 `json:"buyerPaidShippingFee,omitempty"`
	ShopeeShippingRebate       float64 `json:"shopeeShippingRebate,omitempty"`
	ActualShippingFee          float64 `json:"actualShippingFee,omitempty"`
	ReverseShippingFee         float64 `json:"reverseShippingFee,omitempty"`
	CommissionFee              float64 `json:"commissionFee,omitempty"`
	ServiceFee                 float64 `json:"serviceFee,omitempty"`
	SellerTransactionFee       float64 `json:"sellerTransactionFee,omitempty"`
	EscrowTax                  float64 `json:"escrowTax,omitempty"`
	SellerLostCompensation     float64 `json:"sellerLostCompensation,omitempty"`
	BuyerTotalAmount           float64 `json:"buyerTotalAmount,omitempty"`
	Coins                      float64 `json:"coins,omitempty"`
	VoucherFromShopee          float64 `json:"voucherFromShopee,omitempty"`
	UpdateTime                 int64   `json:"updateTime,omitempty"`
	ShopTimeZone               string  `json:"shopTimeZone,omitempty"`
	PublicID                   string  `json:"publicId,omitempty"`
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

// TiktokReportDTO represents TikTok transaction details.
type TiktokReportDTO struct {
	ShopID          int64   `json:"shopId"`
	ShopName        string  `json:"shopName"`
	Currency        string  `json:"currency"`
	OrderID         string  `json:"orderId"`
	AdjustmentID    string  `json:"adjustmentId,omitempty"`
	RelatedOrderID  string  `json:"relatedOrderId,omitempty"`
	SettlementTime  int64   `json:"settlementTime,omitempty"`
	SettlementTimeFmt string `json:"settlementTimeFormatted,omitempty"`
	SkuID           string  `json:"skuId,omitempty"`
	SkuName         string  `json:"skuName,omitempty"`
	ProductName     string  `json:"productName,omitempty"`
	UserPay         float64 `json:"userPay,omitempty"`
	PlatformPromotion float64 `json:"platformPromotion,omitempty"`
	ShippingFeeSubsidy float64 `json:"shippingFeeSubsidy,omitempty"`
	Refund          float64 `json:"refund,omitempty"`
	PaymentFee      float64 `json:"paymentFee,omitempty"`
	PlatformCommission float64 `json:"platformCommission,omitempty"`
	AffiliateCommission float64 `json:"affiliateCommission,omitempty"`
	Vat             float64 `json:"vat,omitempty"`
	ShippingFee     float64 `json:"shippingFee,omitempty"`
	SettlementAmount float64 `json:"settlementAmount,omitempty"`
	UpdateTime      int64   `json:"updateTime,omitempty"`
	ShopTimeZone    string  `json:"shopTimeZone,omitempty"`
	PublicID        string  `json:"publicId,omitempty"`
}

// LazadaReportQuery holds params for Lazada transaction details.
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

// TiktokReportQuery holds params for TikTok transaction details.
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

// ----- Shopee Payout -----

// ShopeePayoutQuery holds params for Shopee payout records.
type ShopeePayoutQuery struct {
	Page              int      `json:"page"`
	PageSize          int      `json:"pageSize"`
	ShopIDList        []int64  `json:"shopIdList,omitempty"`
	ShopNameList      []string `json:"shopNameList,omitempty"`
	PayoutTimeFrom    string   `json:"payoutTimeFrom"`
	PayoutTimeTo      string   `json:"payoutTimeTo"`
	UpdateTimeFrom    string   `json:"updateTimeFrom,omitempty"`
	UpdateTimeTo      string   `json:"updateTimeTo,omitempty"`
	Type              string   `json:"type,omitempty"`
	PayoutStatus      string   `json:"payoutStatus,omitempty"`
	TransactionType   string   `json:"transactionType,omitempty"`
	OnlineOrderID     string   `json:"onlineOrderId,omitempty"`
	OnlineOrderIDList []string `json:"onlineOrderIdList,omitempty"`
}

// ShopeePayoutDTO represents a Shopee payout record.
type ShopeePayoutDTO struct {
	ShopID               int64   `json:"shopId,omitempty"`
	OnlineShopID         string  `json:"onlineShopId,omitempty"`
	ShopName             string  `json:"shopName,omitempty"`
	Platform             string  `json:"platform,omitempty"`
	SiteCode             string  `json:"siteCode,omitempty"`
	Currency             string  `json:"currency,omitempty"`
	Type                 string  `json:"type,omitempty"`
	PayType              string  `json:"payType,omitempty"`
	TransactionID        string  `json:"transactionId,omitempty"`
	PayoutStatus         string  `json:"payoutStatus,omitempty"`
	TransactionType      string  `json:"transactionType,omitempty"`
	Amount               float64 `json:"amount,omitempty"`
	CurrentBalance       float64 `json:"currentBalance,omitempty"`
	TransactionTime      int64   `json:"transactionTime,omitempty"`
	TransactionTimeFmt   string  `json:"transactionTimeFormatted,omitempty"`
	UpdateTime           int64   `json:"updateTime,omitempty"`
	OrderSN              string  `json:"orderSn,omitempty"`
	RefundSN             string  `json:"refundSn,omitempty"`
	WithdrawalType       string  `json:"withdrawalType,omitempty"`
	TransactionFee       float64 `json:"transactionFee,omitempty"`
	Description          string  `json:"description,omitempty"`
	BuyerName            string  `json:"buyerName,omitempty"`
	WithdrawID           string  `json:"withdrawId,omitempty"`
	Reason               string  `json:"reason,omitempty"`
	RootWithdrawalID     string  `json:"rootWithdrawalId,omitempty"`
	PayoutExchangeCurrency string `json:"payoutExchangeCurrency,omitempty"`
	PayoutExchangeAmount float64 `json:"payoutExchangeAmount,omitempty"`
	PayoutExchangeRate   string  `json:"payoutExchangeRate,omitempty"`
	PayoutPayeeID        string  `json:"payoutPayeeId,omitempty"`
	ShopTimeZone         string  `json:"shopTimeZone,omitempty"`
	PublicID             string  `json:"publicId,omitempty"`
}

// QueryShopeePayout queries Shopee payout records.
func (s *ReportService) QueryShopeePayout(params *ShopeePayoutQuery) ([]ShopeePayoutDTO, int, error) {
	biz, _ := json.Marshal(params)
	var list []ShopeePayoutDTO
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_SHOPEE_PAYOUT_DETAIL_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// ----- Lazada Account Transaction -----

// LazadaAccountQuery holds params for Lazada account transaction list.
type LazadaAccountQuery struct {
	Page                int      `json:"page"`
	PageSize            int      `json:"pageSize"`
	ShopIDList          []int64  `json:"shopIdList,omitempty"`
	ShopNameList        []string `json:"shopNameList,omitempty"`
	PayoutTimeFrom      string   `json:"payoutTimeFrom"`
	PayoutTimeTo        string   `json:"payoutTimeTo"`
	UpdateTimeFrom      string   `json:"updateTimeFrom,omitempty"`
	UpdateTimeTo        string   `json:"updateTimeTo,omitempty"`
	TransactionNumberList []string `json:"transactionNumberList,omitempty"`
	TypeList            []string `json:"typeList,omitempty"`
}

// LazadaAccountTransactionDTO represents a Lazada account transaction record.
type LazadaAccountTransactionDTO struct {
	CustomerID       int64  `json:"customerId,omitempty"`
	ShopID           int64  `json:"shopId,omitempty"`
	OnlineShopID     string `json:"onlineShopId,omitempty"`
	SiteCode         string `json:"siteCode,omitempty"`
	TransactionNumber string `json:"transactionNumber,omitempty"`
	TransactionTime  int64  `json:"transactionTime,omitempty"`
	TransactionTimeFmt string `json:"transactionTimeFormatted,omitempty"`
	Type             string `json:"type,omitempty"`
	SubType          string `json:"subType,omitempty"`
	Amount           float64 `json:"amount,omitempty"`
	Currency         string `json:"currency,omitempty"`
	Remarks          string `json:"remarks,omitempty"`
	PayeeAccount     string `json:"payeeAccount,omitempty"`
	PayeeDescription string `json:"payeeDescription,omitempty"`
	ShopTimeZone     string `json:"shopTimeZone,omitempty"`
	PublicID         string `json:"publicId,omitempty"`
}

// QueryLazadaAccountTransaction queries Lazada account transaction list.
func (s *ReportService) QueryLazadaAccountTransaction(params *LazadaAccountQuery) ([]LazadaAccountTransactionDTO, int, error) {
	biz, _ := json.Marshal(params)
	var list []LazadaAccountTransactionDTO
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_LAZADA_ACCOUNT_TRANSACTION_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// ----- TikTok V2 Transaction -----

// TiktokV2ReportQuery holds params for TikTok V2 transaction details.
type TiktokV2ReportQuery struct {
	Page             int      `json:"page"`
	PageSize         int      `json:"pageSize"`
	ShopIDList       []int64  `json:"shopIdList,omitempty"`
	ShopNameList     []string `json:"shopNameList,omitempty"`
	PayoutTimeFrom   string   `json:"payoutTimeFrom"`
	PayoutTimeTo     string   `json:"payoutTimeTo"`
	OnlineOrderIDList []string `json:"onlineOrderIdList,omitempty"`
}

// TiktokV2ReportDTO represents a TikTok V2 transaction detail.
type TiktokV2ReportDTO struct {
	ShopID          int64   `json:"shopId,omitempty"`
	ShopName        string  `json:"shopName,omitempty"`
	Currency        string  `json:"currency,omitempty"`
	OrderID         string  `json:"orderId,omitempty"`
	OrderType       string  `json:"orderType,omitempty"`
	OrderStatus     string  `json:"orderStatus,omitempty"`
	SettlementTime  int64   `json:"settlementTime,omitempty"`
	ProductAmount   float64 `json:"productAmount,omitempty"`
	ShippingAmount  float64 `json:"shippingAmount,omitempty"`
	PlatformFee     float64 `json:"platformFee,omitempty"`
	TransactionFee  float64 `json:"transactionFee,omitempty"`
	Commission      float64 `json:"commission,omitempty"`
	AffiliateFee    float64 `json:"affiliateFee,omitempty"`
	Refund          float64 `json:"refund,omitempty"`
	NetAmount       float64 `json:"netAmount,omitempty"`
	SettlementAmount float64 `json:"settlementAmount,omitempty"`
	UpdateTime      int64   `json:"updateTime,omitempty"`
	ShopTimeZone    string  `json:"shopTimeZone,omitempty"`
	PublicID        string  `json:"publicId,omitempty"`
}

// QueryTiktokV2Transaction queries TikTok V2 transaction details.
func (s *ReportService) QueryTiktokV2Transaction(params *TiktokV2ReportQuery) ([]TiktokV2ReportDTO, int, error) {
	biz, _ := json.Marshal(params)
	var list []TiktokV2ReportDTO
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_TIKTOK_V2_TRANSACTION_DETAIL_LIST", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// ----- TikTok Payout -----

// TiktokPayoutQuery holds params for TikTok payout records.
type TiktokPayoutQuery struct {
	Page             int      `json:"page"`
	PageSize         int      `json:"pageSize"`
	ShopIDList       []int64  `json:"shopIdList,omitempty"`
	ShopNameList     []string `json:"shopNameList,omitempty"`
	PayoutTimeFrom   string   `json:"payoutTimeFrom"`
	PayoutTimeTo     string   `json:"payoutTimeTo"`
}

// TiktokPayoutDTO represents a TikTok payout record.
type TiktokPayoutDTO struct {
	ShopID            int64   `json:"shopId,omitempty"`
	OnlineShopID      string  `json:"onlineShopId,omitempty"`
	ShopName          string  `json:"shopName,omitempty"`
	Currency          string  `json:"currency,omitempty"`
	PayoutTime        int64   `json:"payoutTime,omitempty"`
	PayoutTimeFmt     string  `json:"payoutTimeFormatted,omitempty"`
	TransactionID     string  `json:"transactionId,omitempty"`
	TransactionType   string  `json:"transactionType,omitempty"`
	Amount            float64 `json:"amount,omitempty"`
	Balance           float64 `json:"balance,omitempty"`
	Status            string  `json:"status,omitempty"`
	Reference         string  `json:"reference,omitempty"`
	Remark            string  `json:"remark,omitempty"`
	UpdateTime        int64   `json:"updateTime,omitempty"`
	ShopTimeZone      string  `json:"shopTimeZone,omitempty"`
	PublicID          string  `json:"publicId,omitempty"`
}

// QueryTiktokPayout queries TikTok payout records.
func (s *ReportService) QueryTiktokPayout(params *TiktokPayoutQuery) ([]TiktokPayoutDTO, int, error) {
	biz, _ := json.Marshal(params)
	var list []TiktokPayoutDTO
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_TIKTOK_PAYOUT_RECORD", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}

// ----- Inventory Daily Report -----

// InventoryDailyReportQuery holds params for inventory daily statement.
type InventoryDailyReportQuery struct {
	Page       int      `json:"page"`
	PageSize   int      `json:"pageSize"`
	DateFrom   string   `json:"dateFrom"`
	DateTo     string   `json:"dateTo"`
	ShopIDList []int64  `json:"shopIdList,omitempty"`
}

// InventoryDailyReportDTO represents an inventory daily statement record.
type InventoryDailyReportDTO struct {
	Sku                  string `json:"sku,omitempty"`
	SkuName              string `json:"skuName,omitempty"`
	WarehouseName        string `json:"warehouseName,omitempty"`
	Date                 string `json:"date,omitempty"`
	BeginQuantity        int64  `json:"beginQuantity,omitempty"`
	InQuantity           int64  `json:"inQuantity,omitempty"`
	OutQuantity          int64  `json:"outQuantity,omitempty"`
	EndQuantity          int64  `json:"endQuantity,omitempty"`
	AdjustmentQuantity   int64  `json:"adjustmentQuantity,omitempty"`
	UpdateTime           int64  `json:"updateTime,omitempty"`
}

// QueryInventoryDailyReport queries inventory daily statement.
func (s *ReportService) QueryInventoryDailyReport(params *InventoryDailyReportQuery) ([]InventoryDailyReportDTO, int, error) {
	biz, _ := json.Marshal(params)
	var list []InventoryDailyReportDTO
	w := &ResponseWrapper{Result: &list}
	if err := s.client.Do("QUERY_INVENTORY_DAILY_REPORT", string(biz), w); err != nil {
		return nil, 0, err
	}
	if w.HasError() {
		return nil, 0, &APIError{ErrorCode: w.ErrorCode, Message: w.ErrorMsg, RequestID: w.RequestID}
	}
	return list, w.BizContent.Total, nil
}
