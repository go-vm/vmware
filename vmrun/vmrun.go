// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package vmrun implements a VMware vmrun command wrapper.
package vmrun

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"

	"github.com/go-vm/vmware/internal/vmwareutil"
)

var vmrunPath = vmwareutil.LookPath("vmrun")

// vmrun run the vmrun command with the app name and args, return the stdout result and cmd error.
func vmrun(app string, arg ...string) (string, error) {
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

// POWER COMMANDS           PARAMETERS           DESCRIPTION
// --------------           ----------           -----------
// start                    Path to vmx file     Start a VM or Team
//                          [gui|nogui]
//
// stop                     Path to vmx file     Stop a VM or Team
//                          [hard|soft]
//
// reset                    Path to vmx file     Reset a VM or Team
//                          [hard|soft]
//
// suspend                  Path to vmx file     Suspend a VM or Team
//                          [hard|soft]
//
// pause                    Path to vmx file     Pause a VM
//
// unpause                  Path to vmx file     Unpause a VM

// Start start a VM or Team.
func Start(app, vmx string, gui bool) error {
	flag := "nogui"
	if gui {
		flag = "gui"
	}

	if _, err := vmrun(app, "start", vmx, flag); err != nil {
		return err
	}

	return nil
}

// Stop stop a VM or Team.
func Stop(app, vmx string, hard bool) error {
	flag := "soft"
	if hard {
		flag = "hard"
	}

	if _, err := vmrun(app, "stop", vmx, flag); err != nil {
		return err
	}

	return nil
}

// Reset reset a VM or Team.
func Reset(app, vmx string, hard bool) error {
	flag := "soft"
	if hard {
		flag = "hard"
	}

	if _, err := vmrun(app, "reset", vmx, flag); err != nil {
		return err
	}

	return nil
}

// Suspend Suspend a VM or Team.
func Suspend(app, vmx string, hard bool) error {
	flag := "soft"
	if hard {
		flag = "hard"
	}

	if _, err := vmrun(app, "suspend", vmx, flag); err != nil {
		return err
	}

	return nil
}

// Pause pause a VM.
func Pause(app, vmx string) error {
	if _, err := vmrun(app, "pause", vmx); err != nil {
		return err
	}

	return nil
}

// Unpause unpause a VM.
func Unpause(app, vmx string) error {
	if _, err := vmrun(app, "unpause", vmx); err != nil {
		return err
	}

	return nil
}

// SNAPSHOT COMMANDS        PARAMETERS           DESCRIPTION
// -----------------        ----------           -----------
// listSnapshots            Path to vmx file     List all snapshots in a VM
//                          [showTree]
//
// snapshot                 Path to vmx file     Create a snapshot of a VM
//                          Snapshot name
//
// deleteSnapshot           Path to vmx file     Remove a snapshot from a VM
//                          Snapshot name
//                          [andDeleteChildren]
//
// revertToSnapshot         Path to vmx file     Set VM state to a snapshot
//                          Snapshot name

var listSnapshotsRe = regexp.MustCompile(`[^Total snapshots: \d](\w+)`)

// ListSnapshots list all snapshots in a VM.
func ListSnapshots(app, vmx string) ([]string, int, error) {
	stdout, err := vmrun(app, "listSnapshots", vmx)
	if err != nil {
		return nil, 0, err
	}

	list := listSnapshotsRe.FindAllString(stdout, -1)
	// TODO(zchee): need re-append with TrimSpace?
	var snapshotList []string
	for _, snapshotName := range list {
		snapshotList = append(snapshotList, strings.TrimSpace(snapshotName))
	}

	return snapshotList, len(snapshotList), nil
}

// Snapshot create a snapshot of a VM.
func Snapshot(app, vmx, snapshotName string) error {
	if _, err := vmrun(app, "snapshot", vmx, snapshotName); err != nil {
		return err
	}

	return nil
}

// DeleteSnapshot remove a snapshot from a VM.
func DeleteSnapshot(app, vmx, snapshotName string, deleteChildren bool) error {
	args := []string{"deleteSnapshot", vmx, snapshotName}
	if deleteChildren {
		args = append(args, "andDeleteChildren")
	}

	if _, err := vmrun(app, args...); err != nil {
		return err
	}

	return nil
}

// RevertToSnapshot set VM state to a snapshot.
func RevertToSnapshot(app, vmx, snapshotName string) error {
	if _, err := vmrun(app, "revertToSnapshot", vmx, snapshotName); err != nil {
		return err
	}

	return nil
}

// GUEST OS COMMANDS        PARAMETERS           DESCRIPTION
// -----------------        ----------           -----------
// runProgramInGuest        Path to vmx file     Run a program in Guest OS
//                          [-noWait]
//                          [-activeWindow]
//                          [-interactive]
//                          Complete-Path-To-Program
//                          [Program arguments]
//
// fileExistsInGuest        Path to vmx file     Check if a file exists in Guest OS
//                          Path to file in guest
//
// directoryExistsInGuest   Path to vmx file     Check if a directory exists in Guest OS
//                          Path to directory in guest
//
// setSharedFolderState     Path to vmx file     Modify a Host-Guest shared folder
//                          Share name
//                          Host path
//                          writable | readonly
//
// addSharedFolder          Path to vmx file     Add a Host-Guest shared folder
//                          Share name
//                          New host path
//
// removeSharedFolder       Path to vmx file     Remove a Host-Guest shared folder
//                          Share name
//
// enableSharedFolders      Path to vmx file     Enable shared folders in Guest
//                          [runtime]
//
// disableSharedFolders     Path to vmx file     Disable shared folders in Guest
//                          [runtime]
//
// listProcessesInGuest     Path to vmx file     List running processes in Guest OS
//
// killProcessInGuest       Path to vmx file     Kill a process in Guest OS
//                          process id
//
// runScriptInGuest         Path to vmx file     Run a script in Guest OS
//                          [-noWait]
//                          [-activeWindow]
//                          [-interactive]
//                          Interpreter path
//                          Script text
//
// deleteFileInGuest        Path to vmx file     Delete a file in Guest OS
//                          Path in guest
//
// createDirectoryInGuest   Path to vmx file     Create a directory in Guest OS
//                          Directory path in guest
//
// deleteDirectoryInGuest   Path to vmx file     Delete a directory in Guest OS
//                          Directory path in guest
//
// CreateTempfileInGuest    Path to vmx file     Create a temporary file in Guest OS
//
// listDirectoryInGuest     Path to vmx file     List a directory in Guest OS
//                          Directory path in guest
//
// CopyFileFromHostToGuest  Path to vmx file     Copy a file from host OS to guest OS
//                          Path on host
//                          Path in guest
//
//
// CopyFileFromGuestToHost  Path to vmx file     Copy a file from guest OS to host OS
//                          Path in guest
//                          Path on host
//
//
// renameFileInGuest        Path to vmx file     Rename a file in Guest OS
//                          Original name
//                          New name
//
// captureScreen            Path to vmx file     Capture the screen of the VM to a local file
//                          Path on host
//
// writeVariable            Path to vmx file     Write a variable in the VM state
//                          [runtimeConfig|guestEnv|guestVar]
//                          variable name
//                          variable value
//
// readVariable             Path to vmx file     Read a variable in the VM state
//                          [runtimeConfig|guestEnv|guestVar]
//                          variable name
//
// getGuestIPAddress        Path to vmx file     Gets the IP address of the guest
//                          [-wait]

// RunInGuestConfig represents a runProgramInGuest and runScriptInGuest commands flag.
type RunInGuestConfig int

const (
	// NoWait returns a prompt immediately after the program starts in the guest, rather than waiting for it to finish.
	// This option is useful for interactive programs.
	NoWait RunInGuestConfig = 1 << iota
	// ActiveWindow ensures that the Windows GUI is visible, not minimized.
	// It has no effect on Linux.
	ActiveWindow
	// Interactive forces interactive guest login.
	// It is useful for Vista and Windows 7 guests to make the program visible in the console window.
	Interactive
)

// String implements a fmt.Stringer interface.
func (r RunInGuestConfig) String() string {
	switch r {
	case NoWait:
		return "-noWait"
	case ActiveWindow:
		return "-activeWindow"
	case Interactive:
		return "-interactive"
	default:
		return ""
	}
}

// RunProgramInGuest run a program in Guest OS.
func RunProgramInGuest(app, vmx string, username, password string, config RunInGuestConfig, cmdPath string, cmdArgs ...string) error {
	args := []string{"-gu", username, "-gp", password, "runProgramInGuest", vmx}

	if config&NoWait > 0 {
		args = append(args, NoWait.String())
	}
	if config&ActiveWindow > 0 {
		args = append(args, ActiveWindow.String())
	}
	if config&Interactive > 0 {
		args = append(args, Interactive.String())
	}

	args = append(args, cmdPath)
	args = append(args, cmdArgs...)

	if _, err := vmrun(app, args...); err != nil {
		return err
	}

	return nil
}

// FileExistsInGuest check if a file exists in Guest OS.
func FileExistsInGuest(app, vmx string, username, password string, filename string) bool {
	if _, err := vmrun(app, "-gu", username, "-gp", password, "fileExistsInGuest", vmx, filename); err != nil {
		return false
	}

	return true
}

// DirectoryExistsInGuest check if a directory exists in Guest OS.
func DirectoryExistsInGuest(app, vmx string, username, password string, dir string) bool {
	if _, err := vmrun(app, "-gu", username, "-gp", password, "directoryExistsInGuest", vmx, dir); err != nil {
		return false
	}

	return true
}

// SetSharedFolderState modify a Host-Guest shared folder.
func SetSharedFolderState(app, vmx string, shareName, hostPath string, writable bool) error {
	flag := "readonly"
	if writable {
		flag = "writable"
	}
	if _, err := vmrun(app, "setSharedFolderState", vmx, shareName, hostPath, flag); err != nil {
		return err
	}

	return nil
}

// AddSharedFolder add a Host-Guest shared folder.
func AddSharedFolder(app, vmx string, shareName, newHostPath string) error {
	if _, err := vmrun(app, "addSharedFolder", vmx, shareName, newHostPath); err != nil {
		return err
	}

	return nil
}

// RemoveSharedFolder remove a Host-Guest shared folder.
func RemoveSharedFolder(app, vmx, shareName string) error {
	if _, err := vmrun(app, "removeSharedFolder", vmx, shareName); err != nil {
		return err
	}

	return nil
}

// EnableSharedFolders enable shared folders in Guest.
//
// The optional runtime argument means to share folders only until the virtual machine is powered off.
// Otherwise, the setting persists at next power on.
func EnableSharedFolders(app, vmx string, runtime bool) error {
	args := []string{"enableSharedFolders", vmx}
	if runtime {
		args = append(args, "runtime")
	}

	if _, err := vmrun(app, args...); err != nil {
		return err
	}

	return nil
}

// DisableSharedFolders disable shared folders in Guest.
// Stops the guest virtual machine, specified by .vmx file, from sharing folders with its host.
//
// The optional runtime argument means to stop sharing folders only until the virtual machine is powered off.
// Otherwise, the setting persists at next power on.
func DisableSharedFolders(app, vmx string, runtime bool) error {
	args := []string{"disableSharedFolders", vmx}
	if runtime {
		args = append(args, "runtime")
	}

	if _, err := vmrun(app, args...); err != nil {
		return err
	}

	return nil
}

// ListProcessesInGuestInfo represents a result of listprocessesinguest command.
type ListProcessesInGuestInfo struct {
	Pid   string
	Owner string
	Cmd   string
}

var listProcessesInGuestRe = regexp.MustCompile(`pid=(\d+), owner=(\w+), cmd=([[:print:]]+)`)

// ListProcessesInGuest List running processes in Guest OS.
func ListProcessesInGuest(app, vmx, username, password string) ([]ListProcessesInGuestInfo, error) {
	stdout, err := vmrun(app, "-gu", username, "-gp", password, "listprocessesinguest", vmx)
	if err != nil {
		return nil, err
	}

	cmdList := listProcessesInGuestRe.FindAllStringSubmatch(stdout, -1)

	var list []ListProcessesInGuestInfo
	for _, cmd := range cmdList {
		list = append(list, ListProcessesInGuestInfo{
			Pid:   cmd[1],
			Owner: cmd[2],
			Cmd:   cmd[3],
		})
	}

	return list, nil
}

// KillProcessInGuest kill a process in Guest OS.
func KillProcessInGuest(app, vmx, username, password string, pid int) error {
	if _, err := vmrun(app, "-gu", username, "-gp", password, "killprocessinguest", vmx, strconv.Itoa(pid)); err != nil {
		return err
	}

	return nil
}

// RunScriptInGuest run a script in Guest OS.
func RunScriptInGuest(app, vmx, username, password string, config RunInGuestConfig, interpreter, script string) error {
	args := []string{"-gu", username, "-gp", password, "runScriptInGuest", vmx}

	if config&NoWait > 0 {
		args = append(args, NoWait.String())
	}
	if config&ActiveWindow > 0 {
		args = append(args, ActiveWindow.String())
	}
	if config&Interactive > 0 {
		args = append(args, Interactive.String())
	}

	args = append(args, interpreter, script)

	if _, err := vmrun(app, args...); err != nil {
		return err
	}

	return nil
}

// DeleteFileInGuest delete a file in Guest OS.
func DeleteFileInGuest(app, vmx, username, password, filename string) error {
	if _, err := vmrun(app, "-gu", username, "-gp", password, "deleteFileInGuest", vmx, filename); err != nil {
		return err
	}

	return nil
}

// CreateTempfileInGuest create a temporary file in Guest OS.
func CreateTempfileInGuest(app, vmx, username, password string) (string, error) {
	stdout, err := vmrun(app, "-gu", username, "-gp", password, "CreateTempfileInGuest", vmx)
	if err != nil {
		return "", err
	}

	return filepath.Clean(stdout), nil
}

var listDirectoryInGuestRe = regexp.MustCompile(`[^Directory list: \d](\w+)`)

// ListDirectoryInGuest list a directory in Guest OS.
func ListDirectoryInGuest(app, vmx, username, password, dir string) ([]string, error) {
	stdout, err := vmrun(app, "-gu", username, "-gp", password, "listDirectoryInGuest", vmx, dir)
	if err != nil {
		return nil, err
	}

	files := listDirectoryInGuestRe.FindAllString(stdout, -1)

	return files, nil
}

// CopyFileFromHostToGuest copy a file from host OS to guest OS.
func CopyFileFromHostToGuest(app, vmx, username, password, hostFilepath, guestFilepath string) error {
	if _, err := vmrun(app, "-gu", username, "-gp", password, "CopyFileFromHostToGuest", vmx, hostFilepath, guestFilepath); err != nil {
		return err
	}

	return nil
}

// CopyFileFromGuestToHost copy a file from guest OS to host OS.
func CopyFileFromGuestToHost(app, vmx, username, password, guestFilepath, hostFilepath string) error {
	if _, err := vmrun(app, "-gu", username, "-gp", password, "CopyFileFromGuestToHost", vmx, guestFilepath, hostFilepath); err != nil {
		return err
	}

	return nil
}

// RenameFileInGuest rename a file in Guest OS.
func RenameFileInGuest(app, vmx, username, password, src, dst string) error {
	if _, err := vmrun(app, "-gu", username, "-gp", password, "renameFileInGuest", vmx, src, dst); err != nil {
		return err
	}

	return nil
}

// CaptureScreen capture the screen of the VM to a local file.
func CaptureScreen(app, vmx, username, password, dst string) error {
	if _, err := vmrun(app, "-gu", username, "-gp", password, "captureScreen", vmx, dst); err != nil {
		return err
	}

	return nil
}

// VariableMode represents a writeVariable or readVariable command mode.
type VariableMode int

const (
	// RuntimeConfig runtime‐only value that provides a simple way to pass runtime values in and out of the guest.
	RuntimeConfig VariableMode = 1 << iota
	// GuestEnv non‐persistent guest variable.
	GuestEnv
	// GuestVar runtime configuration parameter as stored in the .vmx file, or an environment variable.
	GuestVar
)

// String implements a fmt.Stringer interface.
func (v VariableMode) String() string {
	switch v {
	case RuntimeConfig:
		return "runtimeConfig"
	case GuestEnv:
		return "guestEnv"
	case GuestVar:
		return "guestVar"
	default:
		return ""
	}
}

// WriteVariable write a variable in the VM state.
func WriteVariable(app, vmx, username, password string, mode VariableMode, env, value string) error {
	if _, err := vmrun(app, "-gu", username, "-gp", password, "writeVariable", vmx, mode.String(), env, value); err != nil {
		return err
	}

	return nil
}
