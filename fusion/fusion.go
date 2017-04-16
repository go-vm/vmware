// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fusion

import "github.com/go-vm/vmware"

const app = "fusion"

// Start start a VM or Team.
func Start(vmx string, gui bool) error {
	cmd := vmware.VMRun(app, "start", vmx)

	if gui {
		cmd.Args = append(cmd.Args, "gui")
	} else {
		cmd.Args = append(cmd.Args, "nogui")
	}

	return cmd.Run()
}

// Stop stop a VM or Team.
func Stop(vmx string, force bool) error {
	cmd := vmware.VMRun(app, "stop", vmx)

	if force {
		cmd.Args = append(cmd.Args, "hard")
	} else {
		cmd.Args = append(cmd.Args, "soft")
	}

	return cmd.Run()
}

// Reset reset a VM or Team.
func Reset(vmx string, force bool) error {
	cmd := vmware.VMRun(app, "reset", vmx)

	if force {
		cmd.Args = append(cmd.Args, "hard")
	} else {
		cmd.Args = append(cmd.Args, "soft")
	}

	return cmd.Run()
}

// Suspend Suspend a VM or Team
func Suspend(vmx string, force bool) error {
	cmd := vmware.VMRun(app, "suspend", vmx)

	if force {
		cmd.Args = append(cmd.Args, "hard")
	} else {
		cmd.Args = append(cmd.Args, "soft")
	}

	return cmd.Run()
}

// Pause pause a VM.
func Pause(vmx string) error {
	cmd := vmware.VMRun(app, "pause", vmx)

	return cmd.Run()
}

// Unpause unpause a VM.
func Unpause(vmx string) error {
	cmd := vmware.VMRun(app, "unpause", vmx)

	return cmd.Run()
}
