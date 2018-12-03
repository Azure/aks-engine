// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package unstable

import (
	"fmt"
	"net"
	"os"
	"strings"
	"testing"

	"github.com/Azure/aks-engine/pkg/openshift/filesystem"
)

type fakefilesystem map[string]filesystem.Fileinfo

func (f fakefilesystem) WriteFile(filename string, data []byte, fi filesystem.Fileinfo) error {
	f[filename] = fi
	// fmt.Printf("Filename: %v	User: %s	Group: %s	Permissions: %04o   string: %v\n", filename, fi.User, fi.Group, fi.Mode, fi.Mode.String())
	return nil
}

func (f fakefilesystem) Mkdir(filename string, fi filesystem.Fileinfo) error {
	// fmt.Printf("Filename: %v	User: %s	Group: %s	Permissions: %04o   string: %v\n", filename, fi.User, fi.Group, fi.Mode, fi.Mode.String())
	f[filename] = fi
	return nil
}

func (fakefilesystem) Close() error {
	return nil
}

var _ filesystem.Writer = &fakefilesystem{}

func TestConfigFilePermissions(t *testing.T) {
	c := Config{
		Master: &Master{
			Hostname: fmt.Sprintf("%s-master-%s-0", "test", "test"),
			IPs: []net.IP{
				net.ParseIP("10.0.0.1"),
			},
		},
	}

	err := c.PrepareMasterCerts()
	if err != nil {
		t.Fatal(err)
	}
	err = c.PrepareMasterKubeConfigs()
	if err != nil {
		t.Fatal(err)
	}
	err = c.PrepareMasterFiles()
	if err != nil {
		t.Fatal(err)
	}

	err = c.PrepareBootstrapKubeConfig()
	if err != nil {
		t.Fatal(err)
	}

	// create mock filesystem
	fs := fakefilesystem{}

	err = c.WriteMaster(fs)
	if err != nil {
		t.Fatal(err)
	}

	for fname, finfo := range fs {
		fi := GetFileInfo(fname)
		// fmt.Printf("fname=>[%s]\n", fname)
		// Verify ownership and permissions are as expected
		if fi.User != finfo.User {
			t.Errorf("File: %s  User does not match.   user: %s  expected: %s", fname, finfo.User, fi.User)
		}
		if fi.Group != finfo.Group {
			t.Errorf("File: %s  Group does not match.  group: %s  expected: %s", fname, finfo.Group, fi.Group)
		}
		if fi.Mode != 0 && fi.Mode != finfo.Mode {
			t.Errorf("File: %s  Mode does not match.   mode: %04o  expected: %04o", fname, finfo.Mode, fi.Mode)
		}
		// Check for .key does _not_ have read or write
		if strings.HasSuffix(fname, ".key") && finfo.Mode&077 != 0 {
			t.Errorf("File: %s  Found Read or Write on key file. mode: %04o  expected: 0600", fname, finfo.Mode)
		}
		if fname == "tmp" && fi.Mode != os.FileMode(01777) {
			t.Errorf("File: %s  /tmp should have 1777 file mode.", fname)
		}
	}
}
