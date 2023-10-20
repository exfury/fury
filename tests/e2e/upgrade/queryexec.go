// Copyright Tharsis Labs Ltd.(Fury)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/exfury/fury/blob/main/LICENSE)
package upgrade

import (
	"fmt"
)

// CreateModuleQueryExec creates a Fury module query
func (m *Manager) CreateModuleQueryExec(moduleName, subCommand, chainID string) (string, error) {
	cmd := []string{
		"furyd",
		"q",
		moduleName,
		subCommand,
		fmt.Sprintf("--chain-id=%s", chainID),
		"--keyring-backend=test",
		"--log_format=json",
	}
	return m.CreateExec(cmd, m.ContainerID())
}
