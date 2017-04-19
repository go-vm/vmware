// Copyright 2017 The go-vm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vdiskmanager

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func createTestVDMK(dst string) error {
	return Create(dst, nil)
}

func TestCreate(t *testing.T) {
	type args struct {
		dst    string
		config *Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "defalut",
			args: args{
				dst:    "create.vmdk",
				config: nil,
			},
			wantErr: false,
		},
		{
			name: "declare size",
			args: args{
				dst: "create-size.vmdk",
				config: &Config{
					Size:     50000,
					DiskType: 0,
					Adapter:  LsiLogic,
				},
			},
			wantErr: false,
		},
		{
			name: "declare diskType",
			args: args{
				dst: "create-disktype.vmdk",
				config: &Config{
					DiskType: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "declare adapter to BusLogic",
			args: args{
				dst: "create-disktype-buslogic.vmdk",
				config: &Config{
					Adapter: BusLogic,
				},
			},
			wantErr: false,
		},
		{
			name: "declare adapter to Ide",
			args: args{
				dst: "create-disktype-ide.vmdk",
				config: &Config{
					Adapter: Ide,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Create(tt.args.dst, tt.args.config); (err != nil) != tt.wantErr {
				t.Fatalf("Create(%v, %v) error = %v, wantErr %v", tt.args.dst, tt.args.config, err, tt.wantErr)
			}

			// fallback to .vmdk file extension
			if !strings.HasSuffix(tt.args.dst, ".vmdk") {
				tt.args.dst = tt.args.dst + ".vmdk"
			}

			// check exist test vmdk file
			if _, err := os.Stat(tt.args.dst); (err != nil) != tt.wantErr {
				t.Fatalf("Create(%v, %v) error = %v, wantErr %v", tt.args.dst, tt.args.config, err, tt.wantErr)
			}

			// remove test vmdk files with globbing
			files, _ := filepath.Glob("*.vmdk")
			for _, file := range files {
				os.Remove(file)
			}
		})
	}
}

func TestDefrag(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "normal",
			args:    args{src: "defrag.vmdk"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create the test vdmk
			if err := createTestVDMK(tt.args.src); err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tt.args.src)

			if err := Defrag(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("Defrag(%v) error = %v, wantErr %v", tt.args.src, err, tt.wantErr)
			}
		})
	}
}

func TestShrink(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "normal",
			args:    args{src: "shrink.vmdk"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create the test vdmk
			if err := createTestVDMK(tt.args.src); err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tt.args.src)

			if err := Shrink(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("Shrink(%v) error = %v, wantErr %v", tt.args.src, err, tt.wantErr)
			}
		})
	}
}

func TestRename(t *testing.T) {
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				src: "rename-src.vmdk",
				dst: "rename-dst.vmdk",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create the test vdmk
			if err := createTestVDMK(tt.args.src); err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tt.args.dst)

			if err := Rename(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Fatalf("Rename(%v, %v) error = %v, wantErr %v", tt.args.src, tt.args.dst, err, tt.wantErr)
			}
			// check exist renamed vmdk file
			if _, err := os.Stat(tt.args.dst); (err != nil) != tt.wantErr {
				t.Fatalf("Rename(%v, %v) error = %v, wantErr %v", tt.args.src, tt.args.dst, err, tt.wantErr)
			}
		})
	}
}

// TODO(zchee): implements Prepare test.
// func TestPrepare(t *testing.T) {
// 	type args struct {
// 		src string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name:    "normal",
// 			args:    args{src: "prepare.vmdk"},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// create the test vdmk
// 			if err := createTestVDMK(tt.args.src); err != nil {
// 				t.Fatal(err)
// 			}
// 			defer os.Remove(tt.args.src)
//
// 			if err := Prepare(tt.args.src); (err != nil) != tt.wantErr {
// 				t.Errorf("Prepare(%v) error = %v, wantErr %v", tt.args.src, err, tt.wantErr)
// 			}
// 		})
// 	}
// }

func TestConvert(t *testing.T) {
	type args struct {
		src      string
		dst      string
		diskType int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				src:      "convert-src.vmdk",
				dst:      "convert-dst.vmdk",
				diskType: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create the test vdmk
			if err := createTestVDMK(tt.args.src); err != nil {
				t.Fatal(err)
			}
			// remove test vmdk files with globbing
			defer func() {
				files, _ := filepath.Glob("*.vmdk")
				for _, file := range files {
					os.Remove(file)
				}
			}()

			if err := Convert(tt.args.src, tt.args.dst, tt.args.diskType); (err != nil) != tt.wantErr {
				t.Errorf("Convert(%v, %v, %v) error = %v, wantErr %v", tt.args.src, tt.args.dst, tt.args.diskType, err, tt.wantErr)
			}
		})
	}
}

func TestExpand(t *testing.T) {
	type args struct {
		capacity int
		src      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				capacity: 30000, // grow up default(20000MB) to 30GB
				src:      "expand.vmdk",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create the test vdmk
			if err := createTestVDMK(tt.args.src); err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tt.args.src)

			if err := Expand(tt.args.capacity, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("Expand(%v, %v) error = %v, wantErr %v", tt.args.capacity, tt.args.src, err, tt.wantErr)
			}
		})
	}
}

func TestRepair(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				src: "repair.vmdk",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create the test vdmk
			if err := createTestVDMK(tt.args.src); err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tt.args.src)

			if err := Repair(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("Repair(%v) error = %v, wantErr %v", tt.args.src, err, tt.wantErr)
			}
		})
	}
}

func TestCheck(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				src: "check.vmdk",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create the test vdmk
			if err := createTestVDMK(tt.args.src); err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tt.args.src)

			if err := Check(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("Check(%v) error = %v, wantErr %v", tt.args.src, err, tt.wantErr)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				src: "delete.vmdk",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create the test vdmk
			if err := createTestVDMK(tt.args.src); err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tt.args.src)

			if err := Delete(tt.args.src); (err != nil) != tt.wantErr {
				t.Fatalf("Delete(%v) error = %v, wantErr %v", tt.args.src, err, tt.wantErr)
			}
			if _, err := os.Stat(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("Delete(%v) error = %v, wantErr %v", tt.args.src, err, tt.wantErr)
			}
		})
	}
}
