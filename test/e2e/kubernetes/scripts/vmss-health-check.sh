#!/bin/bash

if [ -z "$RESOURCE_GROUP" ]; then
    echo "must provide a RESOURCE_GROUP env var"
    exit 1;
fi

# TODO: track VMSS in a "Creating" state, enforce TTL, if "Creating TTL" expires:
#  1. Check if the "Creating" VMSS instance correlates with a running Kubernetes node in the cluster
#     If so, (1) cordon/drain the node
#  2. Delete the instance in a stuck "Creating" state
#  3. Wait for the VMSS to achive a "Succeeded" ProvisioningState
#  4. Scale out the VMSS by 1

# Continually look for non-Succeeded VMSS instances
while true; do
  NUM_VMSS=0
  NUM_TERMINAL_VMSS=0
  echo "$(date)    Starting VMSS Health Remediation loop"
  for VMSS in $(az vmss list -g $RESOURCE_GROUP | jq -r '.[] | .name'); do
    ((NUM_VMSS++))
    NUM_DELETED_INSTANCES=0
    VMSS_CAPACITY=$(az vmss list -g $RESOURCE_GROUP | jq -r --arg VMSS "$VMSS" '.[] | select(.name == $VMSS) | .sku.capacity')
    echo $(date)    VMSS $VMSS has a current capacity of $VMSS_CAPACITY
    for TERMINAL_VMSS in $(az vmss show -g $RESOURCE_GROUP -n $VMSS | jq -r '. | select(.provisioningState == "Succeeded" or .provisioningState == "Failed") | .name'); do
      ((NUM_TERMINAL_VMSS++))
      echo $(date)    VMSS $TERMINAL_VMSS is in a terminal state!
      HAS_FAILED_STATE_INSTANCE="false"
      for TARGET_VMSS_INSTANCE in $(az vmss list-instances -g $RESOURCE_GROUP -n $TERMINAL_VMSS | jq -r '.[] | select(.provisioningState == "Failed") | .name'); do
        HAS_FAILED_STATE_INSTANCE="true"
        echo $(date)    Deleting VMSS $TERMINAL_VMSS instance $TARGET_VMSS_INSTANCE
        if ! az vmss delete-instances -n $TERMINAL_VMSS -g $RESOURCE_GROUP --instance-id ${TARGET_VMSS_INSTANCE##*_} --no-wait; then
          sleep 30
        else
          sleep 1
          ((NUM_DELETED_INSTANCES++))
        fi
      done
      for TARGET_VMSS_INSTANCE in $(az vmss list-instances -g $RESOURCE_GROUP -n $TERMINAL_VMSS | jq -r '.[].resources[] | select(.name == "vmssCSE" and .provisioningState == "Failed") | .id' | awk -F'/' '{print $9}'); do
        echo $(date)    Deleting VMSS $TERMINAL_VMSS instance $TARGET_VMSS_INSTANCE
        if ! az vmss delete-instances -n $TERMINAL_VMSS -g $RESOURCE_GROUP --instance-id ${TARGET_VMSS_INSTANCE##*_}; then
           sleep 30
        else
           sleep 1
           ((NUM_DELETED_INSTANCES++))
        fi
      done
      if [ "$HAS_FAILED_STATE_INSTANCE" == "true" ]; then
        echo $(date)    Waiting for $TERMINAL_VMSS to reach a terminal ProvisioningState after failed instances were deleted...
        sleep 30
        until [[ $(az vmss show -g $RESOURCE_GROUP -n $VMSS | jq -r '. | select(.provisioningState == "Succeeded" or .provisioningState == "Failed") | .name') ]]; do
          for STILL_FAILED_VMSS_INSTANCE in $(az vmss list-instances -g $RESOURCE_GROUP -n $TERMINAL_VMSS | jq -r '.[] | select(.provisioningState == "Failed") | .name'); do
            echo $(date)    Instance $STILL_FAILED_VMSS_INSTANCE is still in a failed state, will attempt to delete again in the next loop
          done
        done
      fi
    done
    if [ "$NUM_DELETED_INSTANCES" -gt "0" ]; then
      echo $(date)    Instances were deleted from VMSS $VMSS, ensuring that capacity is set to $VMSS_CAPACITY
      if ! az vmss scale --new-capacity $VMSS_CAPACITY -n $VMSS -g $RESOURCE_GROUP; then
          exit 1
      fi
    fi
  done
  if [ "$LOOP_FOREVER" == "true" ]; then
    sleep 150
  else
    if [[ "${NUM_VMSS}" == "${NUM_TERMINAL_VMSS}" ]]; then
      exit 0
    fi
  fi
done