#!/bin/bash

retrycmd_get_executable() {
    retries=$1; wait_sleep=$2; filepath=$3; url=$4; validation_args=$5
    echo "${retries} retries"
    for i in $(seq 1 $retries); do
        $filepath $validation_args
        [ $? -eq 0  ] && break || \
        if [ $i -eq $retries ]; then
            return 1
        else
            curl -fsSL $url -o $filepath
            chmod +x $filepath
            sleep $wait_sleep
        fi
    done
}


img_filepath=/usr/local/bin/img
retrycmd_get_executable 120 5 $img_filepath "https://acs-mirror.azureedge.net/img/img-linux-amd64-v0.4.6" lsh
