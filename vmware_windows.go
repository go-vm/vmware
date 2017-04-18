// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vmware

import (
	"os"
	"path/filepath"
)

// vmwareDir is default directory path of vmware commands to fallback when it is not on path.
var vmwareDir string

// vmwareProducts define VMware products those contain vmrun.exe in Windows.
var vmwareProducts = [...]string{"VMware Workstation", "VMware Player", "VMware VIX"}

func init() {
	for _, products := range vmwareProducts {
		path := filepath.Join(os.Getenv("ProgramFiles(x86)"), "VMware", products)
		if _, err := os.Stat(path); err == nil {
			vmwareDir = path
			break
		}
	}
}
