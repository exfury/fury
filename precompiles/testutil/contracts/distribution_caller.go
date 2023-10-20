// Copyright Tharsis Labs Ltd.(Fury)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/exfury/fury/blob/main/LICENSE)

package contracts

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	evmtypes "github.com/exfury/fury/v15/x/evm/types"
)

var (
	//go:embed DistributionCaller.json
	DistributionCallerJSON []byte

	// DistributionCallerContract is the compiled contract calling the distribution precompile
	DistributionCallerContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(DistributionCallerJSON, &DistributionCallerContract)
	if err != nil {
		panic(err)
	}

	if len(DistributionCallerContract.Bin) == 0 {
		panic("failed to load smart contract that calls distribution precompile")
	}
}
