// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fusion

import (
	"github.com/go-vm/vmware"
)

// POWER COMMANDS           PARAMETERS           DESCRIPTION
// --------------           ----------           -----------
// start                    Path to vmx file     Start a VM or Team
//                          [gui|nogui]
//
// stop                     Path to vmx file     Stop a VM or Team
//                          [hard|soft]
//
// reset                    Path to vmx file     Reset a VM or Team
//                          [hard|soft]
//
// suspend                  Path to vmx file     Suspend a VM or Team
//                          [hard|soft]
//
// pause                    Path to vmx file     Pause a VM
//
// unpause                  Path to vmx file     Unpause a VM

// Start start a VM or Team.
func Start(vmx string, gui bool) error {
	flag := "nogui"
	if gui {
		flag = "gui"
	}

	if _, err := vmware.VMRun(app, "start", vmx, flag); err != nil {
		return err
	}

	return nil
}

// Stop stop a VM or Team.
func Stop(vmx string, hard bool) error {
	flag := "soft"
	if hard {
		flag = "hard"
	}

	if _, err := vmware.VMRun(app, "stop", vmx, flag); err != nil {
		return err
	}

	return nil
}

// ShutDown wrap of stop command with hard.
func ShutDown(vmx string) error {
	return Stop(vmx, true)
}

// Halt wrap of stop command with soft.
func Halt(vmx string) error {
	return Stop(vmx, false)
}

// Reset reset a VM or Team.
func Reset(vmx string, hard bool) error {
	flag := "soft"
	if hard {
		flag = "hard"
	}

	if _, err := vmware.VMRun(app, "reset", vmx, flag); err != nil {
		return err
	}

	return nil
}

// Restart restart a VM uses wrap of reset command with soft.
func Restart(vmx string) error {
	return Reset(vmx, false)
}

// Suspend Suspend a VM or Team.
func Suspend(vmx string, hard bool) error {
	flag := "soft"
	if hard {
		flag = "hard"
	}

	if _, err := vmware.VMRun(app, "suspend", vmx, flag); err != nil {
		return err
	}

	return nil
}

// Pause pause a VM.
func Pause(vmx string) error {
	if _, err := vmware.VMRun(app, "pause", vmx); err != nil {
		return err
	}

	return nil
}

// Unpause unpause a VM.
func Unpause(vmx string) error {
	if _, err := vmware.VMRun(app, "unpause", vmx); err != nil {
		return err
	}

	return nil
}
