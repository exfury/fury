// Copyright Tharsis Labs Ltd.(Fury)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/exfury/fury/blob/main/LICENSE)

package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/exfury/fury/v15/x/vesting/client/cli"
)

var RegisterClawbackProposalHandler = govclient.NewProposalHandler(cli.NewClawbackProposalCmd)
