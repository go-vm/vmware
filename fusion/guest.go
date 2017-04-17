// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fusion

import (
	"github.com/go-vm/vmware"
)

// Guest represents a guest login data.
type Guest struct {
	User string
	Pass string
}

type RunProgramInGuestConfig int

const (
	NoWait RunProgramInGuestConfig = 1 << iota
	ActiveWindow
	Interactive
)

// RunProgramInGuest run a program in Guest OS.
func RunProgramInGuest(vmx string, guest Guest, config RunProgramInGuestConfig, cmdPath string, cmdArgs ...string) error {
	args := []string{"-gu", guest.User, "-gp", guest.Pass, "runProgramInGuest", vmx}

	if config&NoWait > 0 {
		args = append(args, "-noWait")
	}
	if config&ActiveWindow > 0 {
		args = append(args, "-activeWindow")
	}
	if config&Interactive > 0 {
		args = append(args, "-interactive")
	}

	args = append(args, cmdPath)
	args = append(args, cmdArgs...)

	if _, err := vmware.VMRun(app, args...); err != nil {
		return err
	}

	return nil
}

// FileExistsInGuest check if a file exists in Guest OS.
func FileExistsInGuest(vmx string, guest Guest, filename string) bool {
	if _, err := vmware.VMRun(app, "-gu", guest.User, "-gp", guest.Pass, "fileExistsInGuest", vmx, filename); err != nil {
		return false
	}

	return true
}

// DirectoryExistsInGuest check if a directory exists in Guest OS.
func DirectoryExistsInGuest(vmx string, guest Guest, dir string) bool {
	if _, err := vmware.VMRun(app, "-gu", guest.User, "-gp", guest.Pass, "directoryExistsInGuest", vmx, dir); err != nil {
		return false
	}

	return true
}
