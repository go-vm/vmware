// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vmware

import (
	"bytes"
	"fmt"
	"os/exec"
)

var vmrun string

func init() {
	var err error

	vmrun, err = exec.LookPath("vmrun")
	if err != nil {
		vmrun = VMRunPath // fallback the vmrun binary path
	}
}

// VMRun return the vmrun execute binary command with the app name.
//
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
func VMRun(app string, arg ...string) (string, error) {
	cmd := exec.Command(vmrun, "-T", app)
	cmd.Args = append(cmd.Args, arg...)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	if runErr := cmd.Run(); runErr != nil {
		if err := runErr.(*exec.ExitError); err != nil {
			return "", fmt.Errorf(stdout.String())
		}
		return "", runErr
	}

	return stdout.String(), nil
}
