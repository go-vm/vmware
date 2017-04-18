// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/go-vm/vmware"
	"github.com/go-vm/vmware/vmrun"
)

const vmxPath = "/Volumes/APFS/VirtualMachine/macOS-10.12.vmwarevm/macOS-10.12.vmx"

func main() {
	fusion := vmware.NewFusion(vmxPath, "darwinstrap", "darwinstrap")

	exist := fusion.DirectoryExistsInGuest("/Volumes/VMware Tools")
	log.Printf("exist: %T => %+v\n", exist, exist)

	if err := fusion.Start(false); err != nil {
		panic(err.Error())
	}

	if _, err := fusion.ListProcessesInGuest(); err != nil {
		panic(err.Error())
	}

	if err := fusion.RunProgramInGuest(vmrun.ActiveWindow, "/usr/bin/env"); err != nil {
		panic(err.Error())
	}

	if err := fusion.DisableSharedFolders(false); err != nil {
		panic(err.Error())
	}

	if err := fusion.Halt(); err != nil {
		panic(err.Error())
	}

	list, num, err := fusion.ListSnapshots()
	if err != nil {
		panic(err.Error())
	}
	log.Printf("list: %T => %+v\n", list, list)
	log.Printf("num: %T => %+v\n", num, num)

	if err := fusion.Snapshot("testSnapshot"); err != nil {
		panic(err.Error())
	}

	list2, num2, err := fusion.ListSnapshots()
	if err != nil {
		panic(err.Error())
	}
	log.Printf("list: %T => %+v\n", list2, list2)
	log.Printf("num: %T => %+v\n", num2, num2)

	if err := fusion.DeleteSnapshot("testSnapshot", true); err != nil {
		panic(err.Error())
	}

	list3, num3, err := fusion.ListSnapshots()
	if err != nil {
		panic(err.Error())
	}
	log.Printf("list: %T => %+v\n", list3, list3)
	log.Printf("num: %T => %+v\n", num3, num3)
}
