// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fusion

import (
	"github.com/go-vm/vmware"
)

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
func RunProgramInGuest(vmx string, auth Auth, config RunProgramInGuestConfig, cmdPath string, cmdArgs ...string) error {
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

	if _, err := vmware.VMRun(app, args...); err != nil {
		return err
	}

	return nil
}

// FileExistsInGuest check if a file exists in Guest OS.
func FileExistsInGuest(vmx string, auth Auth, filename string) bool {
	if _, err := vmware.VMRun(app, "-gu", auth.Username, "-gp", auth.Password, "fileExistsInGuest", vmx, filename); err != nil {
		return false
	}

	return true
}

// DirectoryExistsInGuest check if a directory exists in Guest OS.
func DirectoryExistsInGuest(vmx string, auth Auth, dir string) bool {
	if _, err := vmware.VMRun(app, "-gu", auth.Username, "-gp", auth.Password, "directoryExistsInGuest", vmx, dir); err != nil {
		return false
	}

	return true
}
