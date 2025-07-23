package inventory

import (
	"context"

	"github.com/ehharvey/malleus/internal/outcome"
)

type Repository interface {
	// Domain
	CreateDomain(
		context context.Context,
		input CreateDomainParams,
	) (Domain, outcome.DbResult)
	CheckExistsDomainByName(
		context context.Context,
		input string,
	) (bool, outcome.DbResult)
	// ListDomains(input ListPaginationParams)

	// // Network
	// CreateNetwork(input CreateNetworkParams)
	// ListNetworks(input ListPaginationParams)

	// // NetworkBridge
	// CreateNetworkBridge(input CreateNetworkBridgeParams)
	// ListNetworkBridges(input ListPaginationParams)

	// // NetworkInterface
	// CreateNetworkInterface(input CreateNetworkInterfaceParams)
	// ListNetworkInterfaces(input ListPaginationParams)

	// // NetworkIp
	// CreateNetworkIp(input CreateNetworkIpParams)
	// ListNetworkIps(input ListPaginationParams)

	// // Node
	// CreateNode(input CreateNodeParams)
	// ListNodes(input ListPaginationParams)
}
