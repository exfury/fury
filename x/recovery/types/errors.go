// Copyright Tharsis Labs Ltd.(Fury)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/exfury/fury/blob/main/LICENSE)

package types

import (
	errorsmod "cosmossdk.io/errors"
)

// errors
var (
	ErrBlockedAddress = errorsmod.Register(ModuleName, 2, "blocked address")
)
