// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fusion

import (
	"github.com/go-vm/vmware"
)

const app = "fusion"

// Start start a VM or Team.
func Start(vmx string, gui bool) error {
	flag := "nogui"
	if gui {
		flag = "gui"
	}

	return vmware.VMRun(app, "start", vmx, flag)
}

// Stop stop a VM or Team.
func Stop(vmx string, hard bool) error {
	flag := "soft"
	if hard {
		flag = "hard"
	}

	return vmware.VMRun(app, "stop", vmx, flag)
}

// Reset reset a VM or Team.
func Reset(vmx string, hard bool) error {
	flag := "soft"
	if hard {
		flag = "hard"
	}

	return vmware.VMRun(app, "reset", vmx, flag)
}

// Suspend Suspend a VM or Team.
func Suspend(vmx string, hard bool) error {
	flag := "soft"
	if hard {
		flag = "hard"
	}

	return vmware.VMRun(app, "suspend", vmx, flag)
}

// Pause pause a VM.
func Pause(vmx string) error {
	return vmware.VMRun(app, "pause", vmx)
}

// Unpause unpause a VM.
func Unpause(vmx string) error {
	return vmware.VMRun(app, "unpause", vmx)
}

// ListSnapshots list all snapshots in a VM.
func ListSnapshots(vmx string) error {
	return vmware.VMRun(app, "listSnapshots", vmx)
}

// Snapshot create a snapshot of a VM.
func Snapshot(vmx, snapshotName string) error {
	return vmware.VMRun(app, "snapshot", vmx, snapshotName)
}

// DeleteSnapshot remove a snapshot from a VM.
func DeleteSnapshot(vmx, snapshotName string, deleteChildren bool) error {
	if deleteChildren {
		return vmware.VMRun(app, "deleteSnapshot", vmx, snapshotName, "andDeleteChildren")
	}
	return vmware.VMRun(app, "deleteSnapshot", vmx, snapshotName)
}

// RevertToSnapshot set VM state to a snapshot.
func RevertToSnapshot(vmx, snapshotName string) error {
	return vmware.VMRun(app, "revertToSnapshot", vmx, snapshotName)
}
