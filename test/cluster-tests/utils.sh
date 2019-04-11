#!/bin/bash

function log {
  local message="$1"
  local caller="$(caller 0)"
  now=$(date +"%D %T %Z")

  if [[ ! -z "${LOGFILE:-}" ]]; then
    echo "[${now}] [${caller}] ${message}" | tee -a ${LOGFILE}
  else
    echo "[${now}] [${caller}] ${message}"
  fi
}
