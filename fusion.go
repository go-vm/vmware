// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vmware

import (
	"github.com/go-vm/vmware/vmrun"
)

const fusionApp = "fusion"

// Fusion represents a VMware Fusion application.
type Fusion struct {
	vmx string
}

// NewFusion return the new Fusion.
func NewFusion(vmx string) *Fusion {
	return &Fusion{
		vmx: vmx,
	}
}

// Start start a VM or Team.
func (f *Fusion) Start(gui bool) error {
	return vmrun.Start(fusionApp, f.vmx, gui)
}

// ShutDown wrap of stop command with hard.
func (f *Fusion) ShutDown() error {
	return vmrun.Stop(fusionApp, f.vmx, true)
}

// Halt wrap of stop command with soft.
func (f *Fusion) Halt() error {
	return vmrun.Stop(fusionApp, f.vmx, false)
}

// Reset reset a VM or Team.
func (f *Fusion) Reset() error {
	return vmrun.Reset(fusionApp, f.vmx, true)
}

// Restart restart a VM uses wrap of reset command with soft.
func (f *Fusion) Restart() error {
	return vmrun.Reset(fusionApp, f.vmx, false)
}

// Suspend Suspend a VM or Team.
func (f *Fusion) Suspend(hard bool) error {
	return vmrun.Suspend(fusionApp, f.vmx, hard)
}

// Pause pause a VM.
func (f *Fusion) Pause() error {
	return vmrun.Pause(fusionApp, f.vmx)
}

// Unpause unpause a VM.
func (f *Fusion) Unpause(vmx string) error {
	return vmrun.Unpause(fusionApp, f.vmx)
}

// ListSnapshots list all snapshots in a VM.
func (f *Fusion) ListSnapshots() ([]string, int, error) {
	return vmrun.ListSnapshots(fusionApp, f.vmx)
}

// Snapshot create a snapshot of a VM.
func (f *Fusion) Snapshot(snapshotName string) error {
	return vmrun.Snapshot(fusionApp, f.vmx, snapshotName)
}

// DeleteSnapshot remove a snapshot from a VM.
func (f *Fusion) DeleteSnapshot(snapshotName string, deleteChildren bool) error {
	return vmrun.DeleteSnapshot(fusionApp, f.vmx, snapshotName, deleteChildren)
}

// RevertToSnapshot set VM state to a snapshot.
func (f *Fusion) RevertToSnapshot(snapshotName string) error {
	return vmrun.RevertToSnapshot(fusionApp, f.vmx, snapshotName)
}

// RunProgramInGuest run a program in Guest OS.
func (f *Fusion) RunProgramInGuest(username, password string, config vmrun.RunProgramInGuestConfig, cmdPath string, cmdArgs ...string) error {
	return vmrun.RunProgramInGuest(fusionApp, f.vmx, username, password, config, cmdPath, cmdArgs...)
}

// FileExistsInGuest check if a file exists in Guest OS.
func (f *Fusion) FileExistsInGuest(username, password string, filename string) bool {
	return vmrun.FileExistsInGuest(fusionApp, f.vmx, username, password, filename)
}

// DirectoryExistsInGuest check if a directory exists in Guest OS.
func (f *Fusion) DirectoryExistsInGuest(username, password string, dir string) bool {
	return vmrun.DirectoryExistsInGuest(fusionApp, f.vmx, username, password, dir)
}

// SetSharedFolderState modify a Host-Guest shared folder.
func (f *Fusion) SetSharedFolderState(shareName, hostPath string, writable bool) bool {
	return vmrun.SetSharedFolderState(fusionApp, f.vmx, shareName, hostPath, writable)
}
