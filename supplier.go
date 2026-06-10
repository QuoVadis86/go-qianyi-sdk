package qianyi

import "context"

// SupplierService provides access to supplier API operations.
type SupplierService struct {
	client *Client
}

// NewSupplierService creates a new SupplierService.
func NewSupplierService(client *Client) *SupplierService {
	return &SupplierService{client: client}
}

// Supplier represents a supplier in QERP.
type Supplier struct {
	Name                    string                  `json:"name"`
	Level                   string                  `json:"level,omitempty"`
	CategoryOneName         string                  `json:"categoryOneName,omitempty"`
	CategoryTwoName         string                  `json:"categoryTwoName,omitempty"`
	CategoryThreeName       string                  `json:"categoryThreeName,omitempty"`
	PurchaserName           string                  `json:"purchaserName,omitempty"`
	PurchaserUserName       string                  `json:"purchaserUserName,omitempty"`
	SettlementWay           string                  `json:"settlementWay,omitempty"`
	AccountPeriodList       []SupplierAccountPeriod `json:"accountPeriodList,omitempty"`
	AccountPeriodOpt        string                  `json:"accountPeriodOpt,omitempty"`
	BillingDate             string                  `json:"billingDate,omitempty"`
	EffectiveNode           string                  `json:"effectiveNode,omitempty"`
	PaymentWay              string                  `json:"paymentWay,omitempty"`
	PrepayRate              float64                 `json:"prepayRate,omitempty"`
	TransportParty          string                  `json:"transportParty,omitempty"`
	DefectiveProductResolution string                `json:"defectiveProductResolution,omitempty"`
	Contacts                string                  `json:"contacts,omitempty"`
	ContactsInfo            string                  `json:"contactsInfo,omitempty"`
	ContactsAddress         string                  `json:"contactsAddress,omitempty"`
	Province                string                  `json:"province,omitempty"`
	City                    string                  `json:"city,omitempty"`
	SupplySource            string                  `json:"supplySource,omitempty"`
	LegalRepresentative     string                  `json:"legalRepresentative,omitempty"`
	EntrustedAgent          string                  `json:"entrustedAgent,omitempty"`
	DepositBank             string                  `json:"depositBank,omitempty"`
	Account                 string                  `json:"account,omitempty"`
	Enable                  bool                    `json:"enable"`
	Remark                  string                  `json:"remark,omitempty"`
	Category                string                  `json:"category,omitempty"`
	CreateTime              string                  `json:"createTime,omitempty"`
	UpdateTime              string                  `json:"updateTime,omitempty"`
	Country                 string                  `json:"country,omitempty"`
}

// SupplierAccountPeriod represents a payment period for a supplier.
type SupplierAccountPeriod struct {
	Days    string `json:"days"`
	Percent string `json:"percent"`
}

// QuerySupplierParams holds parameters for querying suppliers.
type QuerySupplierParams struct {
	Page               int    `json:"page"`
	PageSize           int    `json:"pageSize"`
	Name               string `json:"name,omitempty"`
	PurchaserUserName  string `json:"purchaserUserName,omitempty"`
	Category           string `json:"category,omitempty"`
	Enable             *bool  `json:"enable,omitempty"`
	CreateTimeFrom     string `json:"createTimeFrom,omitempty"`
	CreateTimeTo       string `json:"createTimeTo,omitempty"`
}

// CreateSupplierParams holds parameters for creating a supplier.
type CreateSupplierParams struct {
	Name                    string                  `json:"name"`
	Category                string                  `json:"category"`
	PurchaserUserName       string                  `json:"purchaserUserName"`
	SettlementWay           string                  `json:"settlementWay"`
	PaymentWay              string                  `json:"paymentWay"`
	TransportParty          string                  `json:"transportParty"`
	DefectiveProductResolution string                `json:"defectiveProductResolution"`
	Level                   string                  `json:"level,omitempty"`
	CategoryOneName         string                  `json:"categoryOneName,omitempty"`
	CategoryTwoName         string                  `json:"categoryTwoName,omitempty"`
	CategoryThreeName       string                  `json:"categoryThreeName,omitempty"`
	Contacts                string                  `json:"contacts,omitempty"`
	ContactsInfo            string                  `json:"contactsInfo,omitempty"`
	Province                string                  `json:"province,omitempty"`
	Country                 string                  `json:"country,omitempty"`
	City                    string                  `json:"city,omitempty"`
	ContactsAddress         string                  `json:"contactsAddress,omitempty"`
	SupplySource            string                  `json:"supplySource,omitempty"`
	Authorization           *bool                   `json:"authorization,omitempty"`
	AuthUserNameList        []string                `json:"authUserNameList,omitempty"`
	EffectiveNode           string                  `json:"effectiveNode,omitempty"`
	AccountPeriodList       []SupplierAccountPeriod `json:"accountPeriodList,omitempty"`
	AccountPeriodOpt        string                  `json:"accountPeriodOpt,omitempty"`
	BillingDate             string                  `json:"billingDate,omitempty"`
	PrepayRate              float64                 `json:"prepayRate,omitempty"`
	LegalRepresentative     string                  `json:"legalRepresentative,omitempty"`
	EntrustedAgent          string                  `json:"entrustedAgent,omitempty"`
	DepositBank             string                  `json:"depositBank,omitempty"`
	Account                 string                  `json:"account,omitempty"`
	Remark                  string                  `json:"remark,omitempty"`
}

// SupplierSku represents a supplier-SKU relationship in QERP.
type SupplierSku struct {
	Sku                        string                   `json:"sku"`
	Title                      string                   `json:"title"`
	Description                string                   `json:"description,omitempty"`
	SkuSupplierList            []SupplierSkuEntry       `json:"skuSupplierList,omitempty"`
}

// SupplierSkuEntry represents a single supplier record for a SKU.
type SupplierSkuEntry struct {
	DefaultPurchaserName     string  `json:"defaultPurchaserName,omitempty"`
	DefaultPurchaserUserName string  `json:"defaultPurchaserUserName,omitempty"`
	Category                 string  `json:"category"`
	DeliveryCycle            int64   `json:"deliveryCycle"`
	PurchaseMethod           string  `json:"purchaseMethod,omitempty"`
	DefaultValue             string  `json:"defaultValue,omitempty"`
	PurchasePriceUnit        string  `json:"purchasePriceUnit,omitempty"`
	DrawbackRate             float64 `json:"drawbackRate,omitempty"`
	PurchaseTaxRate          float64 `json:"purchaseTaxRate,omitempty"`
	PurchasePrice            string  `json:"purchasePrice,omitempty"`
	MinimumPurchaseQuantity  int64   `json:"minimumPurchaseQuantity,omitempty"`
	SupplierName             string  `json:"supplierName"`
	AssociatedName           string  `json:"associatedName,omitempty"`
}

// QuerySupplierSkuParams holds parameters for querying supplier SKUs.
type QuerySupplierSkuParams struct {
	Page                    int    `json:"page"`
	PageSize                int    `json:"pageSize"`
	Sku                     string `json:"sku,omitempty"`
	Title                   string `json:"title,omitempty"`
	Category                string `json:"category,omitempty"`
	SupplierName            string `json:"supplierName,omitempty"`
	DefaultPurchaserUserName string `json:"defaultPurchaserUserName,omitempty"`
	RelationStatus          string `json:"relationStatus,omitempty"`
	CreateTimeFrom          string `json:"createTimeFrom,omitempty"`
	CreateTimeTo            string `json:"createTimeTo,omitempty"`
}

// CreateSupplierSkuParams holds parameters for creating a supplier-SKU relationship.
type CreateSupplierSkuParams struct {
	Sku                    string  `json:"sku"`
	SupplierName           string  `json:"supplierName"`
	DefaultPurchaserUserName string `json:"defaultPurchaserUserName"`
	DeliveryCycle          int     `json:"deliveryCycle"`
	PurchaseMethod         string  `json:"purchaseMethod"`
	MinimumPurchaseQuantity int    `json:"minimumPurchaseQuantity,omitempty"`
	PurchasePrice          float64 `json:"purchasePrice,omitempty"`
	PurchasePriceUnit      string  `json:"purchasePriceUnit,omitempty"`
	PurchaseTaxRate        float64 `json:"purchaseTaxRate,omitempty"`
	DrawbackRate           float64 `json:"drawbackRate,omitempty"`
	PurchaseURL            string  `json:"purchaseUrl,omitempty"`
	IsDefault              int     `json:"isDefault,omitempty"`
}

// QueryList retrieves suppliers with optional filters.
func (s *SupplierService) QueryList(ctx context.Context, params *QuerySupplierParams) ([]Supplier, int, error) {
	return doList[Supplier](ctx, s.client, ServiceTypeQuerySupplierList, params)
}

// Create creates a new supplier.
func (s *SupplierService) Create(ctx context.Context, params *CreateSupplierParams) error {
	return doAction(ctx, s.client, ServiceTypeCreateSupplier, params)
}

// QuerySkuList retrieves supplier SKU relationships with optional filters.
func (s *SupplierService) QuerySkuList(ctx context.Context, params *QuerySupplierSkuParams) ([]SupplierSku, int, error) {
	return doList[SupplierSku](ctx, s.client, ServiceTypeQuerySupplierSkuList, params)
}

// CreateSku creates a new supplier-SKU relationship.
func (s *SupplierService) CreateSku(ctx context.Context, params *CreateSupplierSkuParams) error {
	return doAction(ctx, s.client, ServiceTypeCreateSupplierSku, params)
}
