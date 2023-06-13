#!/usr/bin/env sh

##
## Input parameters
##
BINARY=/go-fanfury/${BINARY:-fanfuryd}
ID=${ID:-0}
LOG=${LOG:-go-fanfury.log}

##
## Assert linux binary
##
if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'fanfuryd' E.g.: -e BINARY=fanfuryd_my_test_version"
	exit 1
fi
BINARY_CHECK="$(file "$BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"
if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

##
## Run binary with all parameters
##
export GOFANFURYHOME="/go-fanfury/node${ID}/fanfuryd"

if [ -d "$(dirname "${GOFANFURYHOME}"/"${LOG}")" ]; then
  "${BINARY}" --home "${GOFANFURYHOME}" "$@" | tee "${GOFANFURYHOME}/${LOG}"
else
  "${BINARY}" --home "${GOFANFURYHOME}" "$@"
fi