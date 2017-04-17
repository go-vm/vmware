// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "github.com/go-vm/vmware/fusion"

const vmwarevmPath = "/Volumes/APFS/VirtualMachine/macOS-10.12.vmwarevm/macOS-10.12.vmx"

func main() {
	if err := fusion.Start(vmwarevmPath, false); err != nil {
		panic(err.Error())
	}

	if err := fusion.Stop(vmwarevmPath, true); err != nil {
		panic(err.Error())
	}
}
