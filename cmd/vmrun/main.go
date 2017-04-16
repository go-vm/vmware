// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "github.com/go-vm/vmware/fusion"

func main() {
	err := fusion.Start("/Volumes/APFS/VirtualMachine/macOS-10.12.vmwarevm", true)
	if err != nil {
		panic(err.Error())
	}

	err = fusion.Stop("/Volumes/APFS/VirtualMachine/macOS-10.12.vmwarevm", true)
	if err != nil {
		panic(err.Error())
	}
}
