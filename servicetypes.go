package qianyi

// Service type constants for all QERP API endpoints.
const (
	// Shop
	ServiceTypeQueryShopList = "QUERY_SHOP_LIST"

	// Warehouse
	ServiceTypeQueryWarehouseList = "QUERY_WAREHOUSE_LIST"

	// SKU
	ServiceTypeQuerySimpleListSku    = "QUERY_SIMPLE_LIST_SKU"
	ServiceTypeInsertSkuInfo         = "INSERT_SKU_INFO"
	ServiceTypeUpdateSkuInfo         = "UPDATE_SKU_INFO"
	ServiceTypeEnableSku             = "ENABLE_SKU"
	ServiceTypeQuerySysSku           = "QUERY_SYS_SKU"

	// Order
	ServiceTypeCreateSalesOrder            = "CREATE_SALES_ORDER"
	ServiceTypeCloseSalesOrder             = "CLOSE_SALES_ORDER"
	ServiceTypeQuerySalesOrderList         = "QUERY_SALES_ORDER_LIST"
	ServiceTypeQuerySalesOrderNumberList   = "QUERY_SALES_ORDER_NUMBER_LIST"
	ServiceTypeQuerySalesOrderShippingInfo = "QUERY_SALES_ORDER_SHIPPING_INFO"
	ServiceTypeQuerySalesOrderAudit        = "QUERY_SALES_ORDER_AUDIT"
	ServiceTypeCreateWaveOrder             = "CREATE_WAVE_ORDER"
	ServiceTypeSendSalesOrderToWms         = "SEND_SALES_ORDER_TO_WMS"
	ServiceTypeQueryOriginalSalesOrder     = "QUERY_ORIGINAL_SALES_ORDER"
	ServiceTypeQuerySalesOrderPickupStatus = "QUERY_SALES_ORDER_PICKUP_STATUS"
	ServiceTypeQuerySalesOrderDocument     = "QUERY_SALES_ORDER_DOCUMENT"
	ServiceTypeSubscribeOrder              = "SUBSCRIBE_ORDER"

	// Refund
	ServiceTypeCreateReturnOrder    = "CREATE_RETURN_ORDER"
	ServiceTypeCloseReturnOrder     = "CLOSE_RETURN_ORDER"
	ServiceTypeQueryReturnOrderList = "QUERY_RETURN_ORDER_LIST"
	ServiceTypePushReturnOrderInfo  = "PUSH_RETURN_ORDER_INFO"

	// Inventory
	ServiceTypeQuerySimpleListInventory    = "QUERY_SIMPLE_LIST_INVENTORY"
	ServiceTypeQuerySimpleListInventoryV2  = "QUERY_SIMPLE_LIST_INVENTORY_V2"
	ServiceTypeQueryInventoryLogList       = "QUERY_INVENTORY_LOG_LIST"
	ServiceTypeQueryInventoryAssemblyList  = "QUERY_INVENTORY_ASSEMBLY_LIST"
	ServiceTypeCreateTransferOrder         = "CREATE_TRANSFER_ORDER"
	ServiceTypeQueryTransferOrderList      = "QUERY_TRANSFER_ORDER_LIST"
	ServiceTypeQuerySplitOrderList         = "QUERY_SPLIT_ORDER_LIST"
	ServiceTypeQueryStorageLocInventory    = "QUERY_STORAGE_LOC_INVENTORY"
	ServiceTypeQueryBatchInventoryList     = "QUERY_BATCH_INVENTORY_LIST"
	ServiceTypeTransferStorageLocation     = "TRANSFER_STORAGE_LOCATION"
	ServiceTypeQuerySbsInventoryList       = "QUERY_SBS_INVENTORY_LIST"
	ServiceTypeQuerySbsWarehouseList       = "QUERY_SBS_WAREHOUSE_LIST"

	// ASN
	ServiceTypeCreateAsnOrder    = "CREATE_ASN_ORDER"
	ServiceTypeQueryAsnList      = "QUERY_ASN_LIST"
	ServiceTypeCancelAsnOrder    = "CANCEL_ASN_ORDER"
	ServiceTypeDeleteAsnOrder    = "DELETE_ASN_ORDER"
	ServiceTypePushAsnOrder      = "PUSH_ASN_ORDER"
	ServiceTypeQueryAsnBatchList = "QUERY_ASN_BATCH_LIST"

	// ODO
	ServiceTypeQueryOdoList      = "QUERY_ODO_LIST"
	ServiceTypeQuerySalesOdoList = "QUERY_SALES_ODO_LIST"
	ServiceTypeCreateOdoOrder    = "CREATE_ODO_ORDER"
	ServiceTypeCancelOdoOrder    = "CANCEL_ODO_ORDER"
	ServiceTypePushOdoOrder      = "PUSH_ODO_ORDER"

	// Adjustment
	ServiceTypeQueryAdjustmentList   = "QUERY_ADJUSTMENT_LIST"
	ServiceTypeCreateAdjustmentOrder = "CREATE_ADJUSTMENT_ORDER"

	// Purchase
	ServiceTypeQueryPurchaseOrderList = "QUERY_PURCHASE_ORDER_LIST"
	ServiceTypeCreatePurchaseOrder    = "CREATE_PURCHASE_ORDER"

	// Logistics (First Leg)
	ServiceTypeQueryFirstLegOrderList       = "QUERY_FIRST_LEG_ORDER_LIST"
	ServiceTypeCreateFirstLegOrder          = "CREATE_FIRST_LEG_ORDER"
	ServiceTypeQueryFirstLrgLogistics       = "QUERY_FIRST_LRG_LOGISTICS"
	ServiceTypeQueryFirstLrgTrackingPackage = "QUERY_FIRST_LRG_TRACKING_PACKAGE"
	ServiceTypeWithdrawAndDelFirstLeg       = "WITHDRAW_AND_DEL_FIRST_LEG"
	ServiceTypePushTrackingPackage          = "PUSH_TRACKING_PACKAGE"

	// Report
	ServiceTypeQueryShopeeTransactionDetailList    = "QUERY_SHOPEE_TRANSACTION_DETAIL_LIST"
	ServiceTypeQueryLazadaTransactionDetailList    = "QUERY_LAZADA_TRANSACTION_DETAIL_LIST"
	ServiceTypeQueryTiktokTransactionDetailList    = "QUERY_TIKTOK_TRANSACTION_DETAIL_LIST"
	ServiceTypeQueryShopeePayoutDetailList         = "QUERY_SHOPEE_PAYOUT_DETAIL_LIST"
	ServiceTypeQueryLazadaAccountTransactionList   = "QUERY_LAZADA_ACCOUNT_TRANSACTION_LIST"
	ServiceTypeQueryTiktokV2TransactionDetailList  = "QUERY_TIKTOK_V2_TRANSACTION_DETAIL_LIST"
	ServiceTypeQueryTiktokPayoutRecord             = "QUERY_TIKTOK_PAYOUT_RECORD"
	ServiceTypeQueryInventoryDailyReport           = "QUERY_INVENTORY_DAILY_REPORT"

	// CustomerField
	ServiceTypeCustomerFieldQuery = "CUSTOMER_FIELD_QUERY"

	// Supplier
	ServiceTypeQuerySupplierList    = "QUERY_SUPPLIER_LIST"
	ServiceTypeCreateSupplier       = "CREATE_SUPPLIER"
	ServiceTypeQuerySupplierSkuList = "QUERY_SUPPLIER_SKU_LIST"
	ServiceTypeCreateSupplierSku    = "CREATE_SUPPLIER_SKU"
)
