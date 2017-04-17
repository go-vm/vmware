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
func Start(vmx string, gui bool) error {
	flag := "nogui"
	if gui {
		flag = "gui"
	}

	cmd := vmware.VMRun(app, "start", vmx, flag)
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
func Stop(vmx string, hard bool) error {
	flag := "soft"
	if hard {
		flag = "hard"
	}

	cmd := vmware.VMRun(app, "stop", vmx, flag)
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
func Reset(vmx string, hard bool) error {
	flag := "soft"
	if hard {
		flag = "hard"
	}

	cmd := vmware.VMRun(app, "reset", vmx, flag)
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
func Suspend(vmx string, hard bool) error {
	flag := "soft"
	if hard {
		flag = "hard"
	}

	cmd := vmware.VMRun(app, "suspend", vmx, flag)
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
func Pause(vmx string) error {
	cmd := vmware.VMRun(app, "pause", vmx)

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
func Unpause(vmx string) error {
	cmd := vmware.VMRun(app, "unpause", vmx)

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

// ListSnapshots list all snapshots in a VM.
func ListSnapshots(vmx string) error {
	cmd := vmware.VMRun(app, "listSnapshots", vmx)

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

// Snapshot create a snapshot of a VM.
func Snapshot(vmx, snapshotName string) error {
	cmd := vmware.VMRun(app, "snapshot", vmx, snapshotName)

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

// DeleteSnapshot remove a snapshot from a VM.
func DeleteSnapshot(vmx, snapshotName string, deleteChildren bool) error {
	cmd := vmware.VMRun(app, "deleteSnapshot", vmx, snapshotName)
	if deleteChildren {
		cmd.Args = append(cmd.Args, "andDeleteChildren")
	}

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

// RevertToSnapshot set VM state to a snapshot.
func RevertToSnapshot(vmx, snapshotName string) error {
	cmd := vmware.VMRun(app, "revertToSnapshot", vmx, snapshotName)

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
