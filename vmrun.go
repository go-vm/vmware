// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vmware

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"
)

var vmrunPath = vmwareCmd("vmrun")

// VMRun run the vmrun command with the app name and args.
// Return the stdout result and cmd error.
func VMRun(app string, arg ...string) (string, error) {
	// vmrun with nogui on VMware Fusion through at least 8.0.1 doesn't work right
	// if the umask is set to not allow world-readable permissions
	_ = syscall.Umask(022)

	cmd := exec.Command(vmrunPath, "-T", app)
	cmd.Args = append(cmd.Args, arg...)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		if runErr := err.(*exec.ExitError); runErr != nil {
			return "", fmt.Errorf(stdout.String())
		}
	}

	return stdout.String(), err
}
