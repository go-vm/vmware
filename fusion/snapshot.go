// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fusion

import (
	"regexp"
	"strings"

	"github.com/go-vm/vmware"
)

var listSnapshotsRe = regexp.MustCompile(`[^Total snapshots: \d](\w+)`)

// ListSnapshots list all snapshots in a VM.
func ListSnapshots(vmx string) ([]string, int, error) {
	stdout, err := vmware.VMRun(app, "listSnapshots", vmx)
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
