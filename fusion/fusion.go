// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fusion

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/go-vm/vmware"
)

const app = "fusion"

// Start start a VM or Team.
func Start(vmwarevm string, gui bool) error {
	flag := "nogui"
	if gui {
		flag = "gui"
	}

	cmd := vmware.VMRun(app, "start", vmwarevm, flag)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	if runErr := cmd.Run(); runErr != nil {
		if err := runErr.(*exec.ExitError); err != nil {
			return fmt.Errorf(stdout.String())
		}
		return runErr
	}

	return nil
}

// Stop stop a VM or Team.
func Stop(vmwarevm string, hard bool) error {
	flag := "soft"
	if hard {
		flag = "hard"
	}

	cmd := vmware.VMRun(app, "stop", vmwarevm, flag)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	if runErr := cmd.Run(); runErr != nil {
		if err := runErr.(*exec.ExitError); err != nil {
			return fmt.Errorf(stdout.String())
		}
		return runErr
	}

	return nil
}

// Reset reset a VM or Team.
func Reset(vmwarevm string, hard bool) error {
	flag := "soft"
	if hard {
		flag = "hard"
	}

	cmd := vmware.VMRun(app, "reset", vmwarevm, flag)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	if runErr := cmd.Run(); runErr != nil {
		if err := runErr.(*exec.ExitError); err != nil {
			return fmt.Errorf(stdout.String())
		}
		return runErr
	}

	return nil
}

// Suspend Suspend a VM or Team.
func Suspend(vmwarevm string, hard bool) error {
	flag := "soft"
	if hard {
		flag = "hard"
	}

	cmd := vmware.VMRun(app, "suspend", vmwarevm, flag)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	if runErr := cmd.Run(); runErr != nil {
		if err := runErr.(*exec.ExitError); err != nil {
			return fmt.Errorf(stdout.String())
		}
		return runErr
	}

	return nil
}

// Pause pause a VM.
func Pause(vmwarevm string) error {
	cmd := vmware.VMRun(app, "pause", vmwarevm)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	if runErr := cmd.Run(); runErr != nil {
		if err := runErr.(*exec.ExitError); err != nil {
			return fmt.Errorf(stdout.String())
		}
		return runErr
	}

	return nil
}

// Unpause unpause a VM.
func Unpause(vmwarevm string) error {
	cmd := vmware.VMRun(app, "unpause", vmwarevm)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	if runErr := cmd.Run(); runErr != nil {
		if err := runErr.(*exec.ExitError); err != nil {
			return fmt.Errorf(stdout.String())
		}
		return runErr
	}

	return nil
}
