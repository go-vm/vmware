// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vmrun

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"syscall"

	"github.com/go-vm/vmware/internal/vmwareutil"
)

var vmrunPath = vmwareutil.LookPath("vmrun")

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

// Auth represents a guest login data.
type Auth struct {
	HostName          string // -h
	HostPort          string // -P
	HostUserName      string // -u
	HostPassword      string // -p
	EncryptedPassword string // -vp
	Username          string // -gu
	Password          string // -gp
}

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

// ShutDown wrap of stop command with hard.
func ShutDown(app, vmx string) error {
	return Stop(app, vmx, true)
}

// Halt wrap of stop command with soft.
func Halt(app, vmx string) error {
	return Stop(app, vmx, false)
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

// Restart restart a VM uses wrap of reset command with soft.
func Restart(app, vmx string) error {
	return Reset(app, vmx, false)
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
// Path in guest
//
// createDirectoryInGuest   Path to vmx file     Create a directory in Guest OS
// Directory path in guest
//
// deleteDirectoryInGuest   Path to vmx file     Delete a directory in Guest OS
// Directory path in guest
//
// CreateTempfileInGuest    Path to vmx file     Create a temporary file in Guest OS
//
// listDirectoryInGuest     Path to vmx file     List a directory in Guest OS
//                          Directory path in guest
//
// CopyFileFromHostToGuest  Path to vmx file     Copy a file from host OS to guest OS
// Path on host             Path in guest
//
//
// CopyFileFromGuestToHost  Path to vmx file     Copy a file from guest OS to host OS
// Path in guest            Path on host
//
//
// renameFileInGuest        Path to vmx file     Rename a file in Guest OS
//                          Original name
//                          New name
//
// captureScreen            Path to vmx file     Capture the screen of the VM to a local file
// Path on host
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

// RunProgramInGuestConfig represents a runProgramInGuest command flags.
type RunProgramInGuestConfig int

const (
	// NoWait returns a prompt immediately after the program starts in the guest, rather than waiting for it to finish.
	// This option is useful for interactive programs.
	NoWait RunProgramInGuestConfig = 1 << iota
	// ActiveWindow ensures that the Windows GUI is visible, not minimized.
	// It has no effect on Linux.
	ActiveWindow
	// Interactive forces interactive guest login.
	// It is useful for Vista and Windows 7 guests to make the program visible in the console window.
	Interactive
)

// RunProgramInGuest run a program in Guest OS.
func RunProgramInGuest(app, vmx string, auth Auth, config RunProgramInGuestConfig, cmdPath string, cmdArgs ...string) error {
	args := []string{"-gu", auth.Username, "-gp", auth.Password, "runProgramInGuest", vmx}

	if config&NoWait > 0 {
		args = append(args, "-noWait")
	}
	if config&ActiveWindow > 0 {
		args = append(args, "-activeWindow")
	}
	if config&Interactive > 0 {
		args = append(args, "-interactive")
	}

	args = append(args, cmdPath)
	args = append(args, cmdArgs...)

	if _, err := vmrun(app, args...); err != nil {
		return err
	}

	return nil
}

// FileExistsInGuest check if a file exists in Guest OS.
func FileExistsInGuest(app, vmx string, auth Auth, filename string) bool {
	if _, err := vmrun(app, "-gu", auth.Username, "-gp", auth.Password, "fileExistsInGuest", vmx, filename); err != nil {
		return false
	}

	return true
}

// DirectoryExistsInGuest check if a directory exists in Guest OS.
func DirectoryExistsInGuest(app, vmx string, auth Auth, dir string) bool {
	if _, err := vmrun(app, "-gu", auth.Username, "-gp", auth.Password, "directoryExistsInGuest", vmx, dir); err != nil {
		return false
	}

	return true
}

// SetSharedFolderState modify a Host-Guest shared folder.
func SetSharedFolderState(app, vmx string, auth Auth, shareName, hostPath string, writable bool) bool {
	flag := "readonly"
	if writable {
		flag = "writable"
	}
	if _, err := vmrun(app, "-gu", auth.Username, "-gp", auth.Password, "setSharedFolderState", vmx, shareName, hostPath, flag); err != nil {
		return false
	}

	return true
}
