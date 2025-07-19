package resolver

import "github.com/ehharvey/malleus/internal/inventory"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	InventoryService inventory.Service
}
