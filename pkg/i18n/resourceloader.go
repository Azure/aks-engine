// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package i18n

//go:generate go-bindata -nometadata -pkg $GOPACKAGE -prefix ../../ -o translations.go ../../translations/...
//go:generate gofmt -s -l -w translations.go
// resourceloader use go-bindata (https://github.com/go-bindata/go-bindata)
// go-bindata is the way we handle embedded files, like binary, template, etc.
