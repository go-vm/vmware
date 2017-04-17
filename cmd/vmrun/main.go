// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/go-vm/vmware/fusion"
)

const vmwarevmPath = "/Volumes/APFS/VirtualMachine/macOS-10.12.vmwarevm/macOS-10.12.vmx"

func main() {
	if err := fusion.Start(vmwarevmPath, false); err != nil {
		panic(err.Error())
	}

	if err := fusion.Stop(vmwarevmPath, true); err != nil {
		panic(err.Error())
	}

	list, num, err := fusion.ListSnapshots(vmwarevmPath)
	if err != nil {
		panic(err.Error())
	}
	log.Printf("list: %T => %+v\n", list, list)
	log.Printf("num: %T => %+v\n", num, num)

	if err := fusion.Snapshot(vmwarevmPath, "testSnapshot"); err != nil {
		panic(err.Error())
	}

	list2, num2, err := fusion.ListSnapshots(vmwarevmPath)
	if err != nil {
		panic(err.Error())
	}
	log.Printf("list: %T => %+v\n", list2, list2)
	log.Printf("num: %T => %+v\n", num2, num2)

	if err := fusion.DeleteSnapshot(vmwarevmPath, "testSnapshot", true); err != nil {
		panic(err.Error())
	}

	list3, num3, err := fusion.ListSnapshots(vmwarevmPath)
	if err != nil {
		panic(err.Error())
	}
	log.Printf("list: %T => %+v\n", list3, list3)
	log.Printf("num: %T => %+v\n", num3, num3)
}
