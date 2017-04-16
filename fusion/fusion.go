// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fusion

import "github.com/go-vm/vmware"

const fusion = "fusion"

// Start start a VM or Team.
func Start(gui bool) error {
	vmrun := vmware.NewVMRun()
	cmd := vmrun.Cmd(fusion, "start")

	if gui {
		cmd.Args = append(cmd.Args, "gui")
	} else {
		cmd.Args = append(cmd.Args, "nogui")
	}

	return cmd.Run()
}

// Stop stop a VM or Team.
func Stop(force bool) error {
	vmrun := vmware.NewVMRun()
	cmd := vmrun.Cmd(fusion, "start")

	if force {
		cmd.Args = append(cmd.Args, "hard")
	} else {
		cmd.Args = append(cmd.Args, "soft")
	}

	return cmd.Run()
}

// Reset reset a VM or Team.
func Reset(force bool) error {
	vmrun := vmware.NewVMRun()
	cmd := vmrun.Cmd(fusion, "reset")

	if force {
		cmd.Args = append(cmd.Args, "hard")
	} else {
		cmd.Args = append(cmd.Args, "soft")
	}

	return cmd.Run()
}

// Suspend Suspend a VM or Team
func Suspend(force bool) error {
	vmrun := vmware.NewVMRun()
	cmd := vmrun.Cmd(fusion, "suspend")

	if force {
		cmd.Args = append(cmd.Args, "hard")
	} else {
		cmd.Args = append(cmd.Args, "soft")
	}

	return cmd.Run()
}

// Pause pause a VM.
func Pause() error {
	vmrun := vmware.NewVMRun()
	cmd := vmrun.Cmd(fusion, "pause")

	return cmd.Run()
}

// Unpause unpause a VM.
func Unpause() error {
	vmrun := vmware.NewVMRun()
	cmd := vmrun.Cmd(fusion, "unpause")

	return cmd.Run()
}
