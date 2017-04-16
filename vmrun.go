// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vmware

import (
	"os/exec"
)

// VMRun represents a vmrun command.
type VMRun struct {
	vmrunPath string
}

// NewVMRun return the new VmRun.
func NewVMRun() *VMRun {
	vmrunPath, err := exec.LookPath("vmrun")
	if err != nil {
		vmrunPath = VMRunPath // fallback the vmrun binary path
	}

	return &VMRun{
		vmrunPath: vmrunPath,
	}
}

// Usage: vmrun [AUTHENTICATION-FLAGS] COMMAND [PARAMETERS]
//
// AUTHENTICATION-FLAGS
// --------------------
// These must appear before the command and any command parameters.
//
//    -h <hostName>  (not needed for Fusion)
//    -P <hostPort>  (not needed for Fusion)
//    -T <hostType> (ws|fusion)
//    -u <userName in host OS>  (not needed for Fusion)
//    -p <password in host OS>  (not needed for Fusion)
//    -vp <password for encrypted virtual machine>
//    -gu <userName in guest OS>
//    -gp <password in guest OS>
func (v *VMRun) cmd(arg ...string) *exec.Cmd {
	cmd := exec.Command(v.vmrunPath, "-T", "fusion")
	cmd.Args = append(cmd.Args, arg...)

	return cmd
}

// Start start a VM or Team.
func (v *VMRun) Start(gui bool) error {
	cmd := v.cmd("start")

	if gui {
		cmd.Args = append(cmd.Args, "gui")
	} else {
		cmd.Args = append(cmd.Args, "nogui")
	}

	return cmd.Run()
}

// Stop stop a VM or Team.
func (v *VMRun) Stop(force bool) error {
	cmd := v.cmd("stop")

	if force {
		cmd.Args = append(cmd.Args, "hard")
	} else {
		cmd.Args = append(cmd.Args, "soft")
	}

	return cmd.Run()
}

// Reset reset a VM or Team.
func (v *VMRun) Reset(force bool) error {
	cmd := v.cmd("reset")

	if force {
		cmd.Args = append(cmd.Args, "hard")
	} else {
		cmd.Args = append(cmd.Args, "soft")
	}

	return cmd.Run()
}

// Suspend Suspend a VM or Team
func (v *VMRun) Suspend(force bool) error {
	cmd := v.cmd("suspend")

	if force {
		cmd.Args = append(cmd.Args, "hard")
	} else {
		cmd.Args = append(cmd.Args, "soft")
	}

	return cmd.Run()
}

// Pause pause a VM.
func (v *VMRun) Pause() error {
	cmd := v.cmd("unpause")

	return cmd.Run()
}

// Unpause unpause a VM.
func (v *VMRun) Unpause() error {
	cmd := v.cmd("unpause")

	return cmd.Run()
}
