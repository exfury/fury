// Copyright Tharsis Labs Ltd.(Fury)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/exfury/fury/blob/main/LICENSE)
package grpc

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	revtypes "github.com/exfury/fury/v15/x/revenue/v1/types"
)

// GetRevenue returns the revenue for the given address.
func (gqh *IntegrationHandler) GetRevenue(address common.Address) (*revtypes.QueryRevenueResponse, error) {
	revenueClient := gqh.network.GetRevenueClient()
	return revenueClient.Revenue(context.Background(), &revtypes.QueryRevenueRequest{
		ContractAddress: address.String(),
	})
}

// GetRevenueParams gets the revenue module params.
func (gqh *IntegrationHandler) GetRevenueParams() (*revtypes.QueryParamsResponse, error) {
	revenueClient := gqh.network.GetRevenueClient()
	return revenueClient.Params(context.Background(), &revtypes.QueryParamsRequest{})
}
