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
	vmx      string
	username string
	password string
}

// NewFusion return the new Fusion.
func NewFusion(vmx, username, password string) *Fusion {
	return &Fusion{
		vmx:      vmx,
		username: username,
		password: password,
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
func (f *Fusion) RunProgramInGuest(config vmrun.RunInGuestConfig, cmdPath string, cmdArgs ...string) error {
	return vmrun.RunProgramInGuest(fusionApp, f.vmx, f.username, f.password, config, cmdPath, cmdArgs...)
}

// FileExistsInGuest check if a file exists in Guest OS.
func (f *Fusion) FileExistsInGuest(filename string) bool {
	return vmrun.FileExistsInGuest(fusionApp, f.vmx, f.username, f.password, filename)
}

// DirectoryExistsInGuest check if a directory exists in Guest OS.
func (f *Fusion) DirectoryExistsInGuest(dir string) bool {
	return vmrun.DirectoryExistsInGuest(fusionApp, f.vmx, f.username, f.password, dir)
}

// SetSharedFolderState modify a Host-Guest shared folder.
func (f *Fusion) SetSharedFolderState(shareName, hostPath string, writable bool) error {
	return vmrun.SetSharedFolderState(fusionApp, f.vmx, shareName, hostPath, writable)
}

// AddSharedFolder add a Host-Guest shared folder.
func (f *Fusion) AddSharedFolder(shareName, newHostPath string) error {
	return vmrun.AddSharedFolder(fusionApp, f.vmx, shareName, newHostPath)
}

// RemoveSharedFolder remove a Host-Guest shared folder.
func (f *Fusion) RemoveSharedFolder(shareName string) error {
	return vmrun.RemoveSharedFolder(fusionApp, f.vmx, shareName)
}

// EnableSharedFolders enable shared folders in Guest.
func (f *Fusion) EnableSharedFolders(runtime bool) error {
	return vmrun.EnableSharedFolders(fusionApp, f.vmx, runtime)
}

// DisableSharedFolders disable shared folders in Guest.
func (f *Fusion) DisableSharedFolders(runtime bool) error {
	return vmrun.DisableSharedFolders(fusionApp, f.vmx, runtime)
}

// ListProcessesInGuest List running processes in Guest OS.
func (f *Fusion) ListProcessesInGuest() ([]vmrun.ListProcessesInGuestInfo, error) {
	return vmrun.ListProcessesInGuest(fusionApp, f.vmx, f.username, f.password)
}

// KillProcessInGuest kill a process in Guest OS.
func (f *Fusion) KillProcessInGuest(pid int) error {
	return vmrun.KillProcessInGuest(fusionApp, f.vmx, f.username, f.password, pid)
}

// RunScriptInGuest run a script in Guest OS.
func (f *Fusion) RunScriptInGuest(config vmrun.RunInGuestConfig, interpreter, script string) error {
	return vmrun.RunScriptInGuest(fusionApp, f.vmx, f.username, f.password, config, interpreter, script)
}

// DeleteFileInGuest delete a file in Guest OS.
func (f *Fusion) DeleteFileInGuest(filename string) error {
	return vmrun.DeleteFileInGuest(fusionApp, f.vmx, f.username, f.password, filename)
}

// CreateTempfileInGuest create a temporary file in Guest OS.
func (f *Fusion) CreateTempfileInGuest() (string, error) {
	return vmrun.CreateTempfileInGuest(fusionApp, f.vmx, f.username, f.password)
}

// ListDirectoryInGuest list a directory in Guest OS.
func (f *Fusion) ListDirectoryInGuest(dir string) ([]string, error) {
	return vmrun.ListDirectoryInGuest(fusionApp, f.vmx, f.username, f.password, dir)
}

// CopyFileFromHostToGuest copy a file from host OS to guest OS.
func (f *Fusion) CopyFileFromHostToGuest(hostFilepath, guestFilepath string) error {
	return vmrun.CopyFileFromHostToGuest(fusionApp, f.vmx, f.username, f.password, hostFilepath, guestFilepath)
}

// CopyFileFromGuestToHost copy a file from guest OS to host OS.
func (f *Fusion) CopyFileFromGuestToHost(guestFilepath, hostFilepath string) error {
	return vmrun.CopyFileFromGuestToHost(fusionApp, f.vmx, f.username, f.password, guestFilepath, hostFilepath)
}

// RenameFileInGuest rename a file in Guest OS.
func (f *Fusion) RenameFileInGuest(src, dst string) error {
	return vmrun.RenameFileInGuest(fusionApp, f.vmx, f.username, f.password, src, dst)
}

// CaptureScreen capture the screen of the VM to a local file.
func (f *Fusion) CaptureScreen(dst string) error {
	return vmrun.CaptureScreen(fusionApp, f.vmx, f.username, f.password, dst)
}

// WriteVariable write a variable in the VM state.
func (f *Fusion) WriteVariable(mode vmrun.VariableMode, env, value string) error {
	return vmrun.WriteVariable(fusionApp, f.vmx, f.username, f.password, mode, env, value)
}

// ReadVariable read a variable in the VM state.
func (f *Fusion) ReadVariable(mode vmrun.VariableMode, env string) (string, error) {
	return vmrun.ReadVariable(fusionApp, f.vmx, f.username, f.password, mode, env)
}
