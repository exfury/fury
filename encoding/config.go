// Copyright Tharsis Labs Ltd.(Fury)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/exfury/fury/blob/main/LICENSE)
package encoding

import (
	"cosmossdk.io/simapp/params"
	amino "github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"

	enccodec "github.com/exfury/fury/v15/encoding/codec"
)

// MakeConfig creates an EncodingConfig for testing
func MakeConfig(mb module.BasicManager) params.EncodingConfig {
	cdc := amino.NewLegacyAmino()
	interfaceRegistry := types.NewInterfaceRegistry()
	codec := amino.NewProtoCodec(interfaceRegistry)

	encodingConfig := params.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             codec,
		TxConfig:          tx.NewTxConfig(codec, tx.DefaultSignModes),
		Amino:             cdc,
	}

	enccodec.RegisterLegacyAminoCodec(encodingConfig.Amino)
	mb.RegisterLegacyAminoCodec(encodingConfig.Amino)
	enccodec.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	mb.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
