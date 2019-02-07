// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package i18n

//go:generate go-bindata -nometadata -pkg $GOPACKAGE -prefix ../../ -o translations_generated.go ../../translations/...
//go:generate gofmt -s -l -w translations_generated.go
// resourceloader use go-bindata (https://github.com/go-bindata/go-bindata)
// go-bindata is the way we handle embedded files, like binary, template, etc.
