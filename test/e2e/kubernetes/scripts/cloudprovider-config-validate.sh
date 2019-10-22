#!/bin/bash

if [ $(sudo cat /etc/kubernetes/azure.json | jq --arg mode $BACKOFF_MODE '.cloudProviderBackoffMode|contains($mode)') != 'true' ]; then
    exit 1
fi
