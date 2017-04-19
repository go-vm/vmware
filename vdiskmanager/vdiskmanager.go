// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vdiskmanager

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/go-vm/vmware/internal/vmwareutil"
)

// VMware Virtual Disk Manager - build 5192483.
// Usage: vmware-vdiskmanager OPTIONS <disk-name> | <mount-point>
// Offline disk manipulation utility
//   Operations, only one may be specified at a time:
//      -c                   : create disk.  Additional creation options must
//                             be specified.  Only local virtual disks can be
//                             created.
//      -d                   : defragment the specified virtual disk. Only
//                             local virtual disks may be defragmented.
//      -k                   : shrink the specified virtual disk. Only local
//                             virtual disks may be shrunk.
//      -n <source-disk>     : rename the specified virtual disk; need to
//                             specify destination disk-name. Only local virtual
//                             disks may be renamed.
//      -p                   : prepare the mounted virtual disk specified by
//                             the volume path for shrinking.
//      -r <source-disk>     : convert the specified disk; need to specify
//                             destination disk-type.  For local destination disks
//                             the disk type must be specified.
//      -x <new-capacity>    : expand the disk to the specified capacity. Only
//                             local virtual disks may be expanded.
//      -R                   : check a sparse virtual disk for consistency and attempt
//                             to repair any errors.
//      -e                   : check for disk chain consistency.
//      -D                   : make disk deletable.  This should only be used on disks
//                             that have been copied from another product.
//
//   Other Options:
//      -q                   : do not log messages
//
//   Additional options for create and convert:
//      -a <adapter>         : (for use with -c only) adapter type
//                             (ide, buslogic, lsilogic). Pass lsilogic for other adapter types.
//      -s <size>            : capacity of the virtual disk
//      -t <disk-type>       : disk type id
//
//   Disk types:
//       0                   : single growable virtual disk
//       1                   : growable virtual disk split in 2GB files
//       2                   : preallocated virtual disk
//       3                   : preallocated virtual disk split in 2GB files
//       4                   : preallocated ESX-type virtual disk
//       5                   : compressed disk optimized for streaming
//       6                   : thin provisioned virtual disk - ESX 3.x and above
//
//      The capacity can be specified in sectors, KB, MB or GB.
//      The acceptable ranges:
//                            ide/scsi adapter : [1MB, 8192.0GB]
//                            buslogic adapter : [1MB, 2040.0GB]
//         ex 1: vmware-vdiskmanager -c -s 850MB -a ide -t 0 myIdeDisk.vmdk
//         ex 2: vmware-vdiskmanager -d myDisk.vmdk
//         ex 3: vmware-vdiskmanager -r sourceDisk.vmdk -t 0 destinationDisk.vmdk
//         ex 4: vmware-vdiskmanager -x 36GB myDisk.vmdk
//         ex 5: vmware-vdiskmanager -n sourceName.vmdk destinationName.vmdk
//         ex 6: vmware-vdiskmanager -r sourceDisk.vmdk -t 4 -h esx-name.mycompany.com \
//               -u username -f passwordfile "[storage1]/path/to/targetDisk.vmdk"
//         ex 7: vmware-vdiskmanager -k myDisk.vmdk
//         ex 8: vmware-vdiskmanager -p <mount-point>
//               (A virtual disk first needs to be mounted at <mount-point>)
//

var vdiskmanagerPath = vmwareutil.LookPath("vmware-vdiskmanager")

// vdiskmanager wrapper of vmware-vdiskmanager command.
func vdiskmanager(args ...string) error {
	cmd := exec.Command(vdiskmanagerPath, args...)

	if err := cmd.Run(); err != nil {
		if runErr := err.(*exec.ExitError); runErr != nil {
			return runErr
		}
		return err
	}

	return nil
}

// AdapterType represents a adapter type.
type AdapterType int

const (
	// LsiLogic is a lsilogic type.
	LsiLogic AdapterType = iota
	// Ide is a ide type.
	Ide
	// BusLogic is a buslogic type.
	BusLogic
)

// String implements a fmt.Stringer interface.
func (a AdapterType) String() string {
	switch a {
	case LsiLogic:
		return "lsilogic"
	case Ide:
		return "ide"
	case BusLogic:
		return "buslogic"
	default:
		return ""
	}
}

// Config represents a vdiskmanager create config.
type Config struct {
	Size     int
	DiskType int
	Adapter  AdapterType
}

// Create create disk.
func Create(dst string, config *Config) error {
	size := 20000       // default is 20GB
	diskType := 0       // default is 0
	adapter := LsiLogic // default is lsilogic

	if config != nil {
		if config.Size > 0 {
			size = config.Size
		}
		if config.DiskType > 0 {
			diskType = config.DiskType
		}
		if config.Adapter > 0 {
			adapter = config.Adapter
		}
	}

	if !strings.HasSuffix(dst, ".vmdk") {
		dst = dst + ".vmdk"
	}

	return vdiskmanager("-c", "-s", fmt.Sprintf("%dMB", size), "-t", strconv.Itoa(diskType), "-a", adapter.String(), dst)
}

// Defrag defragment the specified virtual disk.
func Defrag(src string) error {
	return vdiskmanager("-d", src)
}

// Shrink shrink the specified virtual disk.
func Shrink(src string) error {
	return vdiskmanager("-k", src)
}

// Rename rename the specified virtual disk.
func Rename(src, dst string) error {
	return vdiskmanager("-n", src, dst)
}

// Prepare the mounted virtual disk specified by the volume path for shrinking.
func Prepare(src string) error {
	return vdiskmanager("-p", src)
}

// Convert convert the specified disk.
func Convert(src, dst string, diskType int) error {
	return vdiskmanager("-r", src, "-t", strconv.Itoa(diskType), dst)
}

// Expand expand the disk to the specified capacity.
func Expand(capacity int, src string) error {
	return vdiskmanager("-x", fmt.Sprintf("%dMB", capacity), src)
}

// Repair check a sparse virtual disk for consistency and attempt to repair any errors.
func Repair(src string) error {
	return vdiskmanager("-R", src)
}

// Check check for disk chain consistency.
func Check(src string) error {
	return vdiskmanager("-e", src)
}

// Delete make disk deletable.
func Delete(src string) error {
	return vdiskmanager("-D", src)
}
