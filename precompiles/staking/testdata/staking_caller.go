// Copyright Tharsis Labs Ltd.(Fury)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/exfury/fury/blob/main/LICENSE)

package testdata

import (
	_ "embed" // embed compiled smart contract
	"encoding/json"

	evmtypes "github.com/exfury/fury/v15/x/evm/types"
)

var (
	//go:embed StakingCaller.json
	StakingCallerJSON []byte

	// StakingCallerContract is the compiled contract calling the staking precompile
	StakingCallerContract evmtypes.CompiledContract
)

func init() {
	err := json.Unmarshal(StakingCallerJSON, &StakingCallerContract)
	if err != nil {
		panic(err)
	}

	if len(StakingCallerContract.Bin) == 0 {
		panic("failed to load smart contract that calls staking precompile")
	}
}
