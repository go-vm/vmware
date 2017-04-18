// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vmwareutil

import (
	"os/exec"
	"path/filepath"
)

// LookPath detect the vmware command binary path.
func LookPath(cmd string) string {
	if path, err := exec.LookPath(cmd); err == nil {
		return path
	}

	return filepath.Join(vmwareDir, cmd) // vmwareDir is OS specific variable.
}
