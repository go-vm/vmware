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
	cmd := vmware.VMRun(app, "start", vmwarevm)

	if gui {
		cmd.Args = append(cmd.Args, "gui")
	} else {
		cmd.Args = append(cmd.Args, "nogui")
	}

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		if e := err.(*exec.ExitError); e != nil {
			return fmt.Errorf(stdout.String())
		}
		return err
	}

	return nil
}

// Stop stop a VM or Team.
func Stop(vmwarevm string, force bool) error {
	cmd := vmware.VMRun(app, "stop", vmwarevm)

	if force {
		cmd.Args = append(cmd.Args, "hard")
	} else {
		cmd.Args = append(cmd.Args, "soft")
	}

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		if e := err.(*exec.ExitError); e != nil {
			return fmt.Errorf(stdout.String())
		}
		return err
	}

	return nil
}

// Reset reset a VM or Team.
func Reset(vmwarevm string, force bool) error {
	cmd := vmware.VMRun(app, "reset", vmwarevm)

	if force {
		cmd.Args = append(cmd.Args, "hard")
	} else {
		cmd.Args = append(cmd.Args, "soft")
	}

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		if e := err.(*exec.ExitError); e != nil {
			return fmt.Errorf(stdout.String())
		}
		return err
	}

	return nil
}

// Suspend Suspend a VM or Team.
func Suspend(vmwarevm string, force bool) error {
	cmd := vmware.VMRun(app, "suspend", vmwarevm)

	if force {
		cmd.Args = append(cmd.Args, "hard")
	} else {
		cmd.Args = append(cmd.Args, "soft")
	}

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		if e := err.(*exec.ExitError); e != nil {
			return fmt.Errorf(stdout.String())
		}
		return err
	}

	return nil
}

// Pause pause a VM.
func Pause(vmwarevm string) error {
	cmd := vmware.VMRun(app, "pause", vmwarevm)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		if e := err.(*exec.ExitError); e != nil {
			return fmt.Errorf(stdout.String())
		}
		return err
	}

	return nil
}

// Unpause unpause a VM.
func Unpause(vmwarevm string) error {
	cmd := vmware.VMRun(app, "unpause", vmwarevm)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		if e := err.(*exec.ExitError); e != nil {
			return fmt.Errorf(stdout.String())
		}
		return err
	}

	return nil
}
