// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package vmware

import (
	"os"
	"path/filepath"
)

// VMRunPath is default path to vmrun to fallback when it is not on path.
var VMRunPath string

// VMwareProducts define VMware products those contain vmrun.exe in Windows.
var VMwareProducts = [...]string{"VMware Workstation", "VMware Player", "VMware VIX"}

func init() {
	for _, products := range VMwareProducts {
		path := filepath.Join(os.Getenv("ProgramFiles(x86)"), "VMware", products, "vmrun.exe")
		if _, err := os.Stat(path); err == nil {
			VMRunPath = path
			break
		}
	}
}
