// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fusion

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

// Authentication represents a guest login data.
type Authentication struct {
	User string
	Pass string
}
