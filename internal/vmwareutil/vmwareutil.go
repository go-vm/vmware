// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package vmwareutil implements VMware commands utitily.
package vmwareutil

import (
	"os/exec"
	"path/filepath"
	"runtime"
)

// LookPath detect the vmware command binary path.
func LookPath(cmd string) string {
	if runtime.GOOS == "windows" {
		cmd = cmd + ".exe"
	}

	if path, err := exec.LookPath(cmd); err == nil {
		return path
	}

	return filepath.Join(vmwareDir, cmd) // vmwareDir is OS specific variable.
}
