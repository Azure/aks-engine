#!/bin/bash

set -e -x -o pipefail

GOTAG="1.13.7"
DOCKERARGS="--network=host"

OUTDIR="$(pwd)/_output"
if [ ! -d $OUTDIR ]; then
    mkdir $OUTDIR
fi

cat <<EOF > $OUTDIR/buildcri.sh
set -e -x -o pipefail
export GOOS=windows
export GOARCH=amd64
go get github.com/Microsoft/hcsshim
cd src/github.com/Microsoft/hcsshim
git rev-parse HEAD > /output/hcsshim-revision.txt
cd \$GOPATH
go build -o /output/containerd-shim-runhcs-v1.exe github.com/Microsoft/hcsshim/cmd/containerd-shim-runhcs-v1
mkdir -p src/github.com/containerd
cd src/github.com/containerd
pwd
git clone https://github.com/containerd/containerd.git
cd containerd
git rev-parse HEAD > /output/containerd-revision.txt
make
# make cri-release # should work, but doesn't
cp bin/ctr.exe /output
#cp bin/containerd.exe /output # missing CRI plugin, so build from containerd/cri
cd \$GOPATH
cd src/github.com/containerd
git clone https://github.com/jterry75/cri.git
cd cri
git checkout windows_port
git rev-parse HEAD > /output/cri-revision.txt
make containerd
cp _output/containerd.exe /output
apt update
apt install -y zip
cd /output
zip windows-cri-containerd.zip *.exe *.txt *.toml
rm -f /output/*.exe
rm -f /output/*.txt
EOF
chmod +x $OUTDIR/buildcri.sh

docker run $DOCKERARGS -v $OUTDIR:/output golang:$GOTAG /bin/bash -c /output/buildcri.sh
