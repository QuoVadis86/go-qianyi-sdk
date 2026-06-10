package main

import (
	"fmt"
	"log"

	"github.com/QuoVadis86/go-qianyi-sdk"
)

func main() {
	sdk := qianyi.NewSDK("your-app-id", "your-app-secret")
	sdk.TestEnv()

	shops, total, err := sdk.Shop.QueryList(1, 10, "", "", "", "")
	if err != nil {
		log.Fatalf("QueryShops failed: %v", err)
	}
	fmt.Printf("Total shops: %d\n", total)
	for _, shop := range shops {
		fmt.Printf("  Shop: %s (platform: %s, status: %s)\n", shop.Name, shop.Platform, shop.Status)
	}

	skus, total, err := sdk.Sku.QueryList(1, 10)
	if err != nil {
		log.Fatalf("QuerySkus failed: %v", err)
	}
	fmt.Printf("Total SKUs: %d\n", total)
	for _, sku := range skus {
		fmt.Printf("  SKU: %s - %s\n", sku.Sku, sku.Title)
	}

	warehouses, total, err := sdk.Warehouse.QueryList(1, 10, "", "")
	if err != nil {
		log.Fatalf("QueryWarehouses failed: %v", err)
	}
	fmt.Printf("Total warehouses: %d\n", total)
	for _, w := range warehouses {
		fmt.Printf("  Warehouse: %s (type: %s, country: %s)\n", w.Name, w.Kind, w.Country)
	}

	inventory, total, err := sdk.Inventory.QueryListV2(&qianyi.InventoryQueryParams{
		Page:      1,
		PageSize:  10,
		Warehouse: "your-warehouse-name",
	})
	if err != nil {
		log.Fatalf("QueryInventory failed: %v", err)
	}
	fmt.Printf("Total inventory items: %d\n", total)
	for _, item := range inventory {
		fmt.Printf("  SKU: %s - available: %d\n", item.Sku, item.Available)
	}
}
