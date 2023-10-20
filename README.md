<!--
parent:
  order: false
-->

<div align="center">
  <h1> Fury </h1>
</div>

<div align="center">
  <a href="https://github.com/exfury/fury/releases/latest">
    <img alt="Version" src="https://img.shields.io/github/tag/tharsis/fury.svg" />
  </a>
  <a href="https://github.com/exfury/fury/blob/main/LICENSE">
    <img alt="License: Apache-2.0" src="https://img.shields.io/github/license/tharsis/fury.svg" />
  </a>
  <a href="https://pkg.go.dev/github.com/exfury/fury">
    <img alt="GoDoc" src="https://godoc.org/github.com/exfury/fury?status.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/exfury/fury">
    <img alt="Go report card" src="https://goreportcard.com/badge/github.com/exfury/fury"/>
  </a>
  <a href="https://bestpractices.coreinfrastructure.org/projects/5018">
    <img alt="Lines of code" src="https://img.shields.io/tokei/lines/github/tharsis/fury">
  </a>
</div>
<div align="center">
  <a href="https://discord.gg/fury">
    <img alt="Discord" src="https://img.shields.io/discord/809048090249134080.svg" />
  </a>
  <a href="https://github.com/exfury/fury/actions?query=branch%3Amain+workflow%3ALint">
    <img alt="Lint Status" src="https://github.com/exfury/fury/actions/workflows/lint.yml/badge.svg?branch=main" />
  </a>
  <a href="https://codecov.io/gh/exfury/fury">
    <img alt="Code Coverage" src="https://codecov.io/gh/exfury/fury/branch/main/graph/badge.svg" />
  </a>
  <a href="https://twitter.com/FuryOrg">
    <img alt="Twitter Follow Fury" src="https://img.shields.io/twitter/follow/FuryOrg"/>
  </a>
</div>

Fury is a scalable, high-throughput Proof-of-Stake blockchain
that is fully compatible and interoperable with Ethereum.
It's built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk/)
which runs on top of the [Tendermint Core](https://github.com/tendermint/tendermint) consensus engine.

## Quick Start

To learn how Fury works from a high-level perspective,
go to the [Protocol Overview](https://docs.fury.org/protocol) section of the documentation.
You can also check the instructions to [Run a Node](https://docs.fury.org/protocol/fury-cli#run-an-fury-node).

## Documentation

Our documentation is hosted in a [separate repository](https://github.com/fury/docs) and can be found at [docs.fury.org](https://docs.fury.org).
Head over there and check it out.

## Installation

For prerequisites and detailed build instructions
please read the [Installation](https://docs.fury.org/protocol/fury-cli) instructions.
Once the dependencies are installed, run:

```bash
make install
```

Or check out the latest [release](https://github.com/exfury/fury/releases).

## Community

The following chat channels and forums are great spots to ask questions about Fury:

- [Fury Twitter](https://twitter.com/FuryOrg)
- [Fury Discord](https://discord.gg/fury)
- [Fury Forum](https://commonwealth.im/fury)

## Contributing

Looking for a good place to start contributing?
Check out some
[`good first issues`](https://github.com/exfury/fury/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22).

For additional instructions, standards and style guides, please refer to the [Contributing](./CONTRIBUTING.md) document.

## Careers

See our open positions on [Greenhouse](https://boards.eu.greenhouse.io/fury).

## Licensing

Starting from April 21st, 2023, the Fury repository will update its License
from GNU Lesser General Public License v3.0 (LGPLv3) to Fury Non-Commercial
License 1.0 (ENCL-1.0). This license applies to all software released from Fury
version 13 or later, except for specific files, as follows, which will continue
to be licensed under LGPLv3:

- `x/revenue/v1/` (all files in this folder)
- `x/claims/genesis.go`
- `x/erc20/keeper/proposals.go`
- `x/erc20/types/utils.go`

LGPLv3 will continue to apply to older versions (<v13.0.0) of the Fury
repository. For more information see LICENSE.

### SPDX Identifier

The following header including a license identifier in SPDX short form has been added to all ENCL-1.0 files:

```go
// Copyright Tharsis Labs Ltd.(Fury)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/exfury/fury/blob/main/LICENSE)
```

Exempted files contain the following SPDX ID:

```go
// Copyright Tharsis Labs Ltd.(Fury)
// SPDX-License-Identifier:LGPL-3.0-only
```

### License FAQ

Find below an overview of the Permissions and Limitations of the Fury Non-Commercial License 1.0.
For more information, check out the full ENCL-1.0 FAQ [here](/LICENSE_FAQ.md).

| Permissions                                                                                                                                                                  | Prohibited                                                                 |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| - Private Use, including distribution and modification<br />- Commercial use on designated blockchains<br />- Commercial use with Fury permit (to be separately negotiated) | - Commercial use, other than on designated blockchains, without Fury permit |
