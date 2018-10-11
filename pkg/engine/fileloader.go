package engine

//go:generate go-bindata -nometadata -pkg $GOPACKAGE -prefix ../../parts/ -o templates.go ../../parts/...
//go:generate gofmt -s -l -w templates.go
// fileloader use go-bindata (https://github.com/go-bindata/go-bindata)
// go-bindata is the way we handle embedded files, like binary, template, etc.
