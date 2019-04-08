#!/usr/bin/env bash
#
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT license.

set -eo pipefail

PROJECT_NAME="aks-engine"

: "${USE_SUDO:="true"}"
: "${AKSE_INSTALL_DIR:="/usr/local/bin"}"

# initArch discovers the architecture for this system.
initArch() {
  ARCH=$(uname -m)
  case $ARCH in
    armv5*) ARCH="armv5";;
    armv6*) ARCH="armv6";;
    armv7*) ARCH="arm";;
    aarch64) ARCH="arm64";;
    x86) ARCH="386";;
    x86_64) ARCH="amd64";;
    i686) ARCH="386";;
    i386) ARCH="386";;
  esac
}

# initOS discovers the operating system for this system.
initOS() {
  OS=$(uname | tr '[:upper:]' '[:lower:]')

  case "$OS" in
    # Minimalist GNU for Windows
    mingw*) OS='windows';;
  esac
}

# runs the given command as root (detects if we are root already)
runAsRoot() {
  local CMD="$*"

  if [ $EUID -ne 0 ] && [ "$USE_SUDO" = "true" ]; then
    CMD="sudo $CMD"
  fi

  $CMD
}

# verifySupported checks that the os/arch combination is supported for
# binary builds.
verifySupported() {
  local supported="darwin-386\ndarwin-amd64\nlinux-386\nlinux-amd64\nlinux-arm\nlinux-arm64\nlinux-ppc64le\nwindows-386\nwindows-amd64"
  if ! echo "${supported}" | grep -q "${OS}-${ARCH}"; then
    echo "No prebuilt binary for ${OS}-${ARCH}."
    echo "To build from source, go to https://github.com/Azure/aks-engine"
    exit 1
  fi

  if ! type "curl" > /dev/null && ! type "wget" > /dev/null; then
    echo "Either curl or wget is required"
    exit 1
  fi
}

# checkDesiredVersion checks if the desired version is available.
checkDesiredVersion() {
  # Use the GitHub releases webpage for the project to find the desired version for this project.
  local release_url="https://github.com/Azure/aks-engine/releases/${DESIRED_VERSION:-latest}"
  # shellcheck disable=SC2086
  if type "curl" > /dev/null; then
    TAG=$(curl -SsL $release_url | awk '/\/tag\//' | grep -v no-underline | grep "<a href=\"/Azure/aks-engine/releases" | head -n 1 | cut -d '"' -f 2 | awk '{n=split($NF,a,"/");print a[n]}' | awk 'a !~ $0{print}; {a=$0}')
  elif type "wget" > /dev/null; then
    TAG=$(wget -q -O - $release_url | awk '/\/tag\//' | grep -v no-underline | grep "<a href=\"/Azure/aks-engine/releases" | head -n 1 | cut -d '"' -f 2 | awk '{n=split($NF,a,"/");print a[n]}' | awk 'a !~ $0{print}; {a=$0}')
  fi
  if [ "x$TAG" == "x" ]; then
    echo "Cannot determine ${DESIRED_VERSION} tag."
    exit 1
  fi
}

# checkAKSEInstalledVersion checks which version of AKSE is installed and
# if it needs to be changed.
checkAKSEInstalledVersion() {
  if [[ -f "${AKSE_INSTALL_DIR}/${PROJECT_NAME}" ]]; then
    local version
    version=$(aks-engine version | grep 'Version' | cut -d' ' -f2)
    if [[ "$version" == "$TAG" ]]; then
      echo "AKS-Engine ${version} is already ${DESIRED_VERSION:-latest}"
      return 0
    else
      echo "AKS-Engine ${TAG} is available. Changing from version ${version}."
      return 1
    fi
  else
    return 1
  fi
}

# downloadFile downloads the latest binary package and also the checksum
# for that binary.
downloadFile() {
  AKSE_DIST="aks-engine-$TAG-$OS-$ARCH.tar.gz"
  DOWNLOAD_URL="https://github.com/Azure/aks-engine/releases/download/$TAG/$AKSE_DIST"
  AKSE_TMP_ROOT="$(mktemp -dt akse-installer-XXXXXX)"
  AKSE_TMP_FILE="$AKSE_TMP_ROOT/$AKSE_DIST"
  echo "Downloading $DOWNLOAD_URL"
  if type "curl" > /dev/null; then
    curl -SsL "$DOWNLOAD_URL" -o "$AKSE_TMP_FILE"
  elif type "wget" > /dev/null; then
    wget -q -O "$AKSE_TMP_FILE" "$DOWNLOAD_URL"
  fi
}

# installFile vunpacks and installs it.
installFile() {
  AKSE_TMP="$AKSE_TMP_ROOT/$PROJECT_NAME"

  mkdir -p "$AKSE_TMP"
  tar xf "$AKSE_TMP_FILE" -C "$AKSE_TMP"
  AKSE_TMP_BIN="$AKSE_TMP/$PROJECT_NAME-$TAG-$OS-$ARCH"
  echo "Preparing to install $PROJECT_NAME into ${AKSE_INSTALL_DIR}"
  runAsRoot cp "$AKSE_TMP_BIN/$PROJECT_NAME" "$AKSE_INSTALL_DIR"
  echo "$PROJECT_NAME installed into $AKSE_INSTALL_DIR/$PROJECT_NAME"
}

# fail_trap is executed if an error occurs.
fail_trap() {
  result=$?
  if [ "$result" != "0" ]; then
    if [[ -n "$INPUT_ARGUMENTS" ]]; then
      echo "Failed to install $PROJECT_NAME with the arguments provided: $INPUT_ARGUMENTS"
      help
    else
      echo "Failed to install $PROJECT_NAME"
    fi
    echo -e "\tFor support, go to https://github.com/Azure/aks-engine."
  fi
  cleanup
  exit $result
}

# testVersion tests the installed client to make sure it is working.
testVersion() {
  set +e
  if [ "$?" = "1" ]; then
    # shellcheck disable=SC2016
    echo "$PROJECT_NAME not found. Is $AKSE_INSTALL_DIR on your "'$PATH?'
    exit 1
  fi
  set -e
  echo "Run '$PROJECT_NAME version' to test."
}

# help provides possible cli installation arguments
help () {
  echo "Accepted cli arguments are:"
  echo -e "\t[--help|-h ] ->> prints this help"
  echo -e "\t[--version|-v <desired_version>] . When not defined it defaults to latest"
  echo -e "\te.g. --version v0.32.3  or -v latest"
  echo -e "\t[--no-sudo]  ->> install without sudo"
}

# cleanup temporary files
cleanup() {
  if [[ -d "${AKSE_TMP_ROOT:-}" ]]; then
    rm -rf "$AKSE_TMP_ROOT"
  fi
}

# Execution

#Stop execution on any error
trap "fail_trap" EXIT
set -e

# Parsing input arguments (if any)
export INPUT_ARGUMENTS=( "${@}" )
set -u
while [[ $# -gt 0 ]]; do
  case $1 in
    '--version'|-v)
       shift
       if [[ $# -ne 0 ]]; then
           export DESIRED_VERSION="${1}"
       else
           echo -e "Please provide the desired version. e.g. --version v0.32.3 or -v latest"
           exit 0
       fi
       ;;
    '--no-sudo')
       USE_SUDO="false"
       ;;
    '--help'|-h)
       help
       exit 0
       ;;
    *) exit 1
       ;;
  esac
  shift
done
set +u

initArch
initOS
verifySupported
checkDesiredVersion
if ! checkAKSEInstalledVersion; then
  downloadFile
  installFile
fi
testVersion
cleanup
