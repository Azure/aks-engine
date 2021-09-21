#!/bin/bash

if [ -z "$RESOURCE_GROUP" ]; then
    echo "must provide a RESOURCE_GROUP env var"
    exit 1;
fi

# Continually look for non-Succeeded VMSS instances
while true; do
  NUM_VMSS=0
  NUM_TERMINAL_VMSS=0
  for VMSS in $(az vmss list -g $RESOURCE_GROUP | jq -r '.[] | .name'); do
    ((NUM_VMSS++))
    NUM_DELETED_INSTANCES=0
    VMSS_CAPACITY=$(az vmss list -g $RESOURCE_GROUP | jq -r --arg VMSS "$VMSS" '.[] | select(.name == $VMSS) | .sku.capacity')
    echo VMSS $VMSS has a current capacity of $VMSS_CAPACITY
    for TERMINAL_VMSS in $(az vmss show -g $RESOURCE_GROUP -n $VMSS | jq -r '. | select(.provisioningState == "Succeeded" or .provisioningState == "Failed") | .name'); do
      ((NUM_TERMINAL_VMSS++))
      echo VMSS $TERMINAL_VMSS is in a terminal state!
      for TARGET_VMSS_INSTANCE in $(az vmss list-instances -g $RESOURCE_GROUP -n $TERMINAL_VMSS | jq -r '.[] | select(.provisioningState == "Failed") | .name'); do
        echo Deleting VMSS $TERMINAL_VMSS instance $TARGET_VMSS_INSTANCE
        if ! az vmss delete-instances -n $TERMINAL_VMSS -g $RESOURCE_GROUP --instance-id ${TARGET_VMSS_INSTANCE##*_} --no-wait; then
          exit 1
        else
          sleep 3
          ((NUM_DELETED_INSTANCES++))
        fi
      done
      until [[ $(az vmss show -g $RESOURCE_GROUP -n $VMSS | jq -r '. | select(.provisioningState == "Succeeded") | .name') ]]; do
        sleep 30
      done
    done
    if [ "$NUM_DELETED_INSTANCES" -gt "0" ]; then
      echo Instances were deleted from VMSS $VMSS, ensuring that capacity is set to $VMSS_CAPACITY
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