// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fusion

import "github.com/go-vm/vmware"

const app = "fusion"

// Start start a VM or Team.
func Start(vmwarevm string, gui bool) error {
	cmd := vmware.VMRun(app, "start", vmwarevm)

	if gui {
		cmd.Args = append(cmd.Args, "gui")
	} else {
		cmd.Args = append(cmd.Args, "nogui")
	}

	return cmd.Run()
}

// Stop stop a VM or Team.
func Stop(vmwarevm string, force bool) error {
	cmd := vmware.VMRun(app, "stop", vmwarevm)

	if force {
		cmd.Args = append(cmd.Args, "hard")
	} else {
		cmd.Args = append(cmd.Args, "soft")
	}

	return cmd.Run()
}

// Reset reset a VM or Team.
func Reset(vmwarevm string, force bool) error {
	cmd := vmware.VMRun(app, "reset", vmwarevm)

	if force {
		cmd.Args = append(cmd.Args, "hard")
	} else {
		cmd.Args = append(cmd.Args, "soft")
	}

	return cmd.Run()
}

// Suspend Suspend a VM or Team
func Suspend(vmwarevm string, force bool) error {
	cmd := vmware.VMRun(app, "suspend", vmwarevm)

	if force {
		cmd.Args = append(cmd.Args, "hard")
	} else {
		cmd.Args = append(cmd.Args, "soft")
	}

	return cmd.Run()
}

// Pause pause a VM.
func Pause(vmwarevm string) error {
	cmd := vmware.VMRun(app, "pause", vmwarevm)

	return cmd.Run()
}

// Unpause unpause a VM.
func Unpause(vmwarevm string) error {
	cmd := vmware.VMRun(app, "unpause", vmwarevm)

	return cmd.Run()
}
