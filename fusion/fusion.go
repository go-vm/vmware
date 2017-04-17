// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fusion

import (
	"regexp"
	"strings"

	"github.com/go-vm/vmware"
)

const app = "fusion"

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

var listSnapshotsRe = regexp.MustCompile(`[^Total snapshots: \d](\w+)`)

// ListSnapshots list all snapshots in a VM.
func ListSnapshots(vmx string) ([]string, int, error) {
	stdout, err := vmware.VMRun(app, "listSnapshots", vmx)
	if err != nil {
		return nil, 0, err
	}

	list := listSnapshotsRe.FindAllString(stdout, -1)
	// TODO: need re-append with TrimSpace?
	var snapshotList []string
	for _, snapshotName := range list {
		snapshotList = append(snapshotList, strings.TrimSpace(snapshotName))
	}

	return snapshotList, len(snapshotList), nil
}

// Snapshot create a snapshot of a VM.
func Snapshot(vmx, snapshotName string) error {
	if _, err := vmware.VMRun(app, "snapshot", vmx, snapshotName); err != nil {
		return err
	}

	return nil
}

// DeleteSnapshot remove a snapshot from a VM.
func DeleteSnapshot(vmx, snapshotName string, deleteChildren bool) error {
	args := []string{"deleteSnapshot", vmx, snapshotName}
	if deleteChildren {
		args = append(args, "andDeleteChildren")
	}

	if _, err := vmware.VMRun(app, args...); err != nil {
		return err
	}

	return nil
}

// RevertToSnapshot set VM state to a snapshot.
func RevertToSnapshot(vmx, snapshotName string) error {
	if _, err := vmware.VMRun(app, "revertToSnapshot", vmx, snapshotName); err != nil {
		return err
	}

	return nil
}
