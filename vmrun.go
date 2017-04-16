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
func (v *VMRun) Cmd(app string, arg ...string) *exec.Cmd {
	cmd := exec.Command(v.vmrunPath, "-T", app)
	cmd.Args = append(cmd.Args, arg...)

	return cmd
}
