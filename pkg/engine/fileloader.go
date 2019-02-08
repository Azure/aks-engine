// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

//go:generate go-bindata -nometadata -pkg $GOPACKAGE -prefix ../../parts/ -o templates_generated.go ../../parts/...
//go:generate gofmt -s -l -w templates_generated.go
// fileloader use go-bindata (https://github.com/go-bindata/go-bindata)
// go-bindata is the way we handle embedded files, like binary, template, etc.
